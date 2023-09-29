package testhelpers

import (
	"proxx/internal/proxx/board"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func EqualBoardStates(t *testing.T, want [][]board.CellValue, got [][]board.CellValue) {
	require.EqualValues(t, len(want), len(got))
	for i := range want {
		assert.EqualValues(t, want[i], got[i])
	}
}
