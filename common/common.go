package common

import (
	"os"
	"strings"

	"golang.org/x/exp/constraints"
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

func Max[T constraints.Ordered](args ...T) T {
	switch len(args) {
	case 0:
		panic("no arguments")
	case 1:
		return args[0]
	case 2:
		if args[0] > args[1] {
			return args[0]
		}
		return args[1]
	default:
		m := args[0]
		for _, v := range args[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}
}

func Min[T constraints.Ordered](args ...T) T {
	switch len(args) {
	case 0:
		panic("no arguments")
	case 1:
		return args[0]
	case 2:
		if args[0] < args[1] {
			return args[0]
		}
		return args[1]
	default:
		m := args[0]
		for _, v := range args[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}
}
