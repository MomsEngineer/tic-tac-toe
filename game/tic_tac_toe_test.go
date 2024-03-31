package game

import (
	"testing"
	"errors"
	"bufio"
	"bytes"
)

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