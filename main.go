package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"fmt"
)

func main() {
	// Day 1
	p1, p2 := day1.Solution()
	fmt.Printf("Day 1 - part 1 answer %d\n", p1)
	fmt.Printf("Day 1 - part 2 answer %d\n", p2)
	fmt.Println("-----------------")

	// Day 2
	p1, p2 = day2.Solution()
	fmt.Printf("Day 2 - part 1 answer %d\n", p1)
	fmt.Printf("Day 2 - part 2 answer %d\n", p2)

	// Day 3
	day3.Solution()
}
