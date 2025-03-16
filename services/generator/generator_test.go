package generator

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/rng"
	"github.com/stretchr/testify/require"
)

func TestSymbols_GetReelSymbols(t *testing.T) {
	const (
		rowsCount = 3
	)

	gameReels := NewSymbols(rowsCount, []symbols.Symbols{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 51, 3, 2, 1, 8, 9, 12, 19, 4},
		{11, 9, 3, 7, 7, 5, 6, 8, 9, 10},
		{10, 21, 33, 44, 55, 66, 67, 88, 99, 11},
	})

	tests := []struct {
		name      string
		reelIndex int
		rowIndex  int
		hasErr    bool
		exp       symbols.Symbols
	}{
		{
			name:      "from_the_beginning_of_the_list",
			reelIndex: 0,
			rowIndex:  0,
			exp:       symbols.Symbols{1, 2, 3},
		},
		{
			name:      "from_the_middle_of_the_list",
			reelIndex: 4,
			rowIndex:  4,
			exp:       symbols.Symbols{55, 66, 67},
		},
		{
			name:      "from_the_end_of_the_list",
			reelIndex: 1,
			rowIndex:  9,
			exp:       symbols.Symbols{1, 10, 9},
		},
		{
			name:      "invalid_reel_index_less_zero",
			reelIndex: -1,
			hasErr:    true,
		},
		{
			name:      "invalid_reel_index_max",
			reelIndex: 15,
			hasErr:    true,
		},
		{
			name:      "invalid_row_index",
			reelIndex: 0,
			rowIndex:  -1,
			hasErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gameReels.GetReelSymbols(tt.reelIndex, tt.rowIndex)
			if tt.hasErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.exp, got)
			}
		})
	}
}

func TestSymbols_Generate(t *testing.T) {
	const (
		rowsCount = 3
	)

	gameReels := NewSymbols(rowsCount, []symbols.Symbols{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 51, 3, 2, 1, 8, 9, 12, 19, 4},
		{11, 9, 3, 7, 7, 5, 6, 8, 9, 10},
		{10, 21, 33, 44, 55, 66, 67, 88, 99, 11},
	})
	tests := []struct {
		name   string
		rng    func(*rng.MockRNG)
		hasErr bool
		exp    symbols.Reels
	}{
		{
			name: "success",
			rng: func(m *rng.MockRNG) {
				m.EXPECT().Random(uint32(0), uint32(9)).Return(uint32(1))
				m.EXPECT().Random(uint32(0), uint32(9)).Return(uint32(5))
				m.EXPECT().Random(uint32(0), uint32(9)).Return(uint32(2))
				m.EXPECT().Random(uint32(0), uint32(9)).Return(uint32(7))
				m.EXPECT().Random(uint32(0), uint32(9)).Return(uint32(7))
			},
			exp: symbols.Reels{
				{
					2, 3, 4,
				},
				{
					5, 4, 3,
				},
				{
					3, 2, 1,
				},
				{
					8, 9, 10,
				},
				{
					88, 99, 11,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := rng.NewMockRNG(ctrl)
			if tt.rng != nil {
				tt.rng(m)
			}

			got, err := gameReels.Generate(m)
			if tt.hasErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.exp, got)
			}
		})
	}
}
