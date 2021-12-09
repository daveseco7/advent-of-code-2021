package day9

import (
	"log"
	"sort"
	"strconv"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type coordinates struct {
	x, y int
}

type heightmap [][]int

func (hm *heightmap) lowPoints() (lowPoints []coordinates, risk int) {
	for i := 0; i < len(*hm); i++ {
		for j := 0; j < len((*hm)[i]); j++ {
			lp := coordinates{x: i, y: j}
			if r, ok := hm.isLowPoint(lp); ok {
				lowPoints = append(lowPoints, lp)
				risk += r
			}
		}
	}

	return
}

func (hm *heightmap) isLowPoint(lp coordinates) (risk int, ok bool) {
	ok = true
	h := (*hm)[lp.x][lp.y]

	for _, a := range hm.getAdjacent(lp) {
		if (*hm)[a.x][a.y] <= h {
			ok = false
		}
	}
	return h + 1, ok
}

func (hm *heightmap) getAdjacent(lp coordinates) (adjacent []coordinates) {
	for i := lp.x - 1; i <= lp.x+1; i++ {
		for j := lp.y - 1; j <= lp.y+1; j++ {
			// ignore diagonals
			if util.Abs(lp.x-i)+util.Abs(lp.y-j) > 1 {
				continue
			}

			// ignore out of bounds
			if i >= len(*hm) || i < 0 || j >= len((*hm)[i]) || j < 0 {
				continue
			}

			// ignore itself
			if i == lp.x && j == lp.y {
				continue
			}
			adjacent = append(adjacent,
				coordinates{
					x: i,
					y: j,
				},
			)
		}
	}
	return adjacent
}

func (hm *heightmap) calcBasin(lp coordinates, visited map[coordinates]struct{}) (basinSize int) {
	basinSize = 1
	visited[lp] = struct{}{}

	adjs := hm.getAdjacent(lp)
	for _, adjacent := range adjs {
		if _, ok := visited[adjacent]; !ok {
			visited[adjacent] = struct{}{}
			if (*hm)[adjacent.x][adjacent.y] != 9 {
				basinSize += hm.calcBasin(adjacent, visited)
			}
		}
	}
	return basinSize
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

	lowPoints, _ := h.lowPoints()

	basinSizes := make([]int, 0)
	visited := make(map[coordinates]struct{})

	for _, lp := range lowPoints {
		basinSize := h.calcBasin(lp, visited)
		basinSizes = append(basinSizes, basinSize)
	}

	sort.SliceStable(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func exe1(lines []string) int {
	h := parseInput(lines)

	_, risk := h.lowPoints()

	return risk
}
