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
				if bit >= 97 {
					total += int(bit - 'a' + 1)

				} else {
					total += int(bit - 'A' + 27)
				}
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

				f := m[bit]
				f[i] = 1
				m[bit] = f
			}

		}

		for k, v := range m {
			c := 0
			for i := range v {
				c += v[i]
			}

			if c == 3 {
				if k >= 97 {
					total2 += int(k - 'a' + 1)

				} else {
					total2 += int(k - 'A' + 27)
				}
			}

		}
	}
	return total, total2
}
