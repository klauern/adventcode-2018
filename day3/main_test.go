package main

import (
	"reflect"
	"testing"
)

func Test_parseClaim(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *claim
	}{
		{
			"example claim",
			args{"#123 @ 3,2: 5x4"},
			&claim{123, &area{3, 2, 5, 4}},
		},
		{
			"first overlap",
			args{"#1 @ 1,3: 4x4"},
			&claim{1, &area{1, 3, 4, 4}},
		},
		{
			"second overlap",
			args{"#2 @ 3,1: 4x4"},
			&claim{2, &area{3, 1, 4, 4}},
		},
		{
			"third overlap",
			args{"#3 @ 5,5: 2x2"},
			&claim{3, &area{5, 5, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseClaim(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mustConv(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mustConv(tt.args.input); got != tt.want {
				t.Errorf("mustConv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_claim_overlaps(t *testing.T) {
	type args struct {
		d *claim
	}
	tests := []struct {
		name string
		c    *claim
		args args
		want bool
	}{
		{
			"example",
			parseClaim("#1 @ 1,3: 4x4"),
			args{
				parseClaim("#2 @ 3,1: 4x4"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.overlaps(tt.args.d); got != tt.want {
				t.Errorf("claim.overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_claim_findOverlap(t *testing.T) {
	type args struct {
		d *claim
	}
	tests := []struct {
		name string
		c    *claim
		args args
		want *area
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.findOverlap(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("claim.findOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyClaimToFabric(t *testing.T) {
	type args struct {
		c *claim
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			applyClaimToFabric(tt.args.c)
		})
	}
}

func Test_countFabricOverlap(t *testing.T) {
	type args struct {
		fabric [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example scenario",
			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFabricOverlap(tt.args.fabric); got != tt.want {
				t.Errorf("countFabricOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
