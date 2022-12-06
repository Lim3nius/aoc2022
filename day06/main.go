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

func main() {
	ln := cmn.ReadLines("input.txt")[0]

	for i := 0; i < len(ln); i++ {
		if detectUnique([]byte(ln[i : i+4])) {
			fmt.Printf("Part1 -> %d\n", i+4)
			break
		}
	}

	for i := 0; i < len(ln); i++ {
		if detectUnique([]byte(ln[i : i+14])) {
			fmt.Printf("Part2 -> %d\n", i+14)
			break
		}
	}

}
