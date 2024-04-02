package main

import (
	"os"
	"bufio"
	"fmt"
	"github.com/MomsEngineer/tic-tac-toe/game"
)

const lang_menu = `
						      Menu
						   1. Русский
						   2. English
Select language: `

func chooseLang() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(lang_menu)
	for scanner.Scan() {
		switch scanner.Text() {
			case "1": return "rus"
			case "2": return "eng"
		}
		fmt.Print(lang_menu)
	}
	return ""
}

func main() {
	lang := chooseLang()
	game.InitializeData(lang)
	game.StartGame()
}