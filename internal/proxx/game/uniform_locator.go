package game

import (
	"math/rand"
	"proxx/internal/proxx/board"
	"time"
)

// UniformBlackHoleLocator uniformly distributes black holes across a game board.
// Each cell has a 50% chance to be chosen as a place for a black hole.
type UniformBlackHoleLocator struct {
	rg *rand.Rand
}

// NewUniformBlackHoleLocator returns new UniformBlackHoleLocator object.
func NewUniformBlackHoleLocator() *UniformBlackHoleLocator {
	return &UniformBlackHoleLocator{rg: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// LocateBlackHolesOnBoard returns positions of black holes on a game board
// with the specified number of rows and columns. The method must locate exactly bhNum black holes.
// Game's configuration is checked to guarantee that this function receives valid values.
func (l UniformBlackHoleLocator) LocateBlackHolesOnBoard(rows int, cols int, bhNum int) []board.Position {
	positions := make([]board.Position, 0, bhNum)

	alreadyOccupied := make(map[board.Position]struct{}, bhNum)

	for len(positions) < bhNum {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				pos := board.Position{Row: i, Col: j}

				if _, occupied := alreadyOccupied[pos]; occupied {
					continue
				}

				if v := l.rg.Intn(2); v == 1 {
					positions = append(positions, pos)

					if len(positions) == bhNum {
						return positions
					}

					alreadyOccupied[pos] = struct{}{}
				}
			}
		}
	}

	return positions
}
