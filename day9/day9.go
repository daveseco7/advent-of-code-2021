package day9

import (
	"fmt"
	"github.com/daveseco7/advent-of-code-2021/util"
	"log"
	"strconv"
)

const filePath = "input1.txt"

type heightmap [][]int

func (hm *heightmap) isLowPoint(x, y int) (risk int, ok bool) {
	ok = true
	h := (*hm)[x][y]

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// ignore diagonals
			if util.Abs(x-i)+util.Abs(y-j) > 1 {
				continue
			}
			// ignore out of bounds
			if i >= len(*hm) || i < 0 || j >= len((*hm)[i]) || j < 0 {
				continue
			}

			// ignore itself
			if i == x && j == y {
				continue
			}

			// if there's at least one smaller neighbor
			if (*hm)[i][j] <= h {
				ok = false
			}
		}
	}

	return h + 1, ok
}

func parseInput(lines []string) heightmap {
	m := make(heightmap, 0)
	for _, strLine := range lines {
		intLine := make([]int, 0)

		for _, r := range strLine {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("invalid number in the puzzle input")
			}

			intLine = append(intLine, n)
		}
		m = append(m, intLine)
	}
	return m
}

func exe2(lines []string) int {
	h := parseInput(lines)

	risk := 0
	for i := 0; i < len(h); i++ {
		for j := 0; j < len(h[i]); j++ {
			if r, ok := h.isLowPoint(i, j); ok {
				risk += r
			}
		}
	}

	fmt.Println(risk)
	return risk
}

func exe1(lines []string) int {
	h := parseInput(lines)

	risk := 0
	for i := 0; i < len(h); i++ {
		for j := 0; j < len(h[i]); j++ {
			if r, ok := h.isLowPoint(i, j); ok {
				risk += r
			}
		}
	}

	return risk
}
