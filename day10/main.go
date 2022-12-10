package main

import (
	"fmt"

	"errors"
	cmn "github.com/Lim3nius/aoc2022/common"
	"strconv"
	"strings"
)

type cpu struct {
	instructions       []string
	pc                 uint
	currentInstruction *instruction
	lastInstruction    *instruction

	middleTickCallback func(*cpu)

	register int
}

func newCpu(instructions []string) *cpu {
	return &cpu{
		instructions: instructions,
		register:     1,
		pc:           0,
	}
}

var errFin = errors.New("finish")

func (c *cpu) tick() error {
	if c.currentInstruction == nil {
		if !c.fetchInstruction() { // failed to fetch, instructions are waisted
			return errFin
		}
	}

	if c.middleTickCallback != nil {
		c.middleTickCallback(c)
	}

	if fin := c.currentInstruction.tick(c); fin {
		c.finishInstruction()
	}
	return nil
}

func (c *cpu) fetchInstruction() bool {
	if int(c.pc) >= len(c.instructions) {
		return false
	}
	ins := parseInstruction(c.instructions[c.pc])
	c.currentInstruction = ins
	return true
}

func (c *cpu) finishInstruction() {
	c.lastInstruction = c.currentInstruction
	c.currentInstruction = nil
	c.pc += 1
}

func (c *cpu) currentIns() *instruction {
	if c.currentInstruction != nil {
		return c.currentInstruction
	}
	return c.lastInstruction
}

type instruction struct {
	ticks int
	cnt   int
	fn    func(*cpu)
	name  string
}

// return true if finished
func (i *instruction) tick(c *cpu) bool {
	i.cnt += 1

	if i.cnt == i.ticks {
		i.fn(c)
		return true
	} else if i.cnt > i.ticks {
		panic("wtf state")
	}

	return false
}

func noopInstruction() *instruction {
	return &instruction{
		ticks: 1,
		fn:    func(*cpu) {},
		name:  "noop",
	}
}

func addxInstruction(v int) *instruction {
	return &instruction{
		ticks: 2,
		fn: func(c *cpu) {
			c.register += v
		},
		name: fmt.Sprintf("addx %d", v),
	}
}

func parseInstruction(ln string) *instruction {
	cmd, arg, _ := strings.Cut(ln, " ")
	switch cmd {
	case "noop":
		return noopInstruction()
	case "addx":
		val := cmn.Must(strconv.Atoi(arg))
		return addxInstruction(val)
	default:
		panic(fmt.Sprintf("unexpected instruction: %q", cmd))
	}
}

func main() {
	lns := cmn.ReadLines("input.txt")

	i := 1
	vals := []int{}
	cpu0 := newCpu(lns)
	cpu0.middleTickCallback = func(c *cpu) {
		if i%40 == 20 {
			vals = append(vals, c.register*i)
		}
	}

	for ; i <= 220; i++ {
		cpu0.tick()
	}

	fmt.Printf("Part1 -> %d\n", cmn.Sum(vals...))

	i = 0
	cpu1 := newCpu(lns)
	cpu1.middleTickCallback = func(c *cpu) {
		if (i%40)-1 <= c.register && c.register <= (i%40)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	for ; cpu1.tick() != errFin; i++ {
	}
}
