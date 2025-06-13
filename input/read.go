package input

import (
	"bufio"
	"io"
	"os"
)

var reader *bufio.Reader
var file *os.File

func Setup() {
	var input io.Reader

	if len(os.Args) > 1 {
		var err error
		file, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		input = file
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

func Close() {
	file.Close()
}
