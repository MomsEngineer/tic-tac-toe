package game

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
)


type basicMessages struct {
	greeting string
	rules string
	congratulation string
	menu string
	inputPrompt string
	pressEnter string
	winMsg string
	сrosses string
	circles string
}

type errorMessages struct {
	invalidCoords string
	busyCoords string
	noEnoughArgs string
	tooManyArgs string
	useDigits string
}

type messages struct {
	basic basicMessages
	err errorMessages
}

var msg messages

func fscanln(b *bufio.Reader, y, x *int) error {
	b.Discard(b.Buffered())
	_, err := fmt.Fscanln(b, y, x)
	return err
}

func InitializeData(lang string) {
	if lang == "rus" {
		msg = messages {
			basic: basicMessages {
				greeting: greeting_rus,
				rules: rules_rus,
				congratulation: congratulation_rus,
				menu: menu_rus,
				inputPrompt: inputPrompt_rus,
				pressEnter: pressEnter_rus,
				winMsg: winMsg_rus,
				сrosses: сrosses_rus,
				circles: circles_rus,
			},
			err: errorMessages {
				invalidCoords: invalidCoords_rus,
				busyCoords: busyCoords_rus,
				noEnoughArgs: noEnoughArgs_rus,
				tooManyArgs: tooManyArgs_rus,
				useDigits: useDigits_rus,
			},
		}
	} else {
		msg = messages {
			basic: basicMessages {
				greeting: greeting_eng,
				rules: rules_eng,
				congratulation: congratulation_eng,
				menu: menu_eng,
				inputPrompt: inputPrompt_eng,
				pressEnter: pressEnter_eng,
				winMsg: winMsg_eng,
				сrosses: сrosses_eng,
				circles: circles_eng,
			},
			err: errorMessages {
				invalidCoords: invalidCoords_eng,
				busyCoords: busyCoords_eng,
				noEnoughArgs: noEnoughArgs_eng,
				tooManyArgs: tooManyArgs_eng,
				useDigits: useDigits_eng,
			},
		}
	}
}

func cleanScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintMenu() {
	cleanScreen()
	fmt.Print(msg.basic.greeting)
	fmt.Print(msg.basic.rules)
	fmt.Print(msg.basic.menu)
}

func Play() {
	state := gameState{player: cross}
	var x, y int
	winner := none

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(msg.basic.inputPrompt)
		if err := fscanln(stdin, &y, &x); err != nil {
			if err.Error() == "unexpected newline" {
				fmt.Println(msg.err.noEnoughArgs)
			} else if err.Error() == "expected newline" {
				fmt.Println(msg.err.tooManyArgs)
			} else if err.Error() == "expected integer" {
				fmt.Println(msg.err.useDigits)
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
	fmt.Print(msg.basic.congratulation)
	state.congratulate(winner)
}

