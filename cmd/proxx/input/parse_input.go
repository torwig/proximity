package input

import (
	"errors"
	"fmt"
	"os"
	"proxx/internal/proxx/game"
	"strconv"
	"strings"
)

var (
	ErrEmptyInput        = errors.New("empty input")
	ErrValueIsNotInteger = errors.New("value should be an integer")
	ErrTwoValuesExpected = errors.New("two values should be provided")
)

func GetGameConfig() (game.Config, error) {
	fmt.Println("Please, configure your game.")
	fmt.Println("Enter a number of rows:")

	var in string

	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	rowNum, err := integerFromString(in)
	if err != nil {
		return game.Config{}, fmt.Errorf("failed to get number of rows: %w", err)
	}

	fmt.Println("Enter a number of columns:")

	_, err = fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	colNum, err := integerFromString(in)
	if err != nil {
		return game.Config{}, fmt.Errorf("failed to get number of columns: %w", err)
	}

	fmt.Println("Enter a number of black holes:")

	_, err = fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	bhNum, err := integerFromString(in)
	if err != nil {
		return game.Config{}, fmt.Errorf("failed to get number of black holes: %w", err)
	}

	cfg := game.Config{
		NumRows:          rowNum,
		NumCols:          colNum,
		NumBlackHoles:    bhNum,
		BlackHoleLocator: game.NewUniformBlackHoleLocator(),
	}

	return cfg, nil
}

func integerFromString(in string) (int, error) {
	if exitTheGame(in) {
		os.Exit(0)
	}

	value, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, ErrValueIsNotInteger
	}

	return int(value), nil
}

func GetCellCoordinates() (int, int, error) {
	fmt.Println("Enter the coordinates of a cell you wish to open \n" +
		"in a format \"rowNo,colNo\" (numeration starts from 1) and press ENTER:")

	var in string

	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	if in == "" {
		return 0, 0, ErrEmptyInput
	}

	if exitTheGame(in) {
		os.Exit(0)
	}

	values := strings.Split(in, ",")
	if len(values) != 2 {
		return 0, 0, ErrTwoValuesExpected
	}

	row, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return 0, 0, ErrValueIsNotInteger
	}

	col, err := strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		return 0, 0, ErrValueIsNotInteger
	}

	return int(row - 1), int(col - 1), nil
}

func exitTheGame(in string) bool {
	return in == "Q" || in == "q"
}

func UserWantToPlayAnotherGame() bool {
	fmt.Println("\nDo you want to play another game? Press 'Y/y' to continue:")

	var in string

	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	return in == "Y" || in == "y"
}
