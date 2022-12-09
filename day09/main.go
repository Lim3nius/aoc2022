package main

import (
	"fmt"
	"strconv"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

type point struct {
	x, y int
}

func (p *point) L1Dist(o *point) int {
	return cmn.Abs(p.x-o.x) + cmn.Abs(p.y-o.y)
}

func (p *point) Vec(o *point) *vec {
	return &vec{
		x: o.x - p.x,
		y: o.y - p.y,
	}
}

func (p *point) Follow(o *point) {
	switch p.L1Dist(o) {
	case 0, 1:
	case 2:
		if cmn.Abs(p.x-o.x) == cmn.Abs(p.y-o.y) {
			return
		}
		p.Move(*p.Vec(o).L1NormVec())
	case 3, 4:
		p.Move(*p.Vec(o).L1NormVec())
	}
}

func (p *point) Move(v vec) {
	p.x += v.x
	p.y += v.y
}

type vec point

func (p *vec) L1NormVec() *vec {
	v := &vec{}

	if p.x < 0 {
		v.x = -1
	} else if p.x > 0 {
		v.x = 1
	}

	if p.y < 0 {
		v.y = -1
	} else if p.y > 0 {
		v.y = 1
	}

	return v
}

func parseCmd(ln string) (vec, uint) {
	d, l, _ := strings.Cut(ln, " ")

	var v vec
	switch d {
	case "D":
		v = vec{x: 0, y: -1}
	case "U":
		v = vec{x: 0, y: +1}
	case "R":
		v = vec{x: +1, y: 0}
	case "L":
		v = vec{x: -1, y: 0}
	}

	return v, uint(cmn.Must(strconv.Atoi(l)))
}

func simulate(cmds []string, nodes int) map[int]map[int]int {
	visited := map[int]map[int]int{}
	end := nodes - 1
	ropes := make([]*point, nodes)
	for i := 0; i < nodes; i++ {
		ropes[i] = &point{}
	}

	for _, cmd := range cmds {
		v, l := parseCmd(cmd)
		for x := 0; x < int(l); x++ {
			ropes[0].Move(v)

			for i := 1; i < len(ropes); i++ {
				ropes[i].Follow(ropes[i-1])
			}

			if _, ok := visited[ropes[end].x]; !ok {
				visited[ropes[end].x] = map[int]int{}
			}
			visited[ropes[end].x][ropes[end].y] += 1
		}
	}
	return visited
}

func main() {
	lns := cmn.ReadLines("input.txt")

	p1Visited := simulate(lns, 2)
	p2Visited := simulate(lns, 10)

	visits := func(m map[int]map[int]int) int {
		acc := 0
		for x := range m {
			for y := range m[x] {
				if m[x][y] >= 1 {
					acc += 1
				}
			}
		}
		return acc
	}

	fmt.Printf("Part1 -> %d\n", visits(p1Visited))
	fmt.Printf("Part2 -> %d\n", visits(p2Visited))
}
