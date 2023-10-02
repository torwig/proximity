package board_test

import (
	"proxx/internal/proxx/board"
	"proxx/internal/proxx/testhelpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_NewBoard(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		cfg         board.Config
	}{
		{name: "Zero number of rows", errExpected: true, cfg: board.Config{NumRows: 0, NumCols: 10}},
		{name: "Negative number of rows", errExpected: true, cfg: board.Config{NumRows: -10, NumCols: 10}},
		{name: "Zero number of columns", errExpected: true, cfg: board.Config{NumRows: 7, NumCols: 0}},
		{name: "Negative number of columns", errExpected: true, cfg: board.Config{NumRows: 7, NumCols: -5}},
		{name: "Valid configuration: common case", errExpected: false, cfg: board.Config{NumRows: 5, NumCols: 5}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gameBoard, err := board.NewBoard(tc.cfg)
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, tc.cfg.NumRows*tc.cfg.NumCols, gameBoard.TotalNumberOfCells())
			}
		})
	}
}

func TestBoard_Init(t *testing.T) {
	t.Run("Black hole distribution across the board", func(t *testing.T) {
		t.Parallel()

		blackHolePositions := []board.Position{
			{0, 3},
			{2, 1}, {2, 2}, {2, 3},
			{3, 1}, {3, 3},
			{4, 1}, {4, 2}, {4, 3},
		}

		boardCfg := board.Config{
			NumRows: 5,
			NumCols: 5,
		}

		gameBoard, err := board.NewBoard(boardCfg)
		require.NoError(t, err)

		gameBoard.Init(blackHolePositions)

		expectedState := [][]board.CellValue{
			{"0", "0", "1", "H", "1"},
			{"1", "2", "4", "3", "2"},
			{"2", "H", "H", "H", "2"},
			{"3", "H", "8", "H", "3"},
			{"2", "H", "H", "H", "2"},
		}

		currentState := gameBoard.DebugState()
		assert.Len(t, currentState, boardCfg.NumRows)
		testhelpers.EqualBoardStates(t, expectedState, currentState)
		assert.EqualValues(t, 0, gameBoard.OpenedCells())
	})
}
