package calculator

import (
	"testing"

	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
	"github.com/stretchr/testify/require"
)

func TestCalculator_Calculate(t *testing.T) {
	payTable := paytable.NewPayTable(map[symbols.Symbol]paytable.Payout{
		1: []uint64{0, 0, 3, 4, 5},
		2: []uint64{0, 0, 0, 6, 7},
		3: []uint64{0, 0, 0, 0, 8},
	})

	line1 := lines.NewLine([]int{0, 0, 0, 0, 0})
	line2 := lines.NewLine([]int{1, 1, 1, 1, 1})
	line3 := lines.NewLine([]int{2, 2, 2, 2, 2})

	gamLines := lines.Lines{*line1, *line2, *line3}

	calculator := NewCalculator(gamLines, payTable)

	tests := []struct {
		name        string
		exp         []win.Win
		spinSymbols symbols.Reels
	}{
		{
			name: "without_win",
			exp:  nil,
			spinSymbols: []symbols.Symbols{
				{1, 2, 1}, // 1 reel
				{3, 2, 1}, // 2 reel
				{1, 3, 3}, // 3 reel
				{4, 1, 3}, // 4 reel
				{1, 4, 4}, // 5 reel
			},
		},
		{
			name: "shuffle_without_win",
			exp:  nil,
			spinSymbols: []symbols.Symbols{
				{3, 2, 1}, // 1 reel
				{3, 2, 1}, // 2 reel
				{3, 3, 3}, // 3 reel
				{1, 3, 3}, // 4 reel
				{4, 4, 4}, // 5 reel
			},
		},
		{
			name: "win_by_3",
			spinSymbols: []symbols.Symbols{
				{2, 3, 1}, // 1 reel
				{2, 3, 1}, // 2 reel
				{3, 3, 3}, // 3 reel
				{1, 3, 1}, // 4 reel
				{4, 3, 4}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(8, symbols.Symbols{3, 3, 3, 3, 3}, 3),
			},
		},
		{
			name: "win_by_2",
			spinSymbols: []symbols.Symbols{
				{2, 3, 1}, // 1 reel
				{2, 3, 1}, // 2 reel
				{2, 3, 3}, // 3 reel
				{2, 6, 1}, // 4 reel
				{4, 5, 4}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(6, symbols.Symbols{2, 2, 2, 2}, 2),
			},
		},
		{
			name: "win_by_1_and_2",
			spinSymbols: []symbols.Symbols{
				{2, 3, 1}, // 1 reel
				{2, 3, 1}, // 2 reel
				{2, 3, 1}, // 3 reel
				{2, 6, 7}, // 4 reel
				{2, 5, 9}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(7, symbols.Symbols{2, 2, 2, 2, 2}, 2),
				win.NewWin(3, symbols.Symbols{1, 1, 1}, 1),
			},
		},
		{
			name: "win_by_1,2,3",
			spinSymbols: []symbols.Symbols{
				{2, 3, 1}, // 1 reel
				{2, 3, 1}, // 2 reel
				{2, 3, 1}, // 3 reel
				{2, 3, 1}, // 4 reel
				{2, 3, 1}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(7, symbols.Symbols{2, 2, 2, 2, 2}, 2),
				win.NewWin(8, symbols.Symbols{3, 3, 3, 3, 3}, 3),
				win.NewWin(5, symbols.Symbols{1, 1, 1, 1, 1}, 1),
			},
		},
		{
			name: "win_with_wildcard",
			spinSymbols: []symbols.Symbols{
				{2, 1, 1}, // 1 reel
				{2, 2, 2}, // 2 reel
				{0, 3, 3}, // 3 reel
				{2, 1, 1}, // 4 reel
				{2, 2, 2}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(7, symbols.Symbols{2, 2, 0, 2, 2}, 2),
			},
		},
		{
			name: "start_win_by_wildcard",
			spinSymbols: []symbols.Symbols{
				{0, 1, 1}, // 1 reel
				{2, 2, 2}, // 2 reel
				{2, 3, 3}, // 3 reel
				{2, 1, 1}, // 4 reel
				{2, 2, 2}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(7, symbols.Symbols{0, 2, 2, 2, 2}, 2),
			},
		},
		{
			name: "wildcard_to_first_win",
			spinSymbols: []symbols.Symbols{
				{1, 1, 1}, // 1 reel
				{1, 2, 2}, // 2 reel
				{0, 3, 3}, // 3 reel
				{2, 1, 1}, // 4 reel
				{2, 2, 2}, // 5 reel
			},
			exp: []win.Win{
				win.NewWin(3, symbols.Symbols{1, 1, 0}, 1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.spinSymbols)
			require.NoError(t, err)
			require.Equal(t, tt.exp, got)
		})
	}
}
