package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type location struct {
	vertical   int
	horizontal int
}
type rope struct {
	head      location
	tail      location
	tailStops map[string]struct{}
}

type instruction struct {
	direction string
	steps     int
}

func (r *rope) moveLeft() {
	r.head.horizontal -= 1
	if r.checkSeparation() {
		r.tail.horizontal = r.head.horizontal + 1
		r.tail.vertical = r.head.vertical
	}
}
func (r *rope) moveRight() {
	r.head.horizontal += 1
	if r.checkSeparation() {
		r.tail.horizontal = r.head.horizontal - 1
		r.tail.vertical = r.head.vertical
	}
}
func (r *rope) moveUp() {
	r.head.vertical += 1
	if r.checkSeparation() {
		r.tail.horizontal = r.head.horizontal
		r.tail.vertical = r.head.vertical - 1
	}
}
func (r *rope) moveDown() {
	r.head.vertical -= 1
	if r.checkSeparation() {
		r.tail.horizontal = r.head.horizontal
		r.tail.vertical = r.head.vertical + 1
	}
}

func (r *rope) checkSeparation() bool {
	hSep := r.head.horizontal - r.tail.horizontal
	vSep := r.head.vertical - r.tail.vertical
	if hSep > 1 || hSep < -1 {
		return true
	}
	if vSep > 1 || vSep < -1 {
		return true
	}

	return false
}

func (r *rope) walkItOut(direction string, steps int) {
	for step := 0; step < steps; step++ {
		switch direction {
		case "L":
			r.moveLeft()
		case "R":
			r.moveRight()
		case "U":
			r.moveUp()
		case "D":
			r.moveDown()
		}
		tailLoc := fmt.Sprintf("%d%d", r.tail.vertical, r.tail.horizontal)
		r.tailStops[tailLoc] = struct{}{}
	}

}
func getInput() []instruction {
	steps := strings.Split(input, "\n")
	returnData := []instruction{}
	for _, step := range steps {
		instructions := strings.Split(step, " ")
		instruction := instruction{}
		count, _ := strconv.Atoi(instructions[1])
		instruction.direction = instructions[0]
		instruction.steps = count
		returnData = append(returnData, instruction)
	}
	return returnData
}
func Solution() {
	data := getInput()
	rope := rope{}

	// Cheesing the starting location because im too dumb to handle negative ints :(
	rope.head.horizontal = 1000
	rope.head.vertical = 1000
	rope.tail.horizontal = 1000
	rope.tail.vertical = 1000

	rope.tailStops = map[string]struct{}{}
	for _, move := range data {
		rope.walkItOut(move.direction, move.steps)
	}
	fmt.Println(len(rope.tailStops))
}
