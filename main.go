package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ejacobg/go-tr/control"
	"github.com/ejacobg/go-tr/input"
	"github.com/ejacobg/go-tr/translator"
)

func init() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Usage of tr:")
		fmt.Fprintln(w, "tr [-c] [-s] string1 string2")
		fmt.Fprintln(w, "tr -s [-c] string1")
		fmt.Fprintln(w, "tr -d [-c] string1")
		fmt.Fprintln(w, "tr -d -s [-c] string1 string2")
		fmt.Fprintln(w, "After giving arguments, program reads from stdin.")
		fmt.Fprintln(w, "To submit input, pass in the EOF character.")
		flag.PrintDefaults()
	}
}

var (
	complement = flag.Bool("c", false, "Complement the set of values specified by string1.")
	delete     = flag.Bool(
		"d", false, "Delete all occurrences of input characters that are specified by string1.",
	)
	squeeze = flag.Bool(
		"s", false, "Replace instances of repeated characters with a single character.",
	)
)

func main() {
	flag.Parse()
	args := flag.Args()
	// Only need 2 args if none or both of delete/squeeze flags are set.
	minArgs := 1
	if *delete == *squeeze {
		minArgs = 2
	}
	if len(args) < minArgs {
		fmt.Println("Incorrect number of arguments.")
		os.Exit(1)
	}

	var string1, string2 translator.CharSet
	string1 = control.Parse(args[0])
	if *delete == *squeeze {
		string2 = control.Parse(args[1])
	}
	if *complement {
		string1 = string1.Complement()
	}

	if len(string2) > len(string1) {
		fmt.Println("string2 is too long.")
		os.Exit(1)
	}

	var t translator.Translator
	switch {
	case *delete && *squeeze:
		t = translator.NewDeleter(string1, t)
		t = translator.NewSqueezer(string2, t)
	case *delete:
		t = translator.NewDeleter(string1, t)
	case *squeeze:
		t = translator.NewSqueezer(string1, t)
	default:
		t = translator.NewReplacer(string1, string2, t)
	}

	var chars []rune
	chars = input.GetChars(os.Stdin)
	chars = t.Translate(chars)
	fmt.Println(string(chars))
}
