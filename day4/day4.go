package day1

import (
	"fmt"
	"github.com/daveseco7/advent-of-code-2021/util"
	"log"
	"strconv"
	"strings"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day4/input1.txt"

type board [][]int

func (b *board) print() {
	for _, i := range *b {
		fmt.Println(i)
	}
	fmt.Println()
}

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
		n, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal("error parsing number input")
		}
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

func Run() {
	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		panic(err)
	}

	//exe1 65325
	//exe2 4624
	fmt.Println(exe1(lineArray))
	fmt.Println(exe2(lineArray))
}
