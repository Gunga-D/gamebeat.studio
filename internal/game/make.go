package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/services/calculator"
	"github.com/releaseband/golang-developer-test/services/generator"
)

const rowsCount = 3

func New() (*Slot, error) {
	l, err := lines.ReadLines()
	if err != nil {
		return nil, err
	}
	payTable, err := paytable.ReadPayTable()
	if err != nil {
		return nil, err
	}
	syms, err := symbols.ReadReels()
	if err != nil {
		return nil, err
	}

	calc := calculator.NewCalculator(l, payTable)
	gen := generator.NewSymbols(rowsCount, syms)
	return newSlot(gen, calc, RoundCost(len(l))), nil
}
