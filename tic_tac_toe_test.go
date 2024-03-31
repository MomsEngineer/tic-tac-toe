package main

import (
	"testing"
	"errors"
	"io"
	"os"
	"reflect"
	"bufio"
	"bytes"
)

func Test_checkCoordinates(t *testing.T) {
	var state gameState
	state.board[1][1] = cross

	t.Run("Good case", func(t *testing.T) {
		err := state.checkCoordinates(1, 1)
		if err != nil {
			t.Errorf("incorrect result: expected %v, got %v", nil, err)
		}
	})

	data := []struct {
		name string
		y int
		x int
		expected error
	} {
		{"Zero coords",      0,  0, errors.New(invalidCoordsErrMsg)},
		{"Negative coords", -1, -1, errors.New(invalidCoordsErrMsg)},
		{"Too big coords",   5,  5, errors.New(invalidCoordsErrMsg)},
		{"Busy cell",        2,  2, errors.New(busyCoordsErrMsg)},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			err := state.checkCoordinates(d.y, d.x)
			if err == nil || err.Error() != d.expected.Error() {
				t.Errorf("incorrect result: expected %v, got %v", d.expected, err)
			}
		})
	}
}

func Test_setMark(t *testing.T) {
	state := gameState{player: cross}

	t.Run("Good case", func(t *testing.T) {
		err := state.setMark(1, 1)
		if err != nil {
			t.Errorf("incorrect result: expected %v, got %v", nil, err)
		}
	})

	t.Run("Busy cell", func(t *testing.T) {
		err := state.setMark(1, 1)
		if err == nil || err.Error() != busyCoordsErrMsg {
			t.Errorf("incorrect result: expected %v, got %v", busyCoordsErrMsg, err)
		}
	})
}

func Test_nextTurn(t *testing.T) {
	state := gameState{player: cross}
	
	state.nextTurn()
	if state.player != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, state.player)
	}

	state.nextTurn()
	if state.player != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, state.player)
	}
}

func Test_checkLine_Horizontal(t *testing.T) {
	var state gameState
	state.board[0] = [size]int{cross, cross, cross}

	result := state.checkLine(0, 0, 0, 1)
	if result != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, result)
	}

	state.board[1] = [size]int{circle, circle, circle}

	result = state.checkLine(1, 0, 0, 1)
	if result != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, result)
	}
}

func Test_checkLine_Vertical(t *testing.T) {
	var state gameState

	for i := 0; i < size; i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				state.board[i][j] = cross
			} else {
				state.board[i][j] = circle
			}
		}
	}

	result := state.checkLine(0, 0, 1, 0)
	if result != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, result)
	}

	result = state.checkLine(0, 1, 1, 0)
	if result != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, result)
	}
}

func Test_checkLine_Diagonal(t *testing.T) {
	var state gameState

	for i := 0; i < size; i++ {
		state.board[i][i] = cross
	}

	result := state.checkLine(0, 0, 1, 1)
	if result != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, result)
	}

	for i := 0; i < size; i++ {
		state.board[i][size - i - 1] = circle
	}

	result = state.checkLine(0, 2, 1, -1)
	if result != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, result)
	}
}

func Test_checkLine_None(t *testing.T) {
	var state gameState

	result := state.checkLine(0, 0, 1, 0)
	if result != none {
		t.Errorf("incorrect result: expected %v, got %v", none, result)
	}
}

func Test_checkWinner_Horizontal(t *testing.T) {
	var state gameState
	state.board[1] = [size]int{cross, cross, cross}

	result := state.checkWinner()
	if result != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, result)
	}
}

func Test_checkWinner_Vertical(t *testing.T) {
	var state gameState

	for i := 0; i < size; i++ {
		state.board[i][1] = circle
	}

	result := state.checkWinner()
	if result != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, result)
	}
}

func Test_checkWinner_MainDiagonal(t *testing.T) {
	var state gameState

	for i := 0; i < size; i++ {
		state.board[i][i] = cross
	}

	result := state.checkWinner()
	if result != cross {
		t.Errorf("incorrect result: expected %v, got %v", cross, result)
	}
}

func Test_checkWinner_CrossDiagonal(t *testing.T) {
	var state gameState

	for i := 0; i < size; i++ {
		state.board[i][size - i - 1] = circle
	}

	result := state.checkWinner()
	if result != circle {
		t.Errorf("incorrect result: expected %v, got %v", circle, result)
	}
}

func Test_checkWinner_None(t *testing.T) {
	var state gameState

	result := state.checkWinner()
	if result != none {
		t.Errorf("incorrect result: expected %v, got %v", none, result)
	}
}

func Test_haveFreeCell(t *testing.T) {
	state := gameState {
		board: [size][size]int {
			{cross, cross, cross},
			{cross, none, cross},
			{cross, cross, cross},
		},
	}
	result := state.haveFreeCell()
	if result != true {
		t.Errorf("incorrect result: expected %v, got %v", true, result)
	}

	state.board[1][1] = circle
	result = state.haveFreeCell()
	if result != false {
		t.Errorf("incorrect result: expected %v, got %v", false, result)
	}
}

func getDataFromConsole(print func()) []byte {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	print()

	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)
	return out[1:]
}

func Test_printBoard(t *testing.T) {
	state := gameState{
		board: [3][3]int{
			{1, 0, 2},
			{0, 1, 0},
			{2, 2, 1},
		},
	}

	result := getDataFromConsole(func () {state.printBoard()})

	tmp := `					 X |   | O 
					-----------
					   | X |   
					-----------
					 O | O | X 
	`
	expected := []byte(tmp[:len(tmp)-1])

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
}

func Test_congratulate(t *testing.T) {
	state := gameState{
		board: [3][3]int{
			{1, 0, 2},
			{0, 1, 0},
			{2, 2, 1},
		},
	}

	result := getDataFromConsole(func () {state.congratulate(1)})

	tmp := `				 X |   | O 
				-----------
				   | X |   	The crosses won!
				-----------
				 O | O | X 	Press 'Enter', please.
	`
	expected := []byte(congratulation + tmp[:len(tmp)-1])

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("incorrect result: expected %v, got %v", expected, result)
	}
}

func Test_fscanln(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		var y, x int
		stdin := bufio.NewReader(bytes.NewBufferString("9 11"))
		err := fscanln(stdin, &y, &x)
		if err != nil {
			t.Errorf("expected error %v, got %v", nil, err)
		}

		if y != 9 || x != 11 {
			t.Errorf("expected (9, 11), got (%d, %d)", y, x)
		}
	})

	data := []struct {
		name string
		input string
		expectedY int
		expectedX int
		expectedErr error
	}{
		{name: "Input letters", input: "a d", expectedY: 0, expectedX: 0, expectedErr: errors.New("expected integer")},
		{name: "Input special characters", input: "- @", expectedY: 0, expectedX: 0, expectedErr: errors.New("expected integer")},
		{name: "Incomplete input", input: "1", expectedY: 0, expectedX: 0, expectedErr: errors.New("EOF")},
		{name: "Empty input", input: "", expectedY: 0, expectedX: 0, expectedErr: errors.New("EOF")},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			var y, x int
			stdin := bufio.NewReader(bytes.NewBufferString(d.input))
			err := fscanln(stdin, &y, &x)
			if err != nil && err.Error() != d.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", d.expectedErr, err)
			}
		})
	}
}