package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	l := len(args)

	var option string
	var file string
	var input io.Reader

	switch l {
	case 1:
		firstArg := args[0]
		if firstArg == "-c" || firstArg == "-l" || firstArg == "-w" {
			option = firstArg
			input = os.Stdin
		} else {
			file = firstArg
			f, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			input = f
		}
	case 2:
		option = args[0]
		file = args[1]
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		input = f
	}

	switch option {
	// output number of bytes
	case "-c":

		byteCount := countBytes(input)
		fmt.Printf("\t %d %s\n", byteCount, file)
	// output number of lines
	case "-l":

		lineCount := countLines(input)
		fmt.Printf("\t %d %s\n", lineCount, file)

	// output number of words
	case "-w":
		wordCount := countWords(input)
		fmt.Printf("\t %d %s\n", wordCount, file)
	default:

		byteCount := countBytes(input)
		lineCount := countLines(input)
		wordCount := countWords(input)
		fmt.Printf("\t %d %d %d %s\n", lineCount, wordCount, byteCount, file)

	}

}

func countWords(f io.Reader) int {

	v, ok := f.(*os.File)
	if ok {
		v.Seek(0, 0)
		f = v
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func countLines(f io.Reader) int {
	v, ok := f.(*os.File)
	if ok {
		v.Seek(0, 0)
		f = v
	}
	reader := bufio.NewReader(f)
	lineCount := 0
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
	}
	return lineCount
}

func countBytes(f io.Reader) int {
	v, ok := f.(*os.File)
	if ok {
		v.Seek(0, 0)
		f = v
	}

	reader := bufio.NewReader(f)
	byteCount := 0
	for {
		_, err := reader.ReadByte()
		if err != nil {

			break
		}
		byteCount++
	}
	return byteCount
}
