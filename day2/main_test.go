package main

import (
	"bufio"
	"reflect"
	"testing"
)

func TestCalcDistance(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			"abcdef",
			args{"abcdef"},
			map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1},
		},
		{
			"babac",
			args{"babac"},
			map[rune]int{'b': 2, 'a': 2, 'c': 1},
		},
		{
			"abbcde",
			args{"abbcde"},
			map[rune]int{'a': 1, 'b': 2, 'c': 1, 'd': 1, 'e': 1},
		},
		{
			"abcccd",
			args{"abcccd"},
			map[rune]int{'a': 1, 'b': 1, 'c': 3, 'd': 1},
		},
		{
			"aabcdd",
			args{"aabcdd"},
			map[rune]int{'a': 2, 'b': 1, 'c': 1, 'd': 2},
		},
		{
			"abcdee",
			args{"abcdee"},
			map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 2},
		},
		{
			"ababab",
			args{"ababab"},
			map[rune]int{'a': 3, 'b': 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcDistance(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustScanFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *bufio.Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustScanFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustScanFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArray(t *testing.T) {
	type args struct {
		input *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArray(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasTwoLetter(t *testing.T) {
	type args struct {
		id map[rune]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcdef",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}},
			false,
		},
		{
			"babac",
			args{map[rune]int{'b': 2, 'a': 2, 'c': 1}},
			true,
		},
		{
			"abbcde",
			args{map[rune]int{'a': 1, 'b': 2, 'c': 1, 'd': 1, 'e': 1}},
			true,
		},
		{
			"abcccd",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 3, 'd': 1}},
			false,
		},
		{
			"aabcdd",
			args{map[rune]int{'a': 2, 'b': 1, 'c': 1, 'd': 2}},
			true,
		},
		{
			"abcdee",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 2}},
			true,
		},
		{
			"ababab",
			args{map[rune]int{'a': 3, 'b': 3}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasTwoLetter(tt.args.id); got != tt.want {
				t.Errorf("HasTwoLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasThreeLetter(t *testing.T) {
	type args struct {
		id map[rune]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcdef",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}},
			false,
		},
		{
			"babac",
			args{map[rune]int{'b': 2, 'a': 2, 'c': 1}},
			false,
		},
		{
			"abbcde",
			args{map[rune]int{'a': 1, 'b': 2, 'c': 1, 'd': 1, 'e': 1}},
			false,
		},
		{
			"abcccd",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 3, 'd': 1}},
			true,
		},
		{
			"aabcdd",
			args{map[rune]int{'a': 2, 'b': 1, 'c': 1, 'd': 2}},
			false,
		},
		{
			"abcdee",
			args{map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 2}},
			false,
		},
		{
			"ababab",
			args{map[rune]int{'a': 3, 'b': 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasThreeLetter(tt.args.id); got != tt.want {
				t.Errorf("HasThreeLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	type args struct {
		boxes []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_input.txt",
			args{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.args.boxes); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
