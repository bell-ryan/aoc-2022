package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func checkLeft(height int, row []int, position int, views []int) bool {
	if position == 0 {
		return true
	}
	if height <= row[position-1] {
		views[0] += 1
		return false
	}
	views[0] += 1
	return checkLeft(height, row, position-1, views)
}

func checkRight(height int, row []int, position int, views []int) bool {
	if position == len(row)-1 {
		return true
	}
	if height <= row[position+1] {
		views[1] += 1
		return false
	}
	views[1] += 1
	return checkRight(height, row, position+1, views)
}

func checkTop(height int, rows [][]int, position []int, views []int) bool {
	if position[0] == 0 {
		return true
	}
	if height <= rows[position[0]-1][position[1]] {
		views[2] += 1
		return false
	}
	p := make([]int, len(position))
	copy(p, position)
	p[0] -= 1
	views[2] += 1
	return checkTop(height, rows, p, views)
}

func checkBottom(height int, rows [][]int, position []int, views []int) bool {
	if position[0] == len(rows)-1 {
		return true
	}
	if height <= rows[position[0]+1][position[1]] {
		views[3] += 1
		return false
	}
	p := make([]int, len(position))
	copy(p, position)
	p[0] += 1
	views[3] += 1
	return checkBottom(height, rows, p, views)
}

func convertGridToInts() [][]int {
	rows := strings.Split(input, "\n")
	grid := [][]int{}

	for row := 0; row < len(rows); row++ {
		tempRow := []int{}
		for col := 0; col < len(rows[0]); col++ {
			val, _ := strconv.Atoi(string(rows[row][col]))
			tempRow = append(tempRow, val)
		}
		grid = append(grid, tempRow)
	}
	return grid
}
func Solution() {
	rows := convertGridToInts()

	visibleTrees := (len(rows)*2 - 4) + (len(rows[0]) * 2)
	theCrib := 0
	var left, right, top, bottom bool

	for row := 1; row < len(rows)-1; row++ {
		for col := 1; col < len(rows[0])-1; col++ {
			position := []int{row, col}
			val := rows[row][col]
			views := []int{0, 0, 0, 0}
			left = checkLeft(val, rows[row], col, views)
			right = checkRight(val, rows[row], col, views)
			top = checkTop(val, rows, position, views)
			bottom = checkBottom(val, rows, position, views)

			// part2
			viewDistance := views[0] * views[1] * views[2] * views[3]
			if viewDistance > theCrib {
				theCrib = viewDistance
			}
			//

			// part1
			if left {
				visibleTrees += 1
				continue
			}

			if right {
				visibleTrees += 1
				continue
			}

			if top {
				visibleTrees += 1
				continue
			}

			if bottom {
				visibleTrees += 1
				continue
			}

		}

	}

	fmt.Printf("The Crib: %d\n", theCrib)
	fmt.Printf("Visible Trees: %d\n", visibleTrees)
}
