package day6

import (
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type school []int

func (s *school) incrementDay() {
	zeroValHolder := (*s)[0]

	(*s)[0] = (*s)[1]
	(*s)[1] = (*s)[2]
	(*s)[2] = (*s)[3]
	(*s)[3] = (*s)[4]
	(*s)[4] = (*s)[5]
	(*s)[5] = (*s)[6]
	(*s)[6] = zeroValHolder + (*s)[7]
	(*s)[7] = (*s)[8]
	(*s)[8] = zeroValHolder
}

func exe(lines []string) (counter int) {
	s := make(school, 10)

	initialState := strings.Split(lines[0], ",")
	for _, state := range initialState {
		n := util.MustAtoi(state)
		s[n]++
	}

	for i := 0; i < 256; i++ {
		s.incrementDay()
	}

	for _, v := range s {
		counter += v
	}

	return counter
}
