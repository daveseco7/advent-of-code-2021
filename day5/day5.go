package day1

import (
	"fmt"
	"github.com/daveseco7/advent-of-code-2021/util"
	"log"
	"math"
	"strconv"
	"strings"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day5/input1.txt"

func getMinMax(a, b int) (min int, max int) {
	if a > b {
		return b, a
	}
	return a, b
}

func parseFromCoordinateToInts(coordinate string) (int, int) {
	s := strings.Split(coordinate, ",")

	x, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal("error parsing input string to int")
	}

	y, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal("error parsing input string to int")
	}

	return x, y
}

func exe1(lines []string) (counter int) {
	points := make(map[string]int)

	for _, line := range lines {
		pointsInput := strings.Split(line, " -> ")

		x1, y1 := parseFromCoordinateToInts(pointsInput[0])
		x2, y2 := parseFromCoordinateToInts(pointsInput[1])

		if x1 == x2 || y1 == y2 {
			x1, x2 = getMinMax(x1, x2)
			y1, y2 = getMinMax(y1, y2)
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
			maxAbs := math.Abs(float64(x2) - float64(x1))

			for i := 0; i < int(maxAbs)+1; i++ {
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

func Run() {
	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		panic(err)
	}

	//exe1 5608
	//exe2 20299
	fmt.Println(exe1(lineArray))
}
