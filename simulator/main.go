package main

import (
	"fmt"
	"github.com/releaseband/golang-developer-test/internal/game"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

const (
	spinsCount = 20_000_00
	seed       = 1471471747174
	expRTP     = 0.9688362
)

func run() error {
	slot, err := game.New()
	if err != nil {
		return fmt.Errorf("new slot: %w", err)
	}

	rand := rng.NewRNG()
	rand.Reseed(seed)

	var totalWin uint64

	for i := 0; i < spinsCount; i++ {
		res, err := slot.Spin(rand)
		if err != nil {
			return fmt.Errorf("spin: %w", err)
		}

		for _, w := range res.Wins() {
			totalWin += w.Amount()
		}
	}

	totalBet := slot.RoundCost() * spinsCount

	gotRTP := float64(totalWin) / float64(totalBet)

	fmt.Printf("exp=%v\n", expRTP)
	fmt.Printf("got=%v\n", gotRTP)

	if expRTP != gotRTP {
		return fmt.Errorf("exp=%f, got=%f", expRTP, gotRTP)
	}

	fmt.Println("< Congratulations >")

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
