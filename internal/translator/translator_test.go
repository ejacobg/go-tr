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

func TestReplacer(t *testing.T) {
	t.Run("replacer does not accept longer 'to' field", func(t *testing.T) {
		from, to := CharSet([]rune("a")), CharSet([]rune("ab"))
		got := NewReplacer(from, to, nil)
		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})

	t.Run("charsets same length", func(t *testing.T) {
		from, to := CharSet([]rune("abcd")), CharSet([]rune("efgh"))
		r := NewReplacer(from, to, nil)
		got := string(r.Translate([]rune("abcd")))
		want := "efgh"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("charsets different length", func(t *testing.T) {
		from, to := CharSet([]rune("abcd")), CharSet([]rune("e"))
		r := NewReplacer(from, to, nil)
		got := string(r.Translate([]rune("abcd")))
		want := "eeee"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// test empty 'from' and 'to' sets?
}

func TestDeleter(t *testing.T) {
	t.Run("deleter deletes characters", func(t *testing.T) {})
}

func TestSqueezer(t *testing.T) {
	t.Run("squeezer squeezes characters", func(t *testing.T) {})
}

func TestDecorator(t *testing.T) {}