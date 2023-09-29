package game_test

import (
	"proxx/internal/proxx/board"
	"proxx/internal/proxx/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniformBlackHoleLocator_LocateBlackHolesOnBoard(t *testing.T) {
	locator := game.NewUniformBlackHoleLocator()

	testCases := []struct {
		name   string
		rowNum int
		colNum int
		bhNum  int
	}{
		{"5x5 with 10 black holes", 5, 5, 10},
		{"3x3 with 1 black hole", 3, 3, 1},
		{"2x2 with 3 black holes", 2, 2, 3},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			positions := locator.LocateBlackHolesOnBoard(tc.rowNum, tc.colNum, tc.bhNum)
			assert.Len(t, positions, tc.bhNum)
			assert.EqualValues(t, len(positions), countUniquePositions(positions))
		})
	}
}

func countUniquePositions(positions []board.Position) int {
	m := make(map[board.Position]struct{}, len(positions))

	for _, p := range positions {
		m[p] = struct{}{}
	}

	return len(m)
}
