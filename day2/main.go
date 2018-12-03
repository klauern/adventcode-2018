package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// MustScanFile will open a file that we know exists.  If there's a file open error,
// it will panic.
func MustScanFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(errors.WithStack(err))
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func ToArray(input *bufio.Scanner) []string {
	ary := []string{}
	for input.Scan() {
		ary = append(ary, input.Text())
	}

	return ary
}

func CalcDistance(id string) map[rune]int {
	count := map[rune]int{}

	for _, runeVal := range id {
		count[runeVal]++
	}
	return count
}

func HasTwoLetter(id map[rune]int) bool {
	for _, v := range id {
		if v == 2 {
			return true
		}
	}
	return false
}

func HasThreeLetter(id map[rune]int) bool {
	for _, v := range id {
		if v == 3 {
			return true
		}
	}
	return false
}

func Checksum(boxes []string) int {
	twoLetter := 0
	threeLetter := 0
	for _, v := range boxes {
		dist := CalcDistance(v)
		if HasTwoLetter(dist) {
			twoLetter++
		}
		if HasThreeLetter(dist) {
			threeLetter++
		}
	}
	return twoLetter * threeLetter
}

func main() {
	scanner := MustScanFile("input.txt")
	strs := ToArray(scanner)
	fmt.Printf("checksum: %v\n", Checksum(strs))
	fmt.Printf("Shortest Pair: %#v\n", FindShortestFromFile("input.txt"))
}
