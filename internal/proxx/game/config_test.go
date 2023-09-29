package game_test

import (
	"proxx/internal/proxx/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	bhLocator := game.NewUniformBlackHoleLocator()

	testCases := []struct {
		name        string
		errExpected bool
		cfg         game.Config
	}{
		{name: "Zero number of black holes", errExpected: true,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 0, BlackHoleLocator: bhLocator}},
		{name: "Negative number of black holes", errExpected: true,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: -10, BlackHoleLocator: bhLocator}},
		// NOTE: proxx.com permits to have the 8x8 board with 64 black holes on it => you will win with the first click
		{name: "Too many black holes: all cells occupied", errExpected: true,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 25, BlackHoleLocator: bhLocator}},
		{name: "Too many black holes: more than cells", errExpected: true,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 30, BlackHoleLocator: bhLocator}},
		{name: "Black hole locator missing", errExpected: true,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 30}},
		{name: "Valid configuration: one cell for a clue", errExpected: false,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 24, BlackHoleLocator: bhLocator}},
		{name: "Valid configuration: enough space for multiple black holes", errExpected: false,
			cfg: game.Config{NumRows: 5, NumCols: 5, NumBlackHoles: 12, BlackHoleLocator: bhLocator}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.cfg.Validate()
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
