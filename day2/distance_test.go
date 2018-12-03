package main

import (
	"reflect"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Asheville/Arizona",
			args{"Asheville", "Arizona"},
			8,
		},
		{
			"fghij/fguij",
			args{"fghij", "fguij"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDistance(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("LevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonChars(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"fghij/fguij",
			args{"fghij", "fguij"},
			"fgij",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonChars(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CommonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindShortestPair(t *testing.T) {
	type args struct {
		blocks []string
	}
	tests := []struct {
		name string
		args args
		want pairDist
	}{
		{
			"example input",
			args{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}},
			pairDist{
				first:    1,
				second:   4,
				distance: 1,
				common:   "fgij"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPair(tt.args.blocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindShortestPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
