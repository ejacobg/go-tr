// Package input generates rune slices from input sources.
package input

import (
	"bufio"
	"io"
)

func GetChars(f io.Reader) (chars []rune) {
	r := bufio.NewReader(f)
	for {
		char, _, err := r.ReadRune()
		if err == io.EOF {
			return
		}
		chars = append(chars, char)
	}
}
