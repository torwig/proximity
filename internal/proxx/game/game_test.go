package game_test

import (
	"proxx/internal/proxx/board"
	"proxx/internal/proxx/game"
	"proxx/internal/proxx/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type predefinedBlackHoleLocator struct {
	positions []board.Position
}

func newPredefinedBlackHoleLocator(positions []board.Position) *predefinedBlackHoleLocator {
	return &predefinedBlackHoleLocator{positions: positions}
}

func (p predefinedBlackHoleLocator) LocateBlackHolesOnBoard(_ int, _ int, _ int) []board.Position {
	return p.positions
}

func TestGame_OpenCell(t *testing.T) {
	gameCfg := game.Config{
		NumRows:          3,
		NumCols:          3,
		NumBlackHoles:    2,
		BlackHoleLocator: newPredefinedBlackHoleLocator([]board.Position{{Row: 1, Col: 2}, {Row: 2, Col: 2}}),
	}

	t.Run("Open a cell outside the board", func(t *testing.T) {
		t.Parallel()

		g, err := game.NewGame(gameCfg)
		require.NoError(t, err)
		require.False(t, g.IsOver())

		err = g.OpenCell(4, 2)
		assert.Error(t, err)

		err = g.OpenCell(1, 5)
		assert.Error(t, err)

		err = g.OpenCell(-1, 3)
		assert.Error(t, err)

		err = g.OpenCell(2, -2)
		assert.Error(t, err)
	})

	t.Run("Open a blank cell", func(t *testing.T) {
		t.Parallel()

		g, err := game.NewGame(gameCfg)
		require.NoError(t, err)
		require.False(t, g.IsOver())

		err = g.OpenCell(2, 0)
		assert.NoError(t, err)

		expectedState := [][]board.CellValue{
			{"0", "1", "?"},
			{"0", "2", "?"},
			{"0", "2", "?"},
		}

		currentState := g.BoardState()
		testhelpers.EqualBoardStates(t, expectedState, currentState)

		// open the same (already opened) cell again, nothing should change
		err = g.OpenCell(2, 0)
		assert.NoError(t, err)

		currentState = g.BoardState()
		testhelpers.EqualBoardStates(t, expectedState, currentState)
	})

	t.Run("Open a cell with a clue", func(t *testing.T) {
		t.Parallel()

		g, err := game.NewGame(gameCfg)
		require.NoError(t, err)
		require.False(t, g.IsOver())

		err = g.OpenCell(1, 1)
		assert.NoError(t, err)

		expectedState := [][]board.CellValue{
			{"?", "?", "?"},
			{"?", "2", "?"},
			{"?", "?", "?"},
		}

		currentState := g.BoardState()
		testhelpers.EqualBoardStates(t, expectedState, currentState)
	})

	t.Run("Open a cell with a black hole", func(t *testing.T) {
		t.Parallel()

		g, err := game.NewGame(gameCfg)
		require.NoError(t, err)
		require.False(t, g.IsOver())

		err = g.OpenCell(2, 2)
		assert.NoError(t, err)

		expectedState := [][]board.CellValue{
			{"?", "?", "?"},
			{"?", "?", "H"},
			{"?", "?", "H"},
		}

		currentState := g.BoardState()
		testhelpers.EqualBoardStates(t, expectedState, currentState)

		assert.True(t, g.IsOver())
		assert.False(t, g.IsWon())
	})

	t.Run("Win a simple game", func(t *testing.T) {
		t.Parallel()

		g, err := game.NewGame(gameCfg)
		require.NoError(t, err)
		require.False(t, g.IsOver())

		err = g.OpenCell(1, 0)
		assert.NoError(t, err)

		err = g.OpenCell(0, 2)
		assert.NoError(t, err)

		expectedState := [][]board.CellValue{
			{"0", "1", "1"},
			{"0", "2", "?"},
			{"0", "2", "?"},
		}

		currentState := g.BoardState()
		testhelpers.EqualBoardStates(t, expectedState, currentState)

		assert.True(t, g.IsOver())
		assert.True(t, g.IsWon())
	})
}
