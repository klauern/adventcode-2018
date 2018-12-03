package main

import (
	"strings"
)

type pairDist struct {
	first    int
	second   int
	distance int
	common   string
}

// LevenshteinDistance will calculate--you guessed it--the Levenshtein Distance
// between two strings.  It's copied nearly verbatim from this link:
// https://www.golangprograms.com/data-structure-and-algorithms/golang-program-for-implementation-of-levenshtein-distance.html
func LevenshteinDistance(first, second string) int {
	s1len := len(first)
	s2len := len(second)

	column := make([]int, len(first)+1)
	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if first[y-1] != second[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

// CommonChars will return the substring from the Levenshtein distance,
// removing all of the different characters in the string, returning a
// "common" substring.
func CommonChars(first, second string) string {
	common := strings.Builder{}
	for i := range first {
		if first[i] == second[i] {
			common.WriteRune(rune(first[i]))
		}
	}
	return common.String()
}

// FindShortestPair will take an array of strings and find the
// pair that is the shortest Levenshtein distance apart, returning a
// pairDist struct.
func FindShortestPair(blocks []string) pairDist {
	shortestPair := pairDist{
		first:    0,
		second:   1,
		distance: LevenshteinDistance(blocks[0], blocks[1]),
		common:   CommonChars(blocks[0], blocks[1]),
	}
	for i := 0; i < len(blocks)-2; i++ {
		for j := i + 1; j < len(blocks)-1; j++ {
			dist := LevenshteinDistance(blocks[i], blocks[j])
			if dist < shortestPair.distance {
				shortestPair.first = i
				shortestPair.second = j
				shortestPair.distance = dist
				shortestPair.common = CommonChars(blocks[i], blocks[j])
			}
		}
	}
	return shortestPair
}

// FindShortestFromFile will find the shortest pair of strings
// using the Levenshtein Distance to calculate them.
func FindShortestFromFile(file string) pairDist {
	scanner := MustScanFile(file)
	blocks := ToArray(scanner)
	return FindShortestPair(blocks)
}
