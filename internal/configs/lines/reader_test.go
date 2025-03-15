package lines

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadLines(t *testing.T) {
	// todo: implement me
	res, err := ReadLines()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 5, len(res))
}
