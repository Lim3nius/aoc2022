package main

import (
	"fmt"
	"strconv"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

type pair struct {
	left, right uint
}

func (p *pair) String() string {
	return fmt.Sprintf("%d-%d", p.left, p.right)
}

func parsePair(s string) *pair {
	l, r, _ := strings.Cut(s, "-")
	ll := cmn.Must(strconv.Atoi(l))
	rr := cmn.Must(strconv.Atoi(r))
	if ll > rr {
		ll, rr = rr, ll
	}
	return &pair{
		left:  uint(ll),
		right: uint(rr),
	}
}

func fullyContained(f, s *pair) bool {
	fn := func(f, s *pair) bool { return f.left <= s.left && s.right <= f.right }
	return fn(f, s) || fn(s, f)
}

func partialOverlap(f, s *pair) bool {
	fn := func(f, s *pair) bool {
		return f.left <= s.left && s.left <= f.right ||
			f.left <= s.right && s.right <= f.right
	}
	return fn(f, s) || fn(s, f)
}

func main() {
	data := cmn.ReadLines("input.txt")
	contained := 0
	partial := 0
	for _, ln := range data {
		l, r, _ := strings.Cut(ln, ",")
		fst := parsePair(l)
		snd := parsePair(r)

		if fullyContained(fst, snd) {
			contained += 1
		} else if partialOverlap(fst, snd) {
			partial += 1
		}

	}
	fmt.Printf("Part1 -> %d\n", contained)
	fmt.Printf("Part2 -> %d\n", contained+partial)
}
