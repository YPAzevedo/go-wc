package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	flags, subject, filename := parseArgs()
	if flags["bytes"] {
		byteCount := countBytes(subject)
		fmt.Printf("%d %s\n", byteCount, filename)
	} else if flags["words"] {
		wordCount := countWords(subject)
		fmt.Printf("%d %s\n", wordCount, filename)
	} else if flags["lines"] {
		lineCount := countLines(subject)
		fmt.Printf("%d %s\n", lineCount, filename)
	} else if flags["chars"] {
		charCount := countChars(subject)
		fmt.Printf("%d %s\n", charCount, filename)
	} else {
		byteCount := countBytes(subject)
		wordCount := countWords(subject)
		lineCount := countLines(subject)
		fmt.Printf("%d %d %d %s\n", byteCount, wordCount, lineCount, filename)
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
	return len(strings.Split(string(b), "\n"))
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
