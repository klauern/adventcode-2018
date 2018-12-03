package main

// https://adventofcode.com/2018/day/1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/klauern/adventcode-2018/helpers"
)

// FrequencyCheck will calculate the overall sum of a list of numbers in a file.
func FrequencyCheck(file string) int {
	scanner := helpers.MustScanFile(file)
	frequency := 0
	for scanner.Scan() {
		chg, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(errors.WithMessage(err, "error reading next line from input"))
		}
		frequency += chg
	}
	return frequency
}

// FindDuplicateFreq will find the first number that occurs in the frequency twice.
func FindDuplicateFreq(file string) int {
	current := 0
	freqMap := map[int]bool{}
	foundDuplicate := false
	for !foundDuplicate {
		scanner := helpers.MustScanFile(file)
		for scanner.Scan() {
			chg, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(errors.WithMessage(err, "error reading next line from input"))
			}
			current += chg
			if freqMap[current] {
				foundDuplicate = true
				return current
			}
			freqMap[current] = true
		}
	}
	return current
}

func main() {
	fmt.Printf("Total Frequency (part 1): %v\n", FrequencyCheck("input.txt"))
	fmt.Printf("Duplicate Frequency (part 2): %v\n", FindDuplicateFreq("input.txt"))
}
