package board

import "strconv"

const (
	valueBlackHole     = -1
	valueBlank         = 0
	CellValueBlackHole = "H"
	CellValueBlank     = "0"
	CellValueUnknown   = "?"
)

type Cell struct {
	value  int
	isOpen bool
}

func newBlankCell() *Cell {
	return &Cell{value: valueBlank}
}

func (c *Cell) IsBlank() bool {
	return c.value == valueBlank
}

func (c *Cell) PutBlackHole() {
	c.value = valueBlackHole
}

func (c *Cell) IsBlackHole() bool {
	return c.value == valueBlackHole
}

func (c *Cell) MarkAsOpen() {
	c.isOpen = true
}

func (c *Cell) IsOpen() bool {
	return c.isOpen
}

func (c *Cell) PutClue(value int) {
	c.value = value
}

func (c *Cell) IsClue() bool {
	return c.value > valueBlank
}

func (c *Cell) GetClue() int {
	return c.value
}

func (c *Cell) Value() CellValue {
	if c.IsBlackHole() {
		return CellValueBlackHole
	}

	if c.IsBlank() {
		return CellValueBlank
	}

	return CellValue(strconv.FormatInt(int64(c.value), 10))
}
