package translator

import "testing"

func TestComplement(t *testing.T) {
	t.Run("ASCII range is complemented", func(t *testing.T) {
		set := CharSet([]rune{0, 1, 2, 3})
		got := set.Complement()
		for _, c := range got {
			if c < 4 {
				t.Errorf("got %v, want < 4", c)
			}
		}
	})

	t.Run("non-ASCII is ignored", func(t *testing.T) {
		set := CharSet([]rune{'😀'})
		got := set.Complement()
		want := 128
		if len(got) != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}