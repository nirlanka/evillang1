package input

import (
	"bufio"
	"io"
	"os"
)

var reader *bufio.Reader

func Setup() {
	var input io.Reader

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		input = f
	} else {
		input = os.Stdin
	}

	reader = bufio.NewReader(input)
}

func ReadRune() (rune, error) {
	ch, _, err := reader.ReadRune()

	if err == io.EOF {
		return -1, nil // EOF
	}

	return ch, err
}
