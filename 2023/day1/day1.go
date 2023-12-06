package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer file.Close()

	value, err := Calibration(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Println(value)

}

func Calibration(input io.Reader) (int, error) {
	s := bufio.NewScanner(input)

	var sum int

	for s.Scan() {
		line := []rune(s.Text())

		var first, last rune

		lineLen := len(line) - 1
		for i := range line {
			f := line[i]
			l := line[lineLen-i]
			if first == 0 && unicode.IsDigit(f) {
				first = f
			}
			if last == 0 && unicode.IsDigit(l) {
				last = l
			}
			if first != 0 && last != 0 {
				break
			}
		}

		numStr := string(first) + string(last)

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		sum += num
	}
	if s.Err() != nil {
		return 0, s.Err()
	}

	return sum, nil
}
