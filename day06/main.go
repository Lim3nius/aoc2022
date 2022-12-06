package main

import (
	"fmt"

	cmn "github.com/Lim3nius/aoc2022/common"
)

func detectUnique(in []byte) bool {
	x := map[byte]uint{}
	for _, v := range in {
		if _, ok := x[v]; ok {
			return false
		}
		x[v] = 1
	}
	return true
}

func detectFirstUniqueSeq(in string, slen int) int {
	for i := 0; i < len(in); i++ {
		if detectUnique([]byte(in[i : i+slen])) {
			return i + slen
		}
	}
	return -42
}

func main() {
	ln := cmn.ReadLines("input.txt")[0]

	fmt.Printf("Part1 -> %d\n", detectFirstUniqueSeq(ln, 4))
	fmt.Printf("Part2 -> %d\n", detectFirstUniqueSeq(ln, 14))
}
