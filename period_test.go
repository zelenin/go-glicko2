package glicko

import (
    "math"
    "testing"
    "strconv"
)

const accuracy int = 2

func TestCalculateMatch(t *testing.T) {
    player1 := NewPlayer(NewRating(1500, 200, 0.06))
    player2 := NewPlayer(NewRating(1400, 30, 0.06))
    player3 := NewPlayer(NewRating(1550, 100, 0.06))
    player4 := NewPlayer(NewRating(1700, 300, 0.06))

    period := NewRatingPeriod()

    period.AddMatch(player1, player2, MATCH_RESULT_WIN)
    period.AddMatch(player1, player3, MATCH_RESULT_LOSS)
    period.AddMatch(player1, player4, MATCH_RESULT_LOSS)

    period.Calculate()

    strAccuracy := strconv.Itoa(accuracy)

    if !isEqual(player1.Rating().R(), 1464.05067) {
        t.Errorf("Fail: %s, %."+strAccuracy+"f", "player.R()", player1.Rating().R())
    }

    if !isEqual(player1.Rating().Rd(), 151.51652) {
        t.Errorf("Fail: %s, %."+strAccuracy+"f", "player.Rd()", player1.Rating().Rd())
    }

    if !isEqual(player1.Rating().Sigma(), 0.05999) {
        t.Errorf("Fail: %s, %."+strAccuracy+"f", "player.Sigma()", player1.Rating().Sigma())
    }
}

func floor(value float64, accuracy int) float64 {
    mp := math.Pow(10, 2)
    return math.Floor(value*mp) / mp
}

func isEqual(a, b float64) bool {
    return floor(a, accuracy) == floor(b, accuracy)
}
