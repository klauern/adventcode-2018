package helpers

import (
	"bufio"
	"os"

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

// ToArray will convert a *bufio.Scanner into a fully-evaluated array of strings.
func ToArray(input *bufio.Scanner) []string {
	ary := []string{}
	for input.Scan() {
		ary = append(ary, input.Text())
	}

	return ary
}
