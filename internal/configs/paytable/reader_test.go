package paytable

import (
	"github.com/stretchr/testify/require"
	"testing"
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
