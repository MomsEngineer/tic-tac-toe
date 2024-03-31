package main

import (
	"os"
	"bufio"
	"github.com/MomsEngineer/tic-tac-toe/game"
)

func main() {
	game.PrintMenu()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
			case "1": game.Play()
			case "2": return
		}
		game.PrintMenu()
	}
}