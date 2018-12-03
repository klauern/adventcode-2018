package main

import (
	"fmt"

	"github.com/klauern/adventcode-2018/helpers"
)

// CalcDistance will calculate the number of occurrances of every letter in the
// given string.
func CalcDistance(id string) map[rune]int {
	count := map[rune]int{}

	for _, runeVal := range id {
		count[runeVal]++
	}
	return count
}

// HasTwoLetter returns whether the given map[rune]int has at least one letter
// with exactly two (2) occurrances of it.
func HasTwoLetter(id map[rune]int) bool {
	for _, v := range id {
		if v == 2 {
			return true
		}
	}
	return false
}

// HasThreeLetter returns whether the given map[rune]int has at least one letter
// with exactly three (3) occurrances of it.
func HasThreeLetter(id map[rune]int) bool {
	for _, v := range id {
		if v == 3 {
			return true
		}
	}
	return false
}

// Checksum will do the instructions form of a "checksum", by multiplying the number of
// 'exactly two letters the same' and number of 'exactly three letters the same' for each
// string in the passed-in array.
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
	scanner := helpers.MustScanFile("input.txt")
	strs := helpers.ToArray(scanner)
	fmt.Printf("checksum: %v\n", Checksum(strs))
	fmt.Printf("Shortest Pair: %#v\n", FindShortestFromFile("input.txt"))
}
