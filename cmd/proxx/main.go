package main

import (
	"fmt"
	"proxx/cmd/proxx/input"
	"proxx/internal/proxx/board"
	"proxx/internal/proxx/game"
	"strings"
)

func main() {
	fmt.Println("Press Q/q to leave the game.")

	for {
		fmt.Printf("\n\n")

		gameCfg, err := input.GetGameConfig()
		if err != nil {
			fmt.Printf("Invalid game configuration: %s", err)
			continue
		}

		proxx, err := game.NewGame(gameCfg)
		if err != nil {
			fmt.Printf("Failed to create a new game: %s", err)
			continue
		}

		fmt.Println("Your game is ready!")

		for !proxx.IsOver() {
			showBoardState(proxx.BoardState())

			row, col, err := input.GetCellCoordinates()
			if err != nil {
				fmt.Printf("Failed to parse cell coordinates: %s", err)
				continue
			}

			if err := proxx.OpenCell(row, col); err != nil {
				fmt.Printf("Error opening the cell: %s", err)
			}
		}

		if proxx.IsOver() {
			if proxx.IsWon() {
				fmt.Println("Great job, champion!")
			} else {
				fmt.Println("Oops! This time a Black Hole captured you!")
			}
		}

		showBoardState(proxx.BoardState())

		if !input.UserWantToPlayAnotherGame() {
			break
		}
	}

	fmt.Println("Bye!")
}

func showBoardState(bs [][]board.CellValue) {
	fmt.Println()
	fmt.Println("Current state of the board:")

	builder := strings.Builder{}

	for _, row := range bs {
		for _, v := range row {
			builder.WriteString(fmt.Sprintf("%s\t", v))
		}
		builder.WriteString("\n\n")
	}

	fmt.Println(builder.String())
}
