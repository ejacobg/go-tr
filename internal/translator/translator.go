// Package translator provides operations for performing text substitutions (translations).
package translator

type CharSet struct {
	Chars []rune
	Compl bool
}

type Translator interface {
	Translate(chars []rune) ([]rune, error)
}

type Replacer map[rune]rune

type Deletor map[rune]struct{}

type Squeezer map[rune]struct{}
