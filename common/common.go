package common

import (
	"os"
	"strings"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func ReadFile(fpath string) string {
	data := string(Must(os.ReadFile(fpath)))
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return data
}

func ReadLines(fpath string) []string {
	lns := []string{}
	data := ReadFile(fpath)
	for _, ln := range strings.Split(data, "\n") {
		ln = strings.TrimRight(ln, "\n")
		lns = append(lns, ln)
	}
	return lns
}
