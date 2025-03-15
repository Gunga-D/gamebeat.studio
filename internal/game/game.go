package game

import (
	"github.com/releaseband/golang-developer-test/internal/game/result"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

// RoundCost - функция, которая возвращет стоимость одного раунда
func RoundCost(linesCount int) uint64 {
	return uint64(linesCount)
}

type Slot struct {
	generator  Generator
	calculator Calculator
	roundCost  uint64
}

func newSlot(generator Generator, calculator Calculator, roundCost uint64) *Slot {
	return &Slot{generator: generator, calculator: calculator, roundCost: roundCost}
}

func (s *Slot) Spin(rng rng.RNG) (*result.Round, error) {
	// todo: implement me
	return nil, nil
}

func (s *Slot) RoundCost() uint64 {
	return s.roundCost
}
