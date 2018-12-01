package main

// https://adventofcode.com/2018/day/1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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

// FrequencyCheck will calculate the overall sum of a list of numbers in a file.
func FrequencyCheck(file string) int {
	scanner := MustScanFile(file)
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
		scanner := MustScanFile(file)
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
