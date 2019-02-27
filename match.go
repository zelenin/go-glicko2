package glicko

type MatchResult float64

const (
    MATCH_RESULT_WIN  MatchResult = 1.0
    MATCH_RESULT_DRAW MatchResult = 0.5
    MATCH_RESULT_LOSS MatchResult = 0.0
)

type match struct {
    player1 *Player
    player2 *Player
    score   MatchResult
}

func (match *match) opponentFor(player *Player) *Player {
    if match.player1 == player {
        return match.player2
    }

    return match.player1
}

func (match *match) resultFor(player *Player) MatchResult {
    if match.player1 == player {
        return match.score
    }

    return 1 - match.score
}

func MatchResultFromScore(score1 float64, score2 float64) MatchResult {
    matchResult := MATCH_RESULT_DRAW

    diff := score1 - score2
    switch {
    case diff < 0:
        matchResult = MATCH_RESULT_LOSS
    case diff > 0:
        matchResult = MATCH_RESULT_WIN
    }

    return matchResult
}
