package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
	"strings"

	cmn "github.com/Lim3nius/aoc2022/common"
)

type monkey struct {
	id        int64
	items     []int64
	operation *operation
	throwFn   func(int64) int64

	inspectedItems int64
}

func (m *monkey) acceptItem(i int64) {
	m.items = append(m.items, i)
}

func (m *monkey) inspectItem(decrease func(int64) int64) (int64, int64, bool) {
	if len(m.items) == 0 {
		return 0, 0, false
	}

	item := m.items[0]
	m.items = m.items[1:]

	item = m.operation.Do(item)

	// item = item / 3
	item = decrease(item)

	m.inspectedItems++

	return m.throwFn(item), item, true
}

var (
	add = func(x, y int64) int64 { return x + y }
	mul = func(x, y int64) int64 { return x * y }

	opMap = map[string]func(int64, int64) int64{
		"*": mul,
		"+": add,
	}
)

var monkeyRe = regexp.MustCompile(
	`Monkey (?P<id>\d+):
\s*Starting items: (?P<list>(\d+, )+)*(?P<last>\d+)
\s*Operation: new = old (?P<op>\*|\+) (?P<arg>\d+|old)
\s*Test: divisible by (?P<div>\d+)
\s*If true: throw to monkey (?P<true>\d+)
\s*If false: throw to monkey (?P<false>\d+)`)

type operation struct {
	val     int64
	squared bool
	fn      func(int64, int64) int64
}

func (o *operation) Do(x int64) int64 {
	if o.squared {
		return o.fn(x, x)
	}
	return o.fn(x, o.val)
}

func parseMonkeys(lns string) []*monkey {
	mons := []*monkey{}
	for _, m := range monkeyRe.FindAllStringSubmatch(lns, -1) {
		mons = append(mons, parseMonkey(m))
	}
	return mons
}

func parseMonkey(matches []string) *monkey {
	// fmt.Printf("%+v\n", matches[1:])
	mon := &monkey{}
	atoi := func(s string) int64 { return cmn.Must(strconv.ParseInt(s, 10, 64)) }
	getMatch := func(key string) string { return matches[monkeyRe.SubexpIndex(key)] }

	mon.id = atoi(getMatch("id"))

	for _, v := range strings.Split(getMatch("list"), ", ") {
		if v != "" {
			mon.items = append(mon.items, atoi(v))
		}
	}
	mon.items = append(mon.items, atoi(getMatch("last")))

	mon.operation = &operation{
		fn: opMap[getMatch("op")],
	}
	if getMatch("arg") == "old" {
		mon.operation.squared = true
	} else {
		mon.operation.val = atoi(getMatch("arg"))
	}

	div := atoi(getMatch("div"))
	ok := atoi(getMatch("true"))
	nok := atoi(getMatch("false"))
	mon.throwFn = func(in int64) int64 {
		if in%div == 0 {
			return ok
		}
		return nok
	}

	return mon
}

func main() {
	lns := cmn.ReadFile("input.txt")
	monkeys := parseMonkeys(lns)

	// for i, m := range monkeys {
	// 	fmt.Printf("M %d: %v\n", i, m.items)
	// }

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for {
				dest, item, ok := m.inspectItem(func(x int64) int64 { return x / 3 })
				if !ok {
					break
				}

				monkeys[dest].acceptItem(item)
			}
		}
		// for i, m := range monkeys {
		// 	fmt.Printf("M %d: %v\n", i, m.items)
		// }
	}

	slices.SortFunc(monkeys, func(a, b *monkey) bool { return a.inspectedItems > b.inspectedItems })

	fmt.Printf("Part1 -> %d\n", monkeys[0].inspectedItems*monkeys[1].inspectedItems)

	// monkeys2 := parseMonkeys(lns)
	// monkeys2[0].items = monkeys2[0].items[1:2]
	// for _, m := range monkeys2[1:] {
	// 	m.items = nil
	// }

	// cycleDetect := []int64{}
	// for i := 0; i < 20; i++ {
	// 	for _, m := range monkeys2 {
	// 		for {
	// 			d, i, ok := m.inspectItem(func(x int64) int64 { return x })
	// 			if !ok {
	// 				break
	// 			}

	// 			fmt.Printf("%d\n", i)
	// 			cycleDetect = append(cycleDetect, d)
	// 			monkeys2[d].acceptItem(i)
	// 		}
	// 	}

	// 	fmt.Printf("%d: %v\n", i, cycleDetect)
	// }
}
