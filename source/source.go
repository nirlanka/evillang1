package source

import (
	"evil/lang1/consts"
	"evil/lang1/state"
	"fmt"
	"strings"
	"unicode"
)

func GetChar() {
	var err error
	state.Look, _, err = state.Reader.ReadRune()
	if err != nil {
		AbortError(err)
	}
}

func AbortCustom(s string) {
	panic("Error: " + s + ".")
}

func AbortError(err error) {
	panic(err)
}

func Expected(s string) {
	AbortCustom(s + " expected")
}

func Match(r rune) {
	if state.Look == r {
		GetChar()
	} else {
		Expected("\"" + string(r) + "\"")
	}
}

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func GetName() {
	if !IsAlpha(state.Look) {
		Expected("Name")
	}

	// test
	state.GetName = strings.ToUpper(string(state.Look))
	GetChar()
}

func Write(a ...any) {
	fmt.Print(a...)
}

func Emit(s string) {
	Write(consts.Tab, s)
}

func EmitLn(s string) {
	Emit(s)
	Write('\n')
}
