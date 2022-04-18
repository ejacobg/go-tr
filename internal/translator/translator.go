// Package translator provides operations for performing text substitutions (translations).
package translator

type CharSet []rune

// Complements the set with respect to the ASCII character set.
// If the character is outside the ASCII range, then it is removed from the new set.
func (cs *CharSet) Complement() (comp CharSet) {
	m := make(map[rune]bool)
	for _, c := range *cs {
		m[c] = true
	}
	for i := rune(0); i < 128; i++ {
		if !m[i] {
			comp = append(comp, i)
		}
	}
	return
}

type Translator interface {
	Translate(chars []rune) ([]rune, error)
}

type Replacer struct {
}

type Deletor map[rune]struct{}

type Squeezer map[rune]struct{}
