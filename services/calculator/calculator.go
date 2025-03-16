package calculator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
const WILD = symbols.Symbol(0)

type Calculator struct {
	lines    lines.Lines
	payTable *paytable.PayTable
}

func NewCalculator(lines lines.Lines, payTable *paytable.PayTable) *Calculator {
	return &Calculator{lines: lines, payTable: payTable}
}

func (c *Calculator) Calculate(spinSymbols symbols.Reels) ([]win.Win, error) {
	var wins []win.Win
	for _, l := range c.lines {
		reelIdx := 0
		nextReelIdx := reelIdx + 1
		winComb := map[symbols.Symbol][]symbols.Symbol{}
		for nextReelIdx < len(l.GetIndices()) {
			rowIdx := l.GetIndices()[reelIdx]
			s := spinSymbols[reelIdx][rowIdx]

			nextRowIdx := l.GetIndices()[nextReelIdx]
			nextS := spinSymbols[nextReelIdx][nextRowIdx]

			// Для специального символа WILD сделано допущение, что в случае, когда
			// следующие символы на рулетке:
			// reel: [1 1 0 2 2]
			// то WILD начинает относится к первому собранному ряду, то есть [1 1 0]
			// Для такого кейса написан тест - "wildcard_to_first_win"
			mainSymbol := s
			if mainSymbol == WILD {
				mainSymbol = nextS
			}
			if nextS == s || nextS == WILD || s == WILD {
				if _, found := winComb[mainSymbol]; !found {
					winComb[mainSymbol] = append(winComb[mainSymbol], s, nextS)
				} else {
					winComb[mainSymbol] = append(winComb[mainSymbol], nextS)
				}
			} else {
				reelIdx = nextReelIdx
			}
			nextReelIdx++
		}

		for symb, comb := range winComb {
			amountToPay, err := c.payTable.Get(symb, len(comb)-1)
			if err != nil {
				return nil, err
			}
			if amountToPay > 0 {
				wins = append(wins, win.NewWin(amountToPay, comb, symb))
			}
		}
	}
	return wins, nil
}
