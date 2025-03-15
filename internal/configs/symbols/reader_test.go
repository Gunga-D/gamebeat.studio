package symbols

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadReels(t *testing.T) {
	resp, err := ReadReels()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 5, len(resp))
	require.Equal(t, 30, len(resp[0]))
	require.Equal(t, 27, len(resp[1]))
	require.Equal(t, 31, len(resp[2]))
	require.Equal(t, 31, len(resp[3]))
	require.Equal(t, 29, len(resp[4]))
}
