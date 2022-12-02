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

		switch mine {
		case "X":
			switch op {
			case "A":
				total += 4
				total2 += 3
			case "B":
				total += 1
				total2 += 1
			case "C":
				total += 7
				total2 += 2
			}
		case "Y":
			switch op {
			case "A":
				total += 8
				total2 += 4
			case "B":
				total += 5
				total2 += 5
			case "C":
				total += 2
				total2 += 6
			}
		case "Z":
			switch op {
			case "A":
				total += 3
				total2 += 8
			case "B":
				total += 9
				total2 += 9
			case "C":
				total += 6
				total2 += 7
			}
		}

	}
	return total, total2
}
