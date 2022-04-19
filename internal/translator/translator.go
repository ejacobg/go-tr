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
	t Translator
	m map[rune]rune
}

func NewReplacer(from, to CharSet) *Replacer {
	if len(to) > len(from) {
		return nil
	}
	r := &Replacer{}
	for i := range from {
		if i >= len(to) {
			r.m[from[i]] = to[len(to)-1]
		} else {
			r.m[from[i]] = to[i]
		}
	}
	return r
}

type Deleter struct {
	t Translator
	m map[rune]struct{}
}

func NewDeleter(cs CharSet) *Deleter {
	d := &Deleter{}
	for _, c := range cs {
		d.m[c] = struct{}{}
	}
	return d
}

type Squeezer struct {
	t Translator
	m map[rune]struct{}
}

func NewSqueezer(cs CharSet) *Squeezer {
	s := &Squeezer{}
	for _, c := range cs {
		s.m[c] = struct{}{}
	}
	return s
}
