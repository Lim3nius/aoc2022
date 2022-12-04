package main

import (
	"fmt"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

func value(b byte) int {
	switch {
	case int('a') <= int(b) && int(b) <= int('z'):
		return int(b) - int('a') + 1
	case int('A') <= int(b) && int(b) <= int('Z'):
		return int(b) - int('A') + 27
	}
	panic("wtf " + string(b))
}

func set(s string) map[byte]bool {
	m := map[byte]bool{}
	for _, c := range s {
		m[byte(c)] = true
	}
	return m
}

func findCommon(lns ...string) byte {
	m := map[byte]int{}
	cnt := len(lns)
	for _, ln := range lns {
		ln = strings.Trim(ln, "\n")
		s := set(ln)
		for k := range s {
			m[k] += 1
			if m[k] == cnt {
				return k
			}
		}
	}
	panic("foo")
}

func main() {
	data := cmn.ReadFile("input.txt")
	acc := 0
	for _, ln := range strings.Split(data, "\n") {
		ln = strings.Trim(ln, "\n")
		acc += value(findCommon(ln[:len(ln)/2], ln[len(ln)/2:]))
	}

	sm := 0
	lns := strings.Split(data, "\n")
	for i := 0; i < len(lns); i += 3 {
		sm += value(findCommon(lns[i : i+3]...))
	}

	fmt.Printf("Part1 -> %d\n", acc)
	fmt.Printf("Part2 -> %d\n", sm)
}
