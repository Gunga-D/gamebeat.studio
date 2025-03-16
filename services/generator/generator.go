package generator

import (
	"errors"

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
	res := symbols.Reels{}
	for reelIdx := 0; reelIdx < len(s.gameTapes); reelIdx++ {
		rowIdx := rng.Random(0, uint32(len(s.gameTapes[reelIdx])-1))
		reelSymbols, err := s.GetReelSymbols(reelIdx, int(rowIdx))
		if err != nil {
			return symbols.Reels{}, err
		}
		res = append(res, reelSymbols)
	}
	return res, nil
}

func (s *Symbols) GetReelSymbols(reelIndex int, rowIndex int) (symbols.Symbols, error) {
	if reelIndex >= len(s.gameTapes) || reelIndex < 0 {
		return symbols.Symbols{}, errors.New("invalid reelIndex")
	}
	reelList := s.gameTapes[reelIndex]

	if rowIndex < 0 {
		return symbols.Symbols{}, errors.New("invalid rowIndex")
	}

	res := symbols.Symbols{}
	for idx := rowIndex; idx < rowIndex+s.rowsCount; idx++ {
		if len(reelList) <= idx {
			res = append(res, reelList[idx-len(reelList)])
			continue
		}
		res = append(res, reelList[idx])
	}
	return res, nil
}
