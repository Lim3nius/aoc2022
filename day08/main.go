package main

import (
	"fmt"

	"strconv"

	cmn "github.com/Lim3nius/aoc2022/common"
)

func visibleRow(grid [][]int, treeR, treeC int) bool {
	treeHeight := grid[treeR][treeC]
	stop := len(grid[treeR])

	left := func() bool {
		for i := 0; i < treeC; i++ {
			if grid[treeR][i] >= treeHeight {
				return false
			}
		}
		return true
	}

	right := func() bool {
		for i := treeC + 1; i < stop; i++ {
			if grid[treeR][i] >= treeHeight {
				return false
			}
		}
		return true
	}

	return left() || right()
}

func visibleCol(grid [][]int, treeR, treeC int) bool {
	treeHeight := grid[treeR][treeC]
	stop := len(grid)

	top := func() bool {
		for i := 0; i < treeR; i++ {
			if grid[i][treeC] >= treeHeight {
				return false
			}
		}
		return true
	}

	bot := func() bool {
		for i := treeR + 1; i < stop; i++ {
			if grid[i][treeC] >= treeHeight {
				return false
			}
		}
		return true
	}

	return top() || bot()
}

func isVisible(grid [][]int, r, c int) bool {
	return visibleCol(grid, r, c) || visibleRow(grid, r, c)
}

func edgeTrees(grid [][]int) int {
	return 2*len(grid) + 2*(len(grid[0])-2)
}

func scenicScore(grid [][]int, r, c int) int {
	treeH := grid[r][c]

	left := func() int {
		t := 0
		for i := c - 1; i >= 0; i-- {
			if grid[r][i] < treeH {
				t += 1
			} else if grid[r][i] == treeH {
				t += 1
				break
			} else {
				t += 1
				break
			}
		}
		return t
	}

	right := func() int {
		// fmt.Printf("%d %d\n", r, c)
		t := 0
		for i := c + 1; i < len(grid[r]); i++ {
			if grid[r][i] < treeH {
				t += 1
			} else if grid[r][i] == treeH {
				t += 1
				break
			} else {
				t += 1
				break
			}
		}
		return t
	}

	top := func() int {
		t := 0
		for i := r - 1; i >= 0; i-- {
			if grid[i][c] < treeH {
				t += 1
			} else if grid[i][c] == treeH {
				t += 1
				break
			} else {
				t += 1
				break
			}
		}
		return t
	}

	bot := func() int {
		t := 0
		for i := r + 1; i < len(grid); i++ {
			if grid[i][c] < treeH {
				t += 1
			} else if grid[i][c] == treeH {
				t += 1
				break
			} else {
				t += 1
				break
			}
		}
		return t
	}

	l, rr, t, b := left(), right(), top(), bot()
	return l * rr * t * b
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

			if m := scenicScore(grid, r, c); m > maxScenic {
				maxScenic = m
			}
		}
	}

	vis += edgeTrees(grid)
	fmt.Printf("Part1 -> %d\n", vis)
	fmt.Printf("Part2 -> %d\n", maxScenic)
}
