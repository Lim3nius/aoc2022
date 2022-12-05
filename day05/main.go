package main

import (
	"fmt"
	"strconv"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

func readCol(col int, ln string) string {
	letPos := col*4 + 1
	if len(ln) < letPos {
		return " "
	}
	return string(ln[letPos])
}

func printStacks[T any](sts []Stack[T]) {
	f := func(arr []T) []string {
		res := []string{}
		for _, e := range arr {
			res = append(res, fmt.Sprintf("%+v", e))
		}
		return res
	}
	for i, s := range sts {
		fmt.Printf("%d -> %s\n", i, strings.Join(f(s), " "))
	}
	fmt.Println()
}

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) PopN(n int) []T {
	v := (*s)[len(*s)-n:]
	(*s) = (*s)[:len(*s)-n]
	return v
}

func (s *Stack[T]) Push(t ...T) {
	*s = append(*s, t...)
}

func main() {
	data := cmn.ReadFile("input.txt")
	boxes, instructions, _ := strings.Cut(data, "\n\n")

	blns := strings.Split(boxes, "\n")
	colCount := len(strings.Fields(blns[len(blns)-1]))
	stacks := make([]Stack[string], colCount)
	stacks2 := make([]Stack[string], colCount)
	for i := len(blns) - 2; i >= 0; i-- {
		for c := 0; c < colCount; c++ {
			if v := readCol(c, blns[i]); v != " " {
				stacks[c].Push(v)
				stacks2[c].Push(v)
			}
		}
	}

	for _, ln := range strings.Split(instructions, "\n") {
		ln = strings.TrimRight(ln, "\n")
		parts := strings.Split(ln, " ")
		f := func(x string) int { return cmn.Must(strconv.Atoi(x)) }
		cnt, from, to := f(parts[1]), f(parts[3])-1, f(parts[5])-1

		for i := 0; i < cnt; i++ {
			stacks[to].Push(stacks[from].Pop())
		}
		stacks2[to].Push(stacks2[from].PopN(cnt)...)

	}

	res := ""
	res2 := ""
	for i := range stacks {
		res += stacks[i][len(stacks[i])-1]
		res2 += stacks2[i][len(stacks2[i])-1]
	}

	fmt.Printf("Part1 -> %s\n", res)
	fmt.Printf("Part2 -> %s\n", res2)
}
