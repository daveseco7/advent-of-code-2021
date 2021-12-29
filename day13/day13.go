package day13

import (
	"fmt"
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type paper [][]int

func (p *paper) print() {
	for y := 0; y < len(*p); y++ {
		for x := 0; x < len((*p)[y]); x++ {
			if (*p)[y][x] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func (p *paper) countPoints() (points int) {
	for y := 0; y < len(*p); y++ {
		for x := 0; x < len((*p)[y]); x++ {
			if (*p)[y][x] == 1 {
				points++
			}
		}
	}

	return points
}

func newPaper(maxX, maxY int) paper {
	p := make(paper, maxY+1)
	for y := 0; y <= maxY; y++ {
		xSlice := make([]int, maxX+1)
		for x := 0; x <= maxX; x++ {
			xSlice[x] = 0
		}
		p[y] = xSlice
	}

	return p
}

type coordinate struct {
	x, y int
}

type fold struct {
	val  int
	axis string
}

func (f *fold) fold(p paper) paper {
	if f.axis == "x" {
		return f.foldX(p)
	}
	return f.foldY(p)
}

func (f *fold) foldX(p paper) paper {
	newPaper := make(paper, 0)
	for y := 0; y < len(p); y++ {
		xSlice := make([]int, 0)
		for x := 0; x < f.val; x++ {
			point := p[y][x] | p[y][len(p[y])-x-1]
			xSlice = append(xSlice, point)
		}
		newPaper = append(newPaper, xSlice)
	}

	return newPaper
}

func (f *fold) foldY(p paper) paper {
	newPaper := make(paper, 0)
	for y := 0; y < f.val; y++ {
		xSlice := make([]int, 0)
		for x := 0; x < len(p[y]); x++ {
			xSlice = append(xSlice, p[y][x])
		}
		newPaper = append(newPaper, xSlice)
	}

	for y := f.val + 1; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			newPaper[2*f.val-y][x] |= p[y][x]
		}
	}
	return newPaper
}

type instructions struct {
	paper paper
	folds []fold
}

func parseFold(line string) fold {
	strs := strings.Fields(line)
	s := strings.Split(strs[2], "=")

	return fold{
		val:  util.MustAtoi(s[1]),
		axis: s[0],
	}
}

func parseCoordinate(line string) coordinate {
	s := strings.Split(line, ",")

	return coordinate{
		x: util.MustAtoi(s[0]),
		y: util.MustAtoi(s[1]),
	}
}

func parseInput(lines []string) instructions {
	maxX, maxY := 0, 0
	f := make([]fold, 0)
	pointHolder := make(map[string]coordinate)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			f = append(f, parseFold(line))
			continue
		}

		c := parseCoordinate(line)

		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}

		pointHolder[line] = c
	}

	p := newPaper(maxX, maxY)
	for _, v := range pointHolder {
		p[v.y][v.x] = 1
	}

	return instructions{
		paper: p,
		folds: f,
	}
}

func exe2(lines []string) int {
	inst := parseInput(lines)

	p := inst.paper
	for _, f := range inst.folds {
		p = f.fold(p)
	}

	p.print()

	// JRZBLGKH
	return 0
}

func exe1(lines []string) int {
	inst := parseInput(lines)
	newPaper := inst.folds[0].fold(inst.paper)
	return newPaper.countPoints()
}
