// Package it contains integration tests.
package it

import (
	"control"
	"input"
	"strings"
	"testing"
	"translator"
)

func TestInputParsing(t *testing.T) {
	t.Run("input is correctly gathered and parsed", func(t *testing.T) {
		inp := strings.NewReader("[a-d]")
		var cs translator.CharSet
		cs = input.GetChars(inp)
		cs = control.Parse(string(cs))
		got := string(cs)
		want := "[abcd]"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestParsedTranslators(t *testing.T) {
	inp1 := strings.NewReader("[a-z] ")
	inp2 := strings.NewReader("[A-Z] ")
	var string1, string2 translator.CharSet
	string1, string2 = input.GetChars(inp1), input.GetChars(inp2)
	string1, string2 = control.Parse(string(string1)), control.Parse(string(string2))

	t.Run("replacer accepts parsed input", func(t *testing.T) {
		inp := []rune("the quick brown fox jumps over the lazy dog")
		tr := translator.NewReplacer(string1, string2, nil)
		got := string(tr.Translate(inp))
		want := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("deleter accepts parsed input", func(t *testing.T) {
		inp := []rune("the quick brown fox jumps over the lazy dog")
		d := translator.NewDeleter(string1, nil)
		got := string(d.Translate(inp))
		want := ""
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("squeezer accepts parsed input", func(t *testing.T) {

	})

}

func TestErrorHandling(t *testing.T) {
	t.Run("empty charsets are handled appropriately", func(t *testing.T) {})
}
