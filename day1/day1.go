package day1

import (
	"aoc/utils"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	data := utils.GetInputData(1)
	maxCal := 0
	elves := strings.Split(data, "\n\n")

	for _, elf := range elves {
		caloryLoad := 0
		cals := strings.Split(elf, "\n")
		for _, cal := range cals {
			n, err := strconv.Atoi(cal)
			utils.ErrorCheck(err)
			caloryLoad += n
		}
		if caloryLoad > maxCal {
			maxCal = caloryLoad
		}
	}

	return (maxCal)
}

func Part2() int {
	data := utils.GetInputData(1)
	top3 := []int{}
	total := 0
	elves := strings.Split(data, "\n\n")

	for _, elf := range elves {
		caloryLoad := 0
		cals := strings.Split(elf, "\n")
		for _, cal := range cals {
			n, err := strconv.Atoi(cal)
			utils.ErrorCheck(err)
			caloryLoad += n
		}
		if len(top3) < 3 {
			top3 = append(top3, caloryLoad)
			sort.Slice(top3, func(i, j int) bool {
				return top3[i] > top3[j]
			})
			continue
		}

		if top3[2] < caloryLoad {
			top3[2] = caloryLoad
		}
		sort.Slice(top3, func(i, j int) bool {
			return top3[i] > top3[j]
		})

	}
	for _, elf := range top3 {
		total += elf
	}
	return (total)
}
