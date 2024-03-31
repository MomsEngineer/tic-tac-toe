package main

import (
	"os"
	"bufio"
	"fmt"
	"github.com/MomsEngineer/tic-tac-toe/game"
)

const lamg_menu = `
						      Menu
						   1. Русский
						   2. English
Select language: `

func choose_lang(scanner *bufio.Scanner) string {
	fmt.Print(lamg_menu)
	for scanner.Scan() {
		switch scanner.Text() {
			case "1": return "rus"
			case "2": return "eng"
		}
		fmt.Print(lamg_menu)
	}
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lang := choose_lang(scanner)
	game.InitializeData(lang)

	game.PrintMenu()
	for scanner.Scan() {
		switch scanner.Text() {
			case "1":game.Play()
			case "2": return
		}
		game.PrintMenu()
	}
}