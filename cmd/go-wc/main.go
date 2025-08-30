package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	SingleValueFormat = "\t%d %s\n"
	DefaultFormat     = "\t%d\t%d\t%d %s\n"
)

func main() {
	flags, subject, filename := parseArgs()
	if flags["bytes"] {
		byteCount := countBytes(subject)
		fmt.Printf(SingleValueFormat, byteCount, filename)
	} else if flags["words"] {
		wordCount := countWords(subject)
		fmt.Printf(SingleValueFormat, wordCount, filename)
	} else if flags["lines"] {
		lineCount := countLines(subject)
		fmt.Printf(SingleValueFormat, lineCount, filename)
	} else if flags["chars"] {
		charCount := countChars(subject)
		fmt.Printf(SingleValueFormat, charCount, filename)
	} else {
		byteCount := countBytes(subject)
		wordCount := countWords(subject)
		lineCount := countLines(subject)
		fmt.Printf(DefaultFormat, lineCount, wordCount, byteCount, filename)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseArgs() (flags map[string]bool, subject []byte, filename string) {
	flags = make(map[string]bool)
	inputFile := os.Stdin
	filename = ""

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			for _, char := range arg[1:] {
				switch char {
				case 'c':
					flags["bytes"] = true
				case 'l':
					flags["lines"] = true
				case 'w':
					flags["words"] = true
				case 'm':
					flags["chars"] = true
				}
			}
		} else {
			filename = arg
			file, err := os.Open(arg)
			check(err)
			inputFile = file
			defer file.Close()
		}
	}

	// Read all data from the file pointer (stdin or file)
	data, err := io.ReadAll(inputFile)
	check(err)
	subject = data

	return flags, subject, filename
}

func countBytes(b []byte) int {
	return len(b)
}

func countWords(b []byte) int {
	return len(strings.Fields(string(b)))
}

func countLines(b []byte) int {
	return bytes.Count(b, []byte{'\n'})
}

func countChars(b []byte) int {
	locale := os.Getenv("LC_CTYPE")
	if locale == "" {
		locale = os.Getenv("LANG")
	}

	// If UTF-8 locale or valid UTF-8 data
	if strings.Contains(strings.ToLower(locale), "utf") || utf8.Valid(b) {
		return utf8.RuneCount(b)
	}

	return len(b)
}
