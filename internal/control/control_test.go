package control

import (
	"fmt"
	"testing"
)

func TestASCII(t *testing.T) {
	tests := []rune{
		0,   // First ASCII character
		48,  // 0
		57,  // 9
		65,  // A
		90,  // Z
		97,  // a
		122, // z
		127, // Last ASCII character
		128, // First UTF-8 character
		'\U0001f600', // 😀
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%U", test), func(t *testing.T) {
			got := isASCII(test)
			want := index <= 7
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestExpansion(t *testing.T) {
	t.Run("Valid range", func(t *testing.T) {
		got := string(expand('a', 'd'))
		want := "abcd"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Invalid range", func(t *testing.T) {
		got := expand('d', 'a')
		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})

	t.Run("Non-ASCII range", func(t *testing.T) {
		got := expand('a', '\U0001f600')
		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
}