package game

import (
	"fmt"
	"errors"
)

const (
	invalidCoordsErrMsg = "There is not a cell with these coordinats."
	busyCoordsErrMsg    = "The cell with these coordinats is busy."
)

const size = 3

const (
	none = iota
	cross
	circle
)

type gameState struct {
	board [size][size]int
	player int
}

func (state *gameState) congratulate(winner int) {
	var winnerName string
	if winner == cross {
		winnerName = "crosses"
	} else {
		winnerName = "circles"
	}

	offset := "				"
	for i, row := range state.board {
		fmt.Print(offset)
		for j, cell := range row {
			switch cell {
				case none: fmt.Print("   ")
				case cross: fmt.Print(" X ")
				case circle: fmt.Print(" O ")
			}

			if j < size - 1 {
				fmt.Printf("|")
			}
		}

		if i == 1 {
			fmt.Printf("	The %v won!", winnerName)
		}

		if i < size - 1 {
			fmt.Printf("\n%v-----------\n", offset)
		}
	}

	fmt.Println("	Press 'Enter', please.")
	fmt.Scanln()
}

func (state *gameState) printBoard() {
	offset := "					"
	for i, row := range state.board {
		fmt.Print(offset)
		for j, cell := range row {
			switch cell {
				case none: fmt.Print("   ")
				case cross: fmt.Print(" X ")
				case circle: fmt.Print(" O ")
			}

			if j < size - 1 {
				fmt.Printf("|")
			}
		}

		if i < size - 1 {
			fmt.Printf("\n%v-----------\n", offset)
		}
	}

	fmt.Println()
}

func (state *gameState) checkCoordinates(x, y int) error {
	if x < 1 || x > size || y < 1 || y > size {
		return errors.New(invalidCoordsErrMsg)
	}

	if state.board[x - 1][y - 1] != none {
		return errors.New(busyCoordsErrMsg)
	}

	return nil
}

func (state *gameState) setMark(x, y int) error {
	err := state.checkCoordinates(x, y)
	if err != nil {
		return err
	}

	state.board[x - 1][y - 1] = state.player
	return nil
}

func (state *gameState) nextTurn() {
	if state.player == cross {
		state.player = circle
	} else {
		state.player = cross
	}
}

func (state *gameState) checkLine(y, x, step_y, step_x int) int {
	var crossCount, circleCount int

	for x >= 0 && x < size && y >= 0 && y < size {
		if state.board[y][x] == cross {
			crossCount++
		} else if state.board[y][x] == circle {
			circleCount++
		}
		x += step_x
		y += step_y
	}

	if crossCount == 3 {
		return cross
	} else if circleCount == 3 {
		return circle
	}

	return none
}

func (state *gameState) checkWinner() int {
	for row := 0; row < size; row++ {
		if result := state.checkLine(row, 0, 0, 1); result != none {
			return result
		}
	}

	for col := 0; col < size; col++ {
		if result := state.checkLine(0, col, 1, 0); result != none {
			return result
		}
	}

	if result := state.checkLine(0, 0, 1, 1); result != none {
		return result
	}

	if result := state.checkLine(0, size - 1, 1, -1); result != none {
		return result
	}

	return none
}

func (state *gameState) haveFreeCell() bool {
	for _, row := range state.board {
		for _, cell := range row {
			if cell == none {
				return true
			}
		}
	}

	return false
}