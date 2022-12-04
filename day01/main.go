package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

type elf struct {
	calories []uint
	calSum   uint
}

func (e *elf) String() string {
	return strconv.FormatUint(uint64(e.calSum), 10)
}

func main() {
	b := must(os.ReadFile("input.txt"))
	data := string(b)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	elv := []*elf{}
	maxElvIdx := 0

	for _, elves := range strings.Split(data, "\n\n") {
		e := &elf{}
		for _, cals := range strings.Split(elves, "\n") {
			c := must(strconv.ParseUint(cals, 10, 32))
			e.calories = append(e.calories, uint(c))
			e.calSum += uint(c)
		}
		elv = append(elv, e)
		if e.calSum > elv[maxElvIdx].calSum {
			maxElvIdx = len(elv) - 1
		}
	}

	fmt.Printf("Part1 -> %d\n", elv[maxElvIdx].calSum)

	slices.SortFunc(elv, func(l, r *elf) bool {
		return l.calSum < r.calSum
	})

	s := uint(0)
	for _, e := range elv[len(elv)-3:] {
		s += e.calSum
	}
	fmt.Printf("Part2 -> %d\n", s)
}
