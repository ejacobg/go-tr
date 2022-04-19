// Package control parses the control strings from the command line arguments.
package control

func isASCII(c rune) bool {
	return c < 128
}

// Takes two endpoints, and returns all the characters in between them (inclusive).
// Returns nil if endpoints are non-ASCII characters.
// Returns nil if end precedes start.
func expand(start, end rune) (expansion []rune) {
	if !isASCII(start) || !isASCII(end) {
		return
	}
	if end < start {
		return
	}
	for start <= end {
		expansion = append(expansion, start)
		start++
	}
	return
}

// Generates the character array from a control string.
func Parse(control string) (arr []rune) {
	runes := []rune(control)
	for i := 0; i < len(runes); i++ {
		switch c := runes[i]; c {
		case '-':
			if i == 0 || i == len(runes)-1 {
				arr = append(arr, c)
			} else {
				arr = append(arr[:i-1], expand(runes[i-1], runes[i+1])...)
				i++
			}
		default:
			arr = append(arr, c)
		}
	}
	return
}
