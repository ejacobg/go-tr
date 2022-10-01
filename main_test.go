// Integration tests.
package main

import (
	"strings"
	"testing"

	"github.com/ejacobg/go-tr/control"
	"github.com/ejacobg/go-tr/input"
	"github.com/ejacobg/go-tr/translator"
)

func TestInputParsing(t *testing.T) {
	t.Run(
		"input is correctly gathered and parsed", func(t *testing.T) {
			inp := strings.NewReader("[a-d]")
			var cs translator.CharSet
			cs = input.GetChars(inp)
			cs = control.Parse(string(cs))
			got := string(cs)
			want := "[abcd]"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)
}

func TestParsedTranslators(t *testing.T) {
	inp1 := strings.NewReader("[a-z] ")
	inp2 := strings.NewReader("[A-Z] ")
	var string1, string2 translator.CharSet
	string1, string2 = input.GetChars(inp1), input.GetChars(inp2)
	string1, string2 = control.Parse(string(string1)), control.Parse(string(string2))

	t.Run(
		"replacer accepts parsed input", func(t *testing.T) {
			inp := []rune("the quick brown fox jumps over the lazy dog")
			tr := translator.NewReplacer(string1, string2, nil)
			got := string(tr.Translate(inp))
			want := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"deleter accepts parsed input", func(t *testing.T) {
			inp := []rune("the quick brown fox jumps over the lazy dog")
			d := translator.NewDeleter(string1, nil)
			got := string(d.Translate(inp))
			want := ""
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"squeezer accepts parsed input", func(t *testing.T) {
			inp := []rune("aaaaa")
			s := translator.NewSqueezer(string1, nil)
			got := string(s.Translate(inp))
			want := "a"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

}

func TestErrorHandling(t *testing.T) {
	empty := strings.NewReader("")
	nonempty := strings.NewReader("[a-z] ")
	var string1, string2 translator.CharSet
	string1, string2 = input.GetChars(empty), input.GetChars(nonempty)
	string1, string2 = control.Parse(string(string1)), control.Parse(string(string2))

	t.Run(
		"replacer handles empty charsets", func(t *testing.T) {
			inp := []rune("nothing is changed")
			tr := translator.NewReplacer(string1, string1, nil)
			got := string(tr.Translate(inp))
			want := "nothing is changed"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"deleter handles empty charsets", func(t *testing.T) {
			inp := []rune("nothing is changed")
			d := translator.NewDeleter(string1, nil)
			got := string(d.Translate(inp))
			want := "nothing is changed"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"squeezer handles empty charsets", func(t *testing.T) {
			inp := []rune("aaaa bbbb")
			s := translator.NewSqueezer(string1, nil)
			got := string(s.Translate(inp))
			want := "aaaa bbbb"
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)
}
