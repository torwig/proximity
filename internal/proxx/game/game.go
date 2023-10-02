package game

import (
	"container/list"
	"errors"
	"fmt"
	"proxx/internal/proxx/board"
)

var (
	ErrCellPositionIsOutsideBoard = errors.New("cell position is out of the board's bounds")
	ErrNumberOfBlackHolesMismatch = errors.New("locator yields different number of black holes than specified via configuration")
)

// Game represents a game.
type Game struct {
	board  *board.Board
	cfg    Config
	isLost bool
	isWon  bool
}

// BlackHoleLocator is the interface that wraps the LocateBlackHolesOnBoard method.
//
// LocateBlackHolesOnBoard returns positions of black holes on a game board
// with the specified number of rows and columns. The method must locate exactly bhNum black holes.
// Game's configuration is checked to guarantee that this function receives valid values.
type BlackHoleLocator interface {
	LocateBlackHolesOnBoard(rows int, cols int, bhNum int) []board.Position
}

// NewGame creates a new game using the specified configuration.
func NewGame(cfg Config) (*Game, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	blackHoles := cfg.BlackHoleLocator.LocateBlackHolesOnBoard(cfg.NumRows, cfg.NumCols, cfg.NumBlackHoles)

	if len(blackHoles) != cfg.NumBlackHoles {
		return nil, ErrNumberOfBlackHolesMismatch
	}

	gameBoard, err := board.NewBoard(board.Config{NumRows: cfg.NumRows, NumCols: cfg.NumCols})
	if err != nil {
		return nil, fmt.Errorf("failed to create a board: %w", err)
	}

	gameBoard.Init(blackHoles)

	return &Game{board: gameBoard, cfg: cfg}, nil
}

// IsOver checks whether a game is over (won or lost).
func (g *Game) IsOver() bool {
	return g.isLost || g.isWon
}

// IsWon return true if a player won the game.
func (g *Game) IsWon() bool {
	return g.isWon
}

// OpenCell opens the specified cell. Returns an error if the position isn't within the board.
func (g *Game) OpenCell(row int, col int) error {
	if !g.board.ValidCellPosition(row, col) {
		return ErrCellPositionIsOutsideBoard
	}

	if g.IsOver() {
		return nil
	}

	cell := g.board.CellAt(row, col)

	if cell.IsOpen() {
		return nil
	}

	if cell.IsBlackHole() {
		g.board.OpenAllBlackHoles()
		g.isLost = true
		return nil
	}

	if cell.IsClue() {
		cell.MarkAsOpen()
	}

	if cell.IsBlank() {
		g.makeCellAndSurroundingCellsOpened(row, col)
	}

	g.isWon = g.board.OpenedCells() == g.maxOpenCells()

	return nil
}

// makeCellAndSurroundingCellsOpened marks the specified cell and surrounding cells with clues opened.
// The effect is also applied to all the surrounding blank cells.
func (g *Game) makeCellAndSurroundingCellsOpened(row int, col int) {
	queue := list.New()
	queue.PushBack(board.Position{Row: row, Col: col})

	for queue.Len() > 0 {
		el := queue.Front()
		pos := el.Value.(board.Position)

		queue.Remove(el)
		c := g.board.CellAt(pos.Row, pos.Col)
		c.MarkAsOpen()

		if c.IsBlank() {
			positions := g.board.GetSurroundingCellPositions(pos.Row, pos.Col)

			for _, p := range positions {
				c = g.board.CellAt(p.Row, p.Col)

				if c.IsOpen() {
					continue
				}

				if c.IsBlank() {
					queue.PushBack(p)
				} else {
					c.MarkAsOpen()
				}
			}
		}
	}
}

func (g *Game) maxOpenCells() int {
	return g.board.TotalNumberOfCells() - g.cfg.NumBlackHoles
}

// BoardState return the current state of a game board.
func (g *Game) BoardState() [][]board.CellValue {
	return g.board.State()
}
