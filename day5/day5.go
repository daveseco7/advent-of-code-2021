package day5

import (
	"fmt"
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

func parseFromCoordinateToInts(coordinate string) (int, int) {
	s := strings.Split(coordinate, ",")

	return util.MustAtoi(s[0]), util.MustAtoi(s[1])
}

func exe(lines []string) (counter int) {
	points := make(map[string]int)

	for _, line := range lines {
		pointsInput := strings.Split(line, " -> ")

		x1, y1 := parseFromCoordinateToInts(pointsInput[0])
		x2, y2 := parseFromCoordinateToInts(pointsInput[1])

		if x1 == x2 || y1 == y2 {
			x1, x2 = util.Min([]int{x1, x2}), util.Max([]int{x1, x2})
			y1, y2 = util.Min([]int{y1, y2}), util.Max([]int{y1, y2})
			for x := x1; x < x2+1; x++ {
				for y := y1; y < y2+1; y++ {
					points[fmt.Sprintf("(%d,%d)", x, y)]++
				}
			}
		} else {
			// comment out to remove diagonals
			dx := -1
			if x2 > x1 {
				dx = 1
			}
			dy := -1
			if y2 > y1 {
				dy = 1
			}
			maxAbs := util.Abs(x2 - x1)
			for i := 0; i < maxAbs+1; i++ {
				points[fmt.Sprintf("(%d,%d)", x1+i*dx, y1+i*dy)]++
			}
		}
	}

	for _, v := range points {
		if v > 1 {
			counter++
		}
	}

	return counter
}
