package main

import (
	"github.com/klauern/adventcode-2018/helpers"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type logEntry struct {
	wake    time.Time
	sleep   time.Time
	begin   time.Time
	guardID int
}

type recordType int


const (
	NONE recordType = iota
	BEGIN
	SLEEP
	WAKE
)

var timeLayout = "2006-01-02 15:04"

func main() {}

func parseTimestamp(segments []string) time.Time {
	newLine := strings.Split(line, "] ")
	t := strings.Trim(newLine[0], "[")
	rec, err := time.Parse(timeLayout, t)
	if err != nil {
		panic(errors.Wrap(err, "unable to parse "+t))
	}
	return rec
}

func parseRecordType(line string) recordType {
	l := strings.Split(line, "] ")

	action := strings.Split(l[1], " ")
	switch action[0] {
	case "Guard":
		return BEGIN
	case "falls":
		return SLEEP
	case "wakes":
		return WAKE
	default:
		return NONE
	}
}

func parseLogEntry(line string) logEntry {
	fields := strings.Split(line, " ")

	
}

func main() {
	var guards := make(map[int][]logEntry)
	scanner := helpers.MustScanFile("input.txt")
	for _, v := range helpers.ToArray(scanner) {
		rec := parseRecordType(v)
		guards[rec.guardID] = append(guards[rec.guardID], rec)
	}
}

