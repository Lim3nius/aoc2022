package common

import (
	"os"
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
