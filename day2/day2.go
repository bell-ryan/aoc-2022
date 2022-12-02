package day2

import (
	"aoc/utils"
	"strings"
)

func Solution() (int, int) {
	data := utils.GetInputData(2)
	total := 0
	total2 := 0

	for _, match := range strings.Split(data, "\n") {
		selections := strings.Split(match, " ")
		op, mine := selections[0], selections[1]

		if mine == "X" {
			if op == "A" {
				total += 4
				total2 += 3
			}
			if op == "B" {
				total += 1
				total2 += 1
			}
			if op == "C" {
				total += 7
				total2 += 2
			}
		}
		if mine == "Y" {
			if op == "A" {
				total += 8
				total2 += 4
			}
			if op == "B" {
				total += 5
				total2 += 5
			}
			if op == "C" {
				total += 2
				total2 += 6
			}
		}
		if mine == "Z" {
			if op == "A" {
				total += 3
				total2 += 8
			}
			if op == "B" {
				total += 9
				total2 += 9
			}
			if op == "C" {
				total += 6
				total2 += 7
			}
		}

	}
	return total, total2
}
