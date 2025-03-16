package paytable

import (
	"testing"

	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadPayTable(t *testing.T) {
	pt, err := ReadPayTable()
	require.NoError(t, err)
	require.NotNil(t, pt)
	require.Equal(t, 8, len(pt.symbolPayouts))

	for _, symbolPayout := range pt.symbolPayouts {
		require.Equal(t, 5, len(symbolPayout))
	}
}

func TestParsePayouts(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    [][]string
		err      bool
		expected map[symbols.Symbol]Payout
	}{
		{
			name:  "no_input",
			input: [][]string{},
			err:   true,
		},
		{
			name: "no_digits_value",
			input: [][]string{
				{
					"1",
					"test",
				},
			},
			err: true,
		},
		{
			name: "success",
			input: [][]string{
				{
					"0",
					"0",
					"50",
					"200",
					"500",
				},
				{
					"1",
					"2",
					"50",
					"100",
					"300",
				},
			},
			err: false,
			expected: map[symbols.Symbol]Payout{
				0: {
					0,
					0,
					50,
					200,
					500,
				},
				1: {
					1,
					2,
					50,
					100,
					300,
				},
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := parsePayouts(tc.input)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NotNil(t, res)
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}
