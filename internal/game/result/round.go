package result

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

type Round struct {
	// symbols - символы, которые выпали в раунде
	symbols []symbols.Symbols
	// wins - выигрыши в раунде
	wins []win.Win
	// cost - стоимость раунда
	cost uint64
}

func NewRound(symbols []symbols.Symbols, wins []win.Win, cost uint64) *Round {
	return &Round{symbols: symbols, wins: wins, cost: cost}
}

func (r *Round) Wins() []win.Win {
	return r.wins
}
