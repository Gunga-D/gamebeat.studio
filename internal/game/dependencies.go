package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

type Calculator interface {
	Calculate(symbols symbols.Reels) ([]win.Win, error)
}

type Generator interface {
	Generate(rng rng.RNG) (symbols.Reels, error)
}
