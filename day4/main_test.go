package main

import (
	"reflect"
	"testing"
	"time"
)

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

func Test_parseRecord(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"1518-11-01 00:00",
			args{"1518-11-01 00:00"},
			time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"1518-11-03 00:24",
			args{"1518-11-03 00:24"},
			time.Date(1518, 11, 3, 0, 24, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRecord(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRecordType(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want recordType
	}{
		{
			"Guard Begins",
			args{"[1518-11-01 00:00] Guard #10 begins shift"},
			BEGIN,
		},
		{
			"falls asleep",
			args{"[1518-11-01 00:30] falls asleep"},
			SLEEP,
		},
		{
			"wakes up",
			args{"[1518-11-01 00:55] wakes up"},
			WAKE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRecordType(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRecordType() = %v, want %v", got, tt.want)
			}
		})
	}
}
