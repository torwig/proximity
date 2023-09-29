package game

import "errors"

var (
	ErrTooManyBlackHoles           = errors.New("too many black holes: at least one cell should be free from them")
	ErrNoBlackHolesProvided        = errors.New("too few black holes: at least one black hole required")
	ErrBlackHoleLocatorNotProvided = errors.New("black hole locator wasn't provided")
)

// Config represents a configuration for the Proxx game.
// BlackHoleLocator is used to assign black holes to cells across a game board.
// Currently, the UniformBlackHoleLocator type is provided to uniformly distribute black holes.
// At least one cell should be left for a clue to successfully create a new game.
// At least 1 black hole should be present to successfully create a new game.
type Config struct {
	NumRows          int
	NumCols          int
	NumBlackHoles    int
	BlackHoleLocator BlackHoleLocator
}

func (cfg Config) Validate() error {
	if cfg.NumBlackHoles >= cfg.NumRows*cfg.NumCols {
		return ErrTooManyBlackHoles
	}

	if cfg.NumBlackHoles < 1 {
		return ErrNoBlackHolesProvided
	}

	if cfg.BlackHoleLocator == nil {
		return ErrBlackHoleLocatorNotProvided
	}

	return nil
}
