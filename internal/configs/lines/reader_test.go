package lines

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	expected := Lines{
		Line{indices: []int{1, 1, 1, 1, 1}},
		Line{indices: []int{0, 0, 0, 0, 0}},
		Line{indices: []int{2, 2, 2, 2, 2}},
		Line{indices: []int{0, 1, 2, 1, 0}},
		Line{indices: []int{2, 1, 0, 1, 2}},
	}
	res, err := ReadLines()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 5, len(res))
	assert.Equal(t, expected, res)
}

func TestParseLine(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    []string
		err      bool
		expected *Line
	}{
		{
			name:  "no_input",
			input: []string{},
			err:   true,
		},
		{
			name: "no_digits_value",
			input: []string{
				"1",
				"test",
			},
			err: true,
		},
		{
			name: "success",
			input: []string{
				"1",
				"1",
				"1",
				"1",
				"1",
			},
			err: false,
			expected: &Line{
				indices: []int{
					1, 1, 1, 1, 1,
				},
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := parseLine(tc.input)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NotNil(t, res)
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}
