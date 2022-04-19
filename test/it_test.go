// Package it contains integration tests.
package it

import "testing"

func TestInputParsing(t *testing.T) {
	t.Run("input is correctly gathered and parsed", func(t *testing.T) {})
}

func TestParsedTranslators(t *testing.T) {
	t.Run("replacer accepts parsed input", func(t *testing.T) {})
	t.Run("deleter accepts parsed input", func(t *testing.T) {})
	t.Run("squeezer accepts parsed input", func(t *testing.T) {})
}

func TestErrorHandling(t *testing.T) {
	t.Run("empty charsets are handled appropriately", func(t *testing.T) {})
}