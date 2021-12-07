package day7

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day7/input1.txt"

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

	minCost := math.MaxInt
	for i := util.Min(p); i <= util.Max(p); i++ {
		cost := 0
		for _, v := range p {
			posDiff := util.Abs(v - i)
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
		distance += util.Abs(v - median)
	}

	return distance
}
