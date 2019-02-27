package glicko

import (
    "math"
)

type RatingPeriod struct {
    tau     float64
    players []*Player
}

func NewRatingPeriod() *RatingPeriod {
    tau := 0.5
    return &RatingPeriod{
        tau:     tau,
        players: []*Player{},
    }
}

func (period *RatingPeriod) AddPlayer(player *Player) {
    // @todo cache
    for _, p := range period.players {
        if p == player {
            return
        }
    }
    period.players = append(period.players, player)
}

func (period *RatingPeriod) AddMatch(player1 *Player, player2 *Player, score MatchResult) {
    period.AddPlayer(player1)
    period.AddPlayer(player2)

    match := &match{
        player1: player1,
        player2: player2,
        score:   score,
    }
    player1.addMatch(match)
    player2.addMatch(match)
}

func (period *RatingPeriod) Calculate() {
    for _, player := range period.players {
        if len(player.matches) > 0 {
            v := v(player)
            dp := delta(player)
            delta := v * dp

            sigmaP := sigmaP(delta, player.pre.sigma, player.pre.phi, v, period.tau)
            phiS := phiA(player.pre.phi, sigmaP)
            phiP := phiP(phiS, v)
            muP := muP(player.pre.mu, phiP, dp)
            player.post.Update(muP, phiP, sigmaP)
        } else {
            player.post.Touch()
        }
    }
}

// step 3
func v(player *Player) float64 {
    v := 0.0
    for _, match := range player.matches {
        opponent := match.opponentFor(player)

        g := g(opponent.pre.phi)
        E := e(player.pre.mu, opponent.pre.mu, opponent.pre.phi)
        vj := g * g * E * (1 - E)
        v += vj
    }

    return 1 / v
}

func g(phiJ float64) float64 {
    return 1 / math.Sqrt(1+3*math.Pow(phiJ, 2)/math.Pow(math.Pi, 2))
}

func e(mu float64, muJ float64, phiJ float64) float64 {
    return 1 / (1 + math.Exp(-g(phiJ)*(mu-muJ)))
}

// step 4
func delta(player *Player) float64 {
    outcomeBasedRating := 0.0
    for _, match := range player.matches {
        opponent := match.opponentFor(player)
        ophi := g(opponent.pre.phi)
        sc := float64(match.resultFor(player))
        e := e(player.pre.mu, opponent.pre.mu, opponent.pre.phi)

        outcomeBasedRating += ophi * (sc - e)
    }

    return outcomeBasedRating
}

// step 5
func sigmaP(delta float64, sigma float64, phi float64, v float64, tau float64) float64 {
    a := math.Log(math.Pow(sigma, 2))
    A := a
    fX := func(x float64, delta float64, phi float64, v float64, a float64, tau float64) float64 {
        return ((math.Exp(x) * (math.Pow(delta, 2) - math.Pow(phi, 2) - v - math.Exp(x))) / (2 * math.Pow((math.Pow(phi, 2) + v + math.Exp(x)), 2))) - ((x - a) / math.Pow(float64(tau), 2))
    }
    epsilon := 0.000001

    var B float64
    if math.Pow(delta, 2) > (math.Pow(phi, 2) + v) {
        B = math.Log(math.Pow(delta, 2) - math.Pow(phi, 2) - v)
    } else {
        k := float64(1)
        for fX(a-k*tau, delta, phi, v, a, tau) < 0 {
            k++
        }
        B = a - k*tau
    }

    fA := fX(A, delta, phi, v, a, tau)
    fB := fX(B, delta, phi, v, a, tau)

    for math.Abs(B-A) > epsilon {
        C := A + fA*(A-B)/(fB-fA)
        fC := fX(C, delta, phi, v, a, tau)
        if (fC * fB) < 0 {
            A = B
            fA = fB
        } else {
            fA = fA / 2
        }
        B = C
        fB = fC
    }

    return math.Exp(A / 2)
}

// step 6
func phiA(phi float64, sigmaP float64) float64 {
    return math.Sqrt(math.Pow(phi, 2) + math.Pow(sigmaP, 2))
}

// step 7
func phiP(phiS float64, v float64) float64 {
    return 1 / math.Sqrt(1/math.Pow(phiS, 2)+1/v)
}

func muP(mu float64, phiP float64, delta float64) float64 {
    return mu + math.Pow(phiP, 2)*delta
}
