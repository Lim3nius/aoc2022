package main

import (
	"fmt"

	cmn "github.com/Lim3nius/aoc2022/common"
)

type pos struct {
	row, col int
}

func (p *pos) String() string {
	return fmt.Sprintf("%d,%d", p.row, p.col)
}

func parseGrid(lns []string) ([][]int, *pos, *pos) {
	grid := make([][]int, len(lns))
	start, target := &pos{}, &pos{}
	for r := range lns {
		grid[r] = make([]int, len(lns[r]))

		for c, v := range lns[r] {
			v := v
			if v == rune('S') {
				v = rune('a')
				start.row = r
				start.col = c
			} else if v == rune('E') {
				v = rune('z')
				target.row = r
				target.col = c
			}

			grid[r][c] = int(v)
		}
	}
	return grid, start, target
}

func createPathGrid(grid [][]int) [][]int {
	ng := make([][]int, len(grid))
	for r := range grid {
		ng[r] = make([]int, len(grid[r]))
	}
	return ng
}

func getPossiblePaths(zoneMap [][]int, from *pos) []*pos {
	check := func(rd, cd int) *pos {
		r, c := from.row+rd, from.col+cd
		if 0 <= r && r < len(zoneMap) &&
			0 <= c && c < len(zoneMap[r]) {
			if zoneMap[r][c]-zoneMap[from.row][from.col] <= 1 {
				return &pos{row: r, col: c}
			}
		}
		return nil
	}

	eval := func(in ...*pos) []*pos {
		res := []*pos{}
		for _, v := range in {
			v := v
			if v != nil {
				res = append(res, v)
			}
		}
		if len(res) > 0 {
			return res
		}
		return nil
	}
	return eval(
		check(1, 0),
		check(-1, 0),
		check(0, 1),
		check(0, -1),
	)
}

func printGrid(grid [][]int) {
	for _, r := range grid {
		for _, v := range r {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
		fmt.Println()
	}
}

func printVisited(grid [][]int) {
	for _, r := range grid {
		for _, v := range r {
			if v != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func travelGrid(zoneMap [][]int, start, target *pos) [][]int {
	pathMap := createPathGrid(zoneMap)
	visitedPositions := []*pos{start}

	for {
		if len(visitedPositions) == 0 {
			break
		}

		// printVisited(pathMap)
		// fmt.Println()

		updatePositions := func(positions []*pos) ([]*pos, bool) {
			lst := []*pos{}

			for _, p := range positions {
				// get possible positions grid
				paths := getPossiblePaths(zoneMap, p)
				// fmt.Printf("Found positions: %v\n", paths)

				// update zone map
				for _, path := range paths {
					path := path
					if pathMap[path.row][path.col] == 0 && (path.row != start.row || path.col != start.col) {
						pathMap[path.row][path.col] = pathMap[p.row][p.col] + 1
						lst = append(lst, path)
					}

					// reached end
					if path.row == target.row && path.col == target.col {
						return nil, true
					}
				}

				// append new positions
			}

			return lst, false
		}

		finished := false
		visitedPositions, finished = updatePositions(visitedPositions)
		if finished {
			break
		}
	}

	return pathMap
}

func main() {
	lns := cmn.ReadLines("input.txt")
	grid, start, target := parseGrid(lns)

	pathMap := travelGrid(grid, start, target)

	fmt.Printf("Part1 -> %d\n", pathMap[target.row][target.col])

	possibleStarts := []*pos{}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == int(rune('a')) {
				possibleStarts = append(possibleStarts, &pos{row: r, col: c})
			}
		}
	}

	fewest := 0
	for _, ps := range possibleStarts {
		res := travelGrid(grid, ps, target)
		if fewest == 0 {
			fewest = res[target.row][target.col]
		}
		if res[target.row][target.col] == 0 {
			continue
		}

		fewest = cmn.Min(fewest, res[target.row][target.col])
	}

	fmt.Printf("Part2 -> %d\n", fewest)
}
