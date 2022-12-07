package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
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
	fmt.Println("-----------------")

	// Day 3
	p1, p2 = day3.Solution()
	fmt.Printf("Day 3 - part 1 answer %d\n", p1)
	fmt.Printf("Day 3 - part 2 answer %d\n", p2)
	fmt.Println("-----------------")

	// Day 4
	p1, p2 = day4.Solution()
	fmt.Printf("Day 4 - part 1 answer %d\n", p1)
	fmt.Printf("Day 4 - part 2 answer %d\n", p2)
	fmt.Println("-----------------")

	// Day 5
	// p1, p2 = day5.Solution()
	// fmt.Printf("Day 5 - part 1 answer %d\n", p1)
	// fmt.Printf("Day 5 - part 2 answer %d\n", p2)
	day5.Solution()
	fmt.Println("-----------------")

	// Day 6
	p1, p2 = day6.Start()
	fmt.Printf("Day 6 - part 1 answer %d\n", p1)
	fmt.Printf("Day 6 - part 2 answer %d\n", p2)
	fmt.Println("-----------------")

	// Day 7
	// p1, p2 = day6.Start()
	// fmt.Printf("Day 6 - part 1 answer %d\n", p1)
	// fmt.Printf("Day 6 - part 2 answer %d\n", p2)
	day7.Solution()
	fmt.Println("-----------------")

}
