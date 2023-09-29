package board

import (
	"container/list"
	"fmt"
)

// CellValue represents a user-friendly value of a cell (blank/black hole/clue/unknown).
type CellValue string

// Position represents a single position on a game board.
type Position struct {
	Row int
	Col int
}

// Board represents a rectangular set of cells.
type Board struct {
	m [][]*Cell
}

// NewBoard return a new board with the specified number of rows and columns.
// The newly created board filled with only blank cells.
func NewBoard(cfg Config) (*Board, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	matrix := make([][]*Cell, cfg.NumRows)

	for i := 0; i < cfg.NumRows; i++ {
		row := make([]*Cell, cfg.NumCols)

		for j := range row {
			row[j] = newBlankCell()
		}

		matrix[i] = row
	}

	return &Board{m: matrix}, nil
}

// Init initializes the board with black holes and clues.
// the black holes specified by their positions via the bhs.
// Clues are placed according to the number of adjacent black holes.
func (b *Board) Init(bhs []Position) {
	b.populateWithBlackHoles(bhs)
	b.populateWithClues()
}

func (b *Board) populateWithBlackHoles(bhs []Position) {
	for _, bh := range bhs {
		b.putBlackHoleAt(bh.Row, bh.Col)
	}
}

func (b *Board) populateWithClues() {
	for i := 0; i < b.height(); i++ {
		for j := 0; j < b.width(); j++ {
			if b.CellAt(i, j).IsBlackHole() {
				continue
			}

			b.putClueAt(b.numberOfAdjacentBlackHoles(i, j), i, j)
		}
	}
}

// OpenedCells returns the number of opened cells.
func (b *Board) OpenedCells() int {
	var opened int

	for i := 0; i < b.height(); i++ {
		for j := 0; j < b.width(); j++ {
			if b.CellAt(i, j).IsOpen() {
				opened++
			}
		}
	}

	return opened
}

// CellAt return the cell on the board that has the specified coordinates.
func (b *Board) CellAt(row int, col int) *Cell {
	return b.m[row][col]
}

func (b *Board) putBlackHoleAt(i int, j int) {
	b.CellAt(i, j).PutBlackHole()
}

func (b *Board) putClueAt(value int, i int, j int) {
	b.CellAt(i, j).PutClue(value)
}

func (b *Board) width() int {
	return len(b.m[0])
}

func (b *Board) height() int {
	return len(b.m)
}

func (b *Board) numberOfAdjacentBlackHoles(i int, j int) int {
	var n int
	positions := b.getSurroundingCellPositions(i, j)

	for _, p := range positions {
		if c := b.CellAt(p.row, p.col); c.IsBlackHole() {
			n++
		}
	}

	return n
}

// OpenAllBlackHoles marks all the black holes on the board as opened.
func (b *Board) OpenAllBlackHoles() {
	for i := 0; i < b.height(); i++ {
		for j := 0; j < b.width(); j++ {
			if c := b.CellAt(i, j); c.IsBlackHole() {
				c.MarkAsOpen()
			}
		}
	}
}

// MakeCellAndSurroundingCellsOpened marks the specified cell and surrounding cells with clues opened.
// The effect is also applied to all the surrounding blank cells.
func (b *Board) MakeCellAndSurroundingCellsOpened(i, j int) {
	queue := list.New()
	queue.PushBack(cellPos{row: i, col: j})

	for queue.Len() > 0 {
		el := queue.Front()
		pos := el.Value.(cellPos)

		queue.Remove(el)
		c := b.CellAt(pos.row, pos.col)
		c.MarkAsOpen()

		if c.IsBlank() {
			positions := b.getSurroundingCellPositions(pos.row, pos.col)

			for _, p := range positions {
				c = b.CellAt(p.row, p.col)

				if c.IsOpen() {
					continue
				}

				if c.IsBlank() {
					queue.PushBack(cellPos{row: p.row, col: p.col})
				} else {
					c.MarkAsOpen()
				}
			}
		}
	}
}

type cellPos struct {
	row, col int
}

func (b *Board) getSurroundingCellPositions(i, j int) []cellPos {
	neighborPositions := []cellPos{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	result := neighborPositions[:0]

	for _, pos := range neighborPositions {
		if b.ValidCellPosition(pos.row, pos.col) {
			result = append(result, pos)
		}
	}

	return result
}

// ValidCellPosition checks if the cell with the specified coordinates is located on the board.
func (b *Board) ValidCellPosition(row int, col int) bool {
	return row >= 0 && col >= 0 && row <= b.height()-1 && col <= b.width()-1
}

// State returns the current state of the board as a two-dimensional matrix of cell values.
// Reveals only opened cells.
func (b *Board) State() [][]CellValue {
	s := make([][]CellValue, 0, b.height())

	for i := 0; i < b.height(); i++ {
		row := make([]CellValue, 0, b.width())
		for j := 0; j < b.width(); j++ {
			cell := b.CellAt(i, j)

			value := CellValue(CellValueUnknown)
			if cell.IsOpen() {
				value = cell.Value()
			}

			row = append(row, value)
		}
		s = append(s, row)
	}

	return s
}

// DebugState returns the current state of the board as a two-dimensional matrix of cell values.
// Reveal all cells. It's here only for test purposes.
func (b *Board) DebugState() [][]CellValue {
	s := make([][]CellValue, 0, b.height())

	for i := 0; i < b.height(); i++ {
		row := make([]CellValue, 0, b.width())
		for j := 0; j < b.width(); j++ {
			row = append(row, b.CellAt(i, j).Value())
		}
		s = append(s, row)
	}

	return s
}

func (b *Board) TotalNumberOfCells() int {
	return b.width() * b.height()
}
