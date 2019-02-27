package glicko

const (
    RATING_SCALE_PARAMETER = 173.7178
    RATING_BASE_R          = 1500
    RATING_BASE_RD         = 350
    RATING_BASE_SIGMA      = 0.06
)

type Rating struct {
    r     float64
    mu    float64
    rd    float64
    phi   float64
    sigma float64
}

func (rating Rating) R() float64 {
    return rating.r
}

func (rating Rating) Rd() float64 {
    return rating.rd
}

func (rating Rating) Sigma() float64 {
    return rating.sigma
}

func (rating Rating) ConfidenceInterval() (float64, float64) {
    return rating.R() - 2*rating.Rd(), rating.R() + 2*rating.Rd()
}

func (rating *Rating) Update(mu float64, phi float64, sigma float64) {
    setMu(rating, mu)
    setPhi(rating, phi)
    setSigma(rating, sigma)
}

func (rating *Rating) Touch() {
    setPhi(rating, phiA(rating.phi, rating.sigma))
}

func setR(rating *Rating, r float64) {
    rating.r = r
    rating.mu = (rating.r - RATING_BASE_R) / RATING_SCALE_PARAMETER
}

func setMu(rating *Rating, mu float64) {
    rating.mu = mu
    rating.r = rating.mu*RATING_SCALE_PARAMETER + RATING_BASE_R
}

func setRd(rating *Rating, rd float64) {
    rating.rd = rd
    rating.phi = rating.rd / RATING_SCALE_PARAMETER
}

func setPhi(rating *Rating, phi float64) {
    rating.phi = phi
    rating.rd = rating.phi * RATING_SCALE_PARAMETER
}

func setSigma(rating *Rating, sigma float64) {
    rating.sigma = sigma
}

func NewRating(r float64, rd float64, sigma float64) *Rating {
    rating := &Rating{}

    setR(rating, r)
    setRd(rating, rd)
    setSigma(rating, sigma)

    return rating
}

func NewDefaultRating() *Rating {
    return NewRating(RATING_BASE_R, RATING_BASE_RD, RATING_BASE_SIGMA)
}
