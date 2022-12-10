package main

import (
	"fmt"
	"strconv"

	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

type dir struct {
	size    int
	name    string
	subDirs []*dir
}

func (d *dir) printTree() {
	var walk func(*dir, uint)
	walk = func(wd *dir, indent uint) {

		prefix := strings.Repeat(" ", int(indent))
		if indent != 0 {
			prefix += "- "
		}
		fmt.Printf("%s (%d)", prefix+wd.name, wd.size)

		for _, sd := range wd.subDirs {
			fmt.Println()
			walk(sd, indent+2)
		}
	}

	walk(d, 0)
	fmt.Println()
}

func (d *dir) lesserThan(threshhold int) []*dir {
	var walk func(*dir) []*dir
	walk = func(wd *dir) []*dir {
		res := []*dir{}

		if wd.size < threshhold {
			res = append(res, wd)
		}

		for _, sub := range wd.subDirs {
			res = append(res, walk(sub)...)
		}

		return res
	}

	return walk(d)
}

func (d *dir) smallestBiggerThan(threshold int) *dir {
	var walk func(*dir) *dir
	walk = func(wd *dir) *dir {
		var cand *dir
		diff := wd.size - threshold
		if diff >= 0 {
			cand = wd
		}

		for _, sd := range wd.subDirs {
			res := walk(sd)

			if res != nil {
				if cand == nil {
					cand = res
				} else if cand.size > res.size {
					cand = res
				}
			}
		}

		return cand
	}

	return walk(d)
}

func buildTree(commands []string) *dir {
	var walk func(int) (int, *dir)
	walk = func(startIdx int) (lastIdx int, d *dir) {
		resDir := &dir{}
		idx := startIdx
		for {
			switch {
			case idx >= len(commands) || commands[idx] == "$ cd ..":
				return idx, resDir
			case strings.HasPrefix(commands[idx], "$ cd "):
				dname := commands[idx][len("$ cd "):]

				lidx, dir := walk(idx + 1)

				dir.name = dname
				resDir.size += dir.size
				resDir.subDirs = append(resDir.subDirs, dir)
				idx = lidx
			case strings.HasPrefix(commands[idx], "dir"):
				// nop
			case commands[idx] == "$ ls":
				// nop
			case strings.ContainsAny(string(commands[idx][0]), "0123456789"):
				num, _, _ := strings.Cut(commands[idx], " ")
				resDir.size += cmn.Must(strconv.Atoi(num))
			default:
				panic(commands[idx])
			}
			idx += 1
		}
	}

	_, root := walk(0)

	return root.subDirs[0]
}

func main() {
	lns := cmn.ReadLines("input.txt")

	root := buildTree(lns)

	dirs := []int{}
	for _, v := range root.lesserThan(100_000) {
		dirs = append(dirs, v.size)
	}

	fmt.Printf("Part1 -> %d\n", cmn.Sum(dirs...))

	diskSize := 70_000_000
	minFree := 30_000_000
	used := root.size
	toFree := cmn.Abs((diskSize - used) - minFree)

	// root.printTree()

	d := root.smallestBiggerThan(toFree)
	fmt.Printf("Part2 -> %d\n", d.size)
}
