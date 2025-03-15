package generator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

type Symbols struct {
	rowsCount int
	gameTapes []symbols.Symbols
}

func NewSymbols(rowsCount int, gameTapes []symbols.Symbols) *Symbols {
	return &Symbols{rowsCount: rowsCount, gameTapes: gameTapes}
}

func (s *Symbols) Generate(rng rng.RNG) (symbols.Reels, error) {
	// todo: implement me
	return nil, nil
}

func (s *Symbols) GetReelSymbols(reelIndex int, rowIndex int) symbols.Symbols {
	// todo: implement me
	return nil
}
