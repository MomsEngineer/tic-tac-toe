package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
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
	cleanScreen()
	fmt.Print(congratulation)

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
	cleanScreen()

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

const greeting = `
----------------------------------------------------------------------------------------------------
     _______  _____   _____          _______             _____          _______   ____   ______
    |__   __||_   _| / ____|        |__   __|    /\     / ____|        |__   __| / __ \ |  ____|
       | |     | |  | |      ______    | |      /  \   | |      ______    | |   | |  | || |__
       | |     | |  | |     |______|   | |     / /\ \  | |     |______|   | |   | |  | ||  __|
       | |    _| |_ | |____            | |    / ____ \ | |____            | |   | |__| || |____
       |_|   |_____| \_____|           |_|   /_/    \_\ \_____|           |_|    \____/ |______|

----------------------------------------------------------------------------------------------------
`

const rules = `
Use y and x coordinates for the game. y and x are vertical and horizontal coordinates, respectively.
Cell coordinates:			           Example:
                			Enter y and x, please: 1 1
   1 1|1 2|1 3  			   |   |               X |   |   
   -----------  			-----------           -----------
   2 1|2 2|2 3  			   |   |       ===>      |   |   
   -----------  			-----------           -----------
   3 1|2 3|3 3  			   |   |                 |   |

----------------------------------------------------------------------------------------------------
`

const congratulation = `
----------------------------------------------------------------------------------------------------
   ____   ___   _   _   ____  ____      _     _____  _   _  _        _    _____  ___   ___   _   _ 
  / ___| / _ \ | \ | | / ___||  _ \    / \   |_   _|| | | || |      / \  |_   _||_ _| / _ \ | \ | |
  | |    | | | |   | || |  _ | |_) |  / _ \    | |  | | | || |     / _ \   | |   | | | | | ||  \| |
  | |___ | | | | |\  || |_| ||  _ <  / ___ \   | |  | |_| || |__  / ___ \  | |   | | | |_| || |\  |
  \____| \___/ |_| \_| \____||_| \_\/_/   \_\  |_|   \___/ |____|/_/   \_\ |_|  |___| \___/ |_| \_|

----------------------------------------------------------------------------------------------------
`

const menu = `
						     Menu
						   1. Play
						   2. Exit
Select the menu item: `

func fscanln(b *bufio.Reader, y, x *int) error {
	b.Discard(b.Buffered())
	_, err := fmt.Fscanln(b, y, x)
	return err
}

func cleanScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printMenu() {
	cleanScreen()
	fmt.Print(greeting)
	fmt.Print(rules)
	fmt.Print(menu)
}

func play() {
	state := gameState{player: cross}
	var x, y int
	winner := none

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter y and x coordinates by space, please: ")
		if err := fscanln(stdin, &y, &x); err != nil {
			if err.Error() == "unexpected newline" {
				fmt.Println("Not enough arguments")
			} else if err.Error() == "expected newline" {
				fmt.Println("Too many arguments")
			} else if err.Error() == "expected integer" {
				fmt.Println("Use digits, please")
			} else {
				fmt.Println(err)
			}
			continue
		}

		if err := state.setMark(y, x); err != nil {
			fmt.Println(err.Error())
			continue
		}

		state.nextTurn()
		state.printBoard()

		if winner = state.checkWinner(); winner != none {
			break
		}

		if !state.haveFreeCell() {
			break
		}
	}

	state.congratulate(winner)
}

func main() {
	printMenu()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
			case "1": play()
			case "2": return
		}
		printMenu()
	}
}

