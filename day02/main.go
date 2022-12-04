package main

import (
	"fmt"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

const (
	Rock     = "X"
	Paper    = "Y"
	Scissors = "Z"
)

var (
	otherToMe = map[string]string{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	points = map[string]uint{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	winMap = map[string]string{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}

	meToWantRes = map[string]uint{
		WinC:  Win,
		DrawC: Draw,
		LoseC: Lose,
	}
)

const (
	Win  = 6
	Draw = 3
	Lose = 0

	LoseC = "X"
	DrawC = "Y"
	WinC  = "Z"
)

func eval(me, ot string) uint {
	switch {
	case me == ot:
		return Draw
	case ot == winMap[me]:
		return Win
	default:
		return Lose
	}
}

func ensureResult(op string, wantRes uint) uint {
	for _, me := range []string{Rock, Paper, Scissors} {
		if wantRes == eval(me, op) {
			return wantRes + points[me]
		}
	}
	return 0
}

func main() {
	data := cmn.ReadFile("input.txt")

	score := uint(0)
	score2 := uint(0)
	for _, ln := range strings.Split(data, "\n") {
		pl := strings.Split(ln, " ")
		ot, me := pl[0], pl[1]
		ot = otherToMe[ot]

		score += eval(me, ot) + points[me]
		score2 += ensureResult(ot, meToWantRes[me])
	}

	fmt.Printf("Part1 -> %d\n", score)
	fmt.Printf("Part2 -> %d\n", score2)
}
