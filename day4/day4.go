package day4

import (
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type board [][]int

func (b *board) getScore() int {
	score := 0
	for x := range *b {
		for y := range (*b)[x] {
			if (*b)[x][y] != -1 {
				score += (*b)[x][y]
			}
		}
	}
	return score
}

func (b *board) solve(n int) (bool, int) {
	for x := range *b {
		for y := range (*b)[x] {
			if (*b)[x][y] == n {
				(*b)[x][y] = -1

				// check line & column
				sumX, sumY := 0, 0
				for i := 0; i < 5; i++ {
					sumX += (*b)[x][i]
					sumY += (*b)[i][y]
				}
				if sumX == -5 || sumY == -5 {
					return true, b.getScore() * n
				}
			}
		}
	}
	return false, 0
}

func convertToIntSlice(strNumbers []string) []int {
	numbers := make([]int, 0)
	for _, strNum := range strNumbers {
		n := util.MustAtoi(strNum)
		numbers = append(numbers, n)
	}
	return numbers
}

func parseInput(lines []string) (drawNumbers []int, boards []board) {
	drawNumbers = convertToIntSlice(strings.Split(lines[0], ","))
	boards = make([]board, 0)

	for i := 2; i < len(lines); i += 6 {
		b := make(board, 0)
		for j := i; j < i+5; j++ {
			boardLine := convertToIntSlice(strings.Fields(lines[j]))
			b = append(b, boardLine)
		}
		boards = append(boards, b)
	}

	return
}

func exe2(lines []string) int {
	drawNumbers, boards := parseInput(lines)

	lastScore := -1
	idxDone := make(map[int]struct{})

	for _, n := range drawNumbers {
		for i, b := range boards {
			if _, ok := idxDone[i]; !ok {
				if ok, score := b.solve(n); ok {
					idxDone[i] = struct{}{}
					lastScore = score
				}
			}
		}
	}

	return lastScore
}

func exe1(lines []string) int {
	drawNumbers, boards := parseInput(lines)

	for _, n := range drawNumbers {
		for _, b := range boards {
			if ok, score := b.solve(n); ok {
				return score
			}
		}
	}
	// this return should only occur if no board is the winner
	return -1
}
