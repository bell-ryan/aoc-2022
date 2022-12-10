package day10

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	cycle        int
	register     int
	tracker      map[int]int
	total        int
	linePosition int
	crt          string
}

func (c *cpu) incrementCycle() {
	c.cycle++
	c.record()
}

func (c *cpu) updateRegister(a int) {
	c.register += a
}
func (c *cpu) record() {
	if c.cycle%20 == 0 {
		c.tracker[c.cycle] = c.register
	}
}

func (c *cpu) calcTotal() {
	for cycle, val := range c.tracker {
		m := 0
		switch cycle {
		case 20:
			m = cycle * val
		case 60:
			m = cycle * val
		case 100:
			m = cycle * val
		case 140:
			m = cycle * val
		case 180:
			m = cycle * val
		case 220:
			m = cycle * val
		}
		c.total += m
	}
}

func (c *cpu) draw() {

	if inRange(c.linePosition, c.register-1, c.register+1) {
		c.crt += "#"
	} else {
		c.crt += "."
	}

	c.linePosition++
	if c.cycle%40 == 0 {
		c.linePosition = 1
		c.crt += "\n"
	}
}

func inRange(i, min, max int) bool {
	return (i >= min) && (i <= max)
}
func Solution() {
	instructionSet := strings.Split(utils.GetInputData(10), "\n")

	cpu := cpu{}
	cpu.tracker = map[int]int{}
	cpu.register = 1
	cpu.linePosition = 1

	for _, instruction := range instructionSet {

		if instruction == "noop" {
			cpu.incrementCycle()
			cpu.draw()
			continue
		}
		amount := strings.Split(instruction, " ")[1]
		a, _ := strconv.Atoi(amount)
		cpu.incrementCycle()
		cpu.draw()
		cpu.incrementCycle()
		cpu.updateRegister(a)
		cpu.draw()

	}
	cpu.calcTotal()

	fmt.Println(cpu.total)
	fmt.Println(cpu.crt)

}
