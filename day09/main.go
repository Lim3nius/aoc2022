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

func main() {
	lns := cmn.ReadLines("input.txt")
	headPos := &point{}
	tailPos := &point{}

	tailTrail := map[int]map[int]int{}

	tailTrails := map[int]map[int]int{}

	ropes := []*point{{
		x: 0,
		y: 0,
	}}
	endsC := 10
	for _, cmd := range lns {
		v, l := parseCmd(cmd)
		for x := 0; x < int(l); x++ {
			headPos.Move(v)
			tailPos.Follow(headPos)

			if _, ok := tailTrail[tailPos.x]; !ok {
				tailTrail[tailPos.x] = map[int]int{}
			}
			tailTrail[tailPos.x][tailPos.y] += 1

			ropes[0].Move(v)
			for i := 1; i < len(ropes); i++ {
				ropes[i].Follow(ropes[i-1])

			}

			if len(ropes) < endsC && (ropes[len(ropes)-1].x != 0 || ropes[len(ropes)-1].y != 0) {
				ropes = append(ropes, &point{x: 0, y: 0})
			}

			if len(ropes) == 10 {
				if _, ok := tailTrails[ropes[9].x]; !ok {
					tailTrails[ropes[9].x] = map[int]int{}
				}
				tailTrails[ropes[9].x][ropes[9].y] += 1
			}

		}
	}

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

	fmt.Printf("Part1 -> %d\n", visits(tailTrail))
	fmt.Printf("Part2 -> %d\n", visits(tailTrails))
}
