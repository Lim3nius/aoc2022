package main

import (
	"fmt"

	"strconv"

	cmn "github.com/Lim3nius/aoc2022/common"
)

func isVisible(grid [][]int, r, c int) bool {
	treeHeight := grid[r][c]

	walkGrid := func(grid [][]int, rv, cv int, sr, sc int, fn func(int) bool) bool {
		r, c := sr, sc

		for 0 <= r && r < len(grid) &&
			0 <= c && c < len(grid[r]) {
			if !fn(grid[r][c]) {
				return false
			}

			r += rv
			c += cv
		}
		return true
	}

	cond := func(th int) bool { return th < treeHeight }
	leftVis := walkGrid(grid, 0, -1, r, c-1, cond)
	rightVis := walkGrid(grid, 0, 1, r, c+1, cond)
	topVis := walkGrid(grid, -1, 0, r-1, c, cond)
	botVis := walkGrid(grid, 1, 0, r+1, c, cond)

	return leftVis || rightVis || topVis || botVis
}

func edgeTrees(grid [][]int) int {
	return 2*len(grid) + 2*(len(grid[0])-2)
}

func scenicScore(grid [][]int, r, c int) int {
	treeH := grid[r][c]

	walkGrid := func(rv, cv int, sr, sc int, fn func(int) bool) int {
		r, c := sr, sc

		acc := 0
		for 0 <= r && r < len(grid) &&
			0 <= c && c < len(grid[0]) {
			if fn(grid[r][c]) {
				acc += 1
			} else {
				acc += 1
				return acc
			}

			r += rv
			c += cv
		}
		return acc
	}

	cond := func(th int) bool { return th < treeH }

	left := walkGrid(0, -1, r, c-1, cond)
	right := walkGrid(0, 1, r, c+1, cond)
	top := walkGrid(-1, 0, r-1, c, cond)
	bot := walkGrid(1, 0, r+1, c, cond)
	return left * right * top * bot
}

func main() {
	lns := cmn.ReadLines("input.txt")

	grid := make([][]int, len(lns))
	for r := range lns {
		for c := range lns[r] {
			grid[r] = append(grid[r], cmn.Must(strconv.Atoi(string(lns[r][c]))))
		}
	}

	vis := 0
	maxScenic := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if isVisible(grid, r, c) {
				vis += 1
			}
			maxScenic = cmn.Max(scenicScore(grid, r, c), maxScenic)
		}
	}

	vis += edgeTrees(grid)
	fmt.Printf("Part1 -> %d\n", vis)
	fmt.Printf("Part2 -> %d\n", maxScenic)
}
