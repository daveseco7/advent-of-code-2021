package day1

import (
	"fmt"
	"github.com/daveseco7/advent-of-code-2021/util"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day7/input1.txt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(lines []string) (positions []int) {
	positions = make([]int, 0)
	for _, line := range lines {
		for _, str := range strings.Split(line, ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal("invalid number in the puzzle input for exe 1")
			}
			positions = append(positions, n)
		}
	}
	return positions
}

func exe2(lines []string) int {
	p := parseInput(lines)

	min := p[0]
	for _, v := range p {
		if v < min {
			min = v
		}
	}

	max := 0
	for _, v := range p {
		if v > max {
			max = v
		}
	}

	minCost := math.MaxInt
	for i := min; i <= max; i++ {
		cost := 0
		for _, v := range p {
			posDiff := abs(v - i)
			cost += posDiff * (posDiff + 1) / 2
		}

		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func exe1(lines []string) (distance int) {
	p := parseInput(lines)

	sort.SliceStable(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	median := p[len(p)/2]
	for _, v := range p {
		distance += abs(v - median)
	}

	return distance
}

func Run() {
	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		panic(err)
	}

	//exe1 347509
	//exe2 98257206
	fmt.Println(exe1(lineArray))
	fmt.Println(exe2(lineArray))
}
