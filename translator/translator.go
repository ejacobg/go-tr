// Package translator provides operations for performing text substitutions (translations).
package translator

type CharSet []rune

// Complement complements the set with respect to the ASCII character set.
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
	Translate(chars []rune) []rune
}

type Replacer struct {
	t Translator
	m map[rune]rune
}

func NewReplacer(from, to CharSet, t Translator) *Replacer {
	if len(to) > len(from) {
		return nil
	}
	r := &Replacer{t: t, m: make(map[rune]rune)}
	if len(to) == 0 {
		return r
	}
	for i := range from {
		if i >= len(to) {
			r.m[from[i]] = to[len(to)-1]
		} else {
			r.m[from[i]] = to[i]
		}
	}
	return r
}

func (r *Replacer) Translate(chars []rune) (translated []rune) {
	if r.t != nil {
		chars = r.t.Translate(chars)
	}
	for _, c := range chars {
		if val, prs := r.m[c]; prs {
			translated = append(translated, val)
		} else {
			translated = append(translated, c)
		}
	}
	return
}

type Deleter struct {
	t Translator
	m map[rune]struct{}
}

func NewDeleter(cs CharSet, t Translator) *Deleter {
	d := &Deleter{t: t, m: make(map[rune]struct{})}
	for _, c := range cs {
		d.m[c] = struct{}{}
	}
	return d
}

func (d *Deleter) Translate(chars []rune) (translated []rune) {
	if d.t != nil {
		chars = d.t.Translate(chars)
	}
	for _, c := range chars {
		if _, prs := d.m[c]; !prs {
			translated = append(translated, c)
		}
	}
	return
}

type Squeezer struct {
	t Translator
	m map[rune]struct{}
}

func NewSqueezer(cs CharSet, t Translator) *Squeezer {
	s := &Squeezer{t: t, m: make(map[rune]struct{})}
	for _, c := range cs {
		s.m[c] = struct{}{}
	}
	return s
}

func (s *Squeezer) Translate(chars []rune) (translated []rune) {
	if s.t != nil {
		chars = s.t.Translate(chars)
	}
	for i := 0; i < len(chars); i++ {
		translated = append(translated, chars[i])
		if _, prs := s.m[chars[i]]; prs {
			c := chars[i]
			for i+1 < len(chars) && chars[i+1] == c {
				i++
			}
		}
	}
	return translated
}
