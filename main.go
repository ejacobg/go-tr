package main

import (
	"control"
	"flag"
	"fmt"
	"input"
	"os"
	"translator"
)

func init() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Usage of tr:")
		fmt.Fprintln(w, "tr [-c] [-s] string1 string2 [files...]")
		fmt.Fprintln(w, "tr -s [-c] string1 [files...]")
		fmt.Fprintln(w, "tr -d [-c] string1 [files...]")
		fmt.Fprintln(w, "tr -d -s [-c] string1 string2 [files...]")
		fmt.Fprintln(w, "If no files are specified, reads from stdin.")
		flag.PrintDefaults()
	}
}

var (
	complement = flag.Bool("c", false, "Complement the set of values specified by string1.")
	delete     = flag.Bool("d", false, "Delete all occurrences of input characters that are specified by string1.")
	squeeze    = flag.Bool("s", false, "Replace instances of repeated characters with a single character.")
)

func main() {
	flag.Parse()
	fmt.Println(*complement, *delete, *squeeze)
	args := flag.Args()
	//   cds n
	// 0 000 2
	// 1 001 1
	// 2 010 1
	// 3 011 2
	// 4 100 2
	// 5 101 1
	// 6 110 1
	// 7 111 2
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

	fmt.Println(string1, string2)
	var chars []rune
	if len(args) == minArgs {
		chars = input.GetChars(os.Stdin)
		chars = t.Translate(chars)
		fmt.Println(string(chars))
	}
}
