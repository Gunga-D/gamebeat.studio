package generator

import (
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSymbols_GetReelSymbols(t *testing.T) {
	const (
		rowsCount = 3
	)

	gameReels := NewSymbols(rowsCount, []symbols.Symbols{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	})

	tests := []struct {
		name      string
		reelIndex int
		rowIndex  int
		exp       symbols.Symbols
	}{
		{
			name:      "from the beginning of the list",
			reelIndex: 2,
			rowIndex:  0,
			exp:       symbols.Symbols{1, 2, 3},
		},
		{
			name:      "from the middle of the list",
			reelIndex: 4,
			rowIndex:  4,
			exp:       symbols.Symbols{5, 6, 7},
		},
		{
			name:      "from the end of the list",
			reelIndex: 1,
			rowIndex:  9,
			exp:       symbols.Symbols{10, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gameReels.GetReelSymbols(tt.reelIndex, tt.rowIndex)
			require.Equal(t, tt.exp, got)
		})
	}
}
