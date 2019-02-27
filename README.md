# Go-Glicko2

A Go implementation of [Glicko2 rating system](http://www.glicko.net/glicko.html)

## Installation

```go get -u github.com/zelenin/go-glicko2```

## Usage

```go
package main

import (
    "fmt"
    "github.com/zelenin/go-glicko2"
)

func main() {
    player1 := glicko.NewPlayer(glicko.NewRating(1500, 200, 0.06))
    player2 := glicko.NewPlayer(glicko.NewRating(1400, 30, 0.06))
    player3 := glicko.NewPlayer(glicko.NewRating(1550, 100, 0.06))
    player4 := glicko.NewPlayer(glicko.NewRating(1700, 300, 0.06))

    period := glicko.NewRatingPeriod()

    period.AddMatch(player1, player2, glicko.MATCH_RESULT_WIN)
    period.AddMatch(player1, player3, glicko.MATCH_RESULT_LOSS)
    period.AddMatch(player1, player4, glicko.MATCH_RESULT_LOSS)

    period.Calculate()

    fmt.Printf("Player #1 rating: %0.2f\n", player1.Rating().R())
    fmt.Printf("Player #2 rating: %0.2f\n", player2.Rating().R())
    fmt.Printf("Player #3 rating: %0.2f\n", player3.Rating().R())
    fmt.Printf("Player #4 rating: %0.2f\n", player4.Rating().R())
}
```

## Author

[Aleksandr Zelenin](https://github.com/zelenin/), e-mail: [aleksandr@zelenin.me](mailto:aleksandr@zelenin.me)
