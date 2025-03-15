package win

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

type Win struct {
	amount  uint64
	symbols symbols.Symbols
	symbol  symbols.Symbol
}

func NewWin(
	amount uint64,
	symbols symbols.Symbols,
	symbol symbols.Symbol,
) Win {
	return Win{
		amount:  amount,
		symbols: symbols,
		symbol:  symbol,
	}
}

func (w Win) Amount() uint64 {
	return w.amount
}
