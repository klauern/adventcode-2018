package main

import (
	"bufio"
	"reflect"
	"testing"
)

func TestFrequencyCheck(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first test file",
			args{"test_input1.txt"},
			4,
		},
		{
			"second file",
			args{"test_input2.txt"},
			1,
		},
		{
			"example input from description",
			args{"example_input.txt"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrequencyCheck(tt.args.file); got != tt.want {
				t.Errorf("FrequencyCheck() = %v, want %v", got, tt.want)
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

func TestFindDuplicateFreq(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first test",
			args{"test_input1.txt"},
			10,
		},
		{
			"second of the tests",
			args{"test_input2.txt"},
			14,
		},
		{
			"example input",
			args{"example_input.txt"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateFreq(tt.args.file); got != tt.want {
				t.Errorf("FindDuplicateFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
