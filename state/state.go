package state

import (
	"bufio"
	"os"
)

var Look rune
var GetName string

var Reader = bufio.NewReader(os.Stdin)
