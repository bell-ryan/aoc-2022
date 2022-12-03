package day3

import (
	"aoc/utils"
	"strings"
)

func Solution() (int, int) {
	data := utils.GetInputData(3)
	rucks := strings.Split(data, "\n")
	total := 0
	for i := 0; i < len(rucks); i++ {
		ruckBytes := []byte(rucks[i])
		c1, c2 := ruckBytes[:len(ruckBytes)/2], ruckBytes[len(ruckBytes)/2:]
		m := map[byte]int{}
		for _, bit := range c1 {
			m[bit] = 1
		}
		for _, bit := range c2 {
			if _, ok := m[bit]; ok {
				total += bitToInt(bit)
				break
			}
		}
	}
	// part 2
	total2 := 0
	for i := 0; i < len(rucks); i += 3 {

		group := rucks[i : i+3]
		m := make(map[byte][3]int)

		for i := range group {
			ruckBytes := []byte(group[i])
			for _, bit := range ruckBytes {
				tempSlice := m[bit]
				tempSlice[i] = 1
				m[bit] = tempSlice
			}

		}

		for k, v := range m {
			c := 0
			for elfI := range v {
				c += v[elfI]
			}

			if c == 3 {
				total2 += bitToInt(k)
			}

		}
	}
	return total, total2
}

func bitToInt(b byte) int {
	var num int
	if b >= 97 {
		num += int(b - 'a' + 1)

	} else {
		num += int(b - 'A' + 27)
	}
	return num
}
