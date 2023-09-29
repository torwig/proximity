package board

import "errors"

var (
	ErrInvalidNumberOfRows    = errors.New("invalid number of rows")
	ErrInvalidNumberOfColumns = errors.New("invalid number of columns")
)

// Config represents a configuration for a game board.
type Config struct {
	NumRows int
	NumCols int
}

// validate checks the board's configuration.
func (cfg Config) validate() error {
	if cfg.NumRows < 1 {
		return ErrInvalidNumberOfRows
	}

	if cfg.NumCols < 1 {
		return ErrInvalidNumberOfColumns
	}

	return nil
}
