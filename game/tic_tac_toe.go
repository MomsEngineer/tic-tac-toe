package game

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
)

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

func PrintMenu() {
	cleanScreen()
	fmt.Print(greeting)
	fmt.Print(rules)
	fmt.Print(menu)
}

func Play() {
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
		cleanScreen()
		state.printBoard()

		if winner = state.checkWinner(); winner != none {
			break
		}

		if !state.haveFreeCell() {
			break
		}
	}

	cleanScreen()
	fmt.Print(congratulation)
	state.congratulate(winner)
}

