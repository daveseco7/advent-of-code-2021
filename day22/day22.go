package day22

import (
	"regexp"
	"sort"
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

const stepRegionRegex = `(\w+) x=(-?\d+..-?\d+),y=(-?\d+..-?\d+),z=(-?\d+..-?\d+)`

const (
	gridMinLimit = -50
	gridMaxLimit = 50
)

type step struct {
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
	state      bool
}

type axisHelper struct {
	s *[]int
	m *map[int]int
}

type cuboid struct {
	xx   axisHelper
	yy   axisHelper
	zz   axisHelper
	grid [][][]bool
}

func newCuboids(steps []step) cuboid {
	xs, ys, zs := make([]int, 0), make([]int, 0), make([]int, 0)
	xm, ym, zm := make(map[int]int), make(map[int]int), make(map[int]int)

	for _, s := range steps {
		xs = append(xs, s.minX, s.maxX+1)
		ys = append(ys, s.minY, s.maxY+1)
		zs = append(zs, s.minZ, s.maxZ+1)
	}

	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)

	for i, x := range xs {
		xm[x] = i
	}

	for j, y := range ys {
		ym[y] = j
	}

	for k, z := range zs {
		zm[z] = k
	}

	return cuboid{
		xx: axisHelper{
			s: &xs,
			m: &xm,
		},
		yy: axisHelper{
			s: &ys,
			m: &ym,
		},
		zz: axisHelper{
			s: &zs,
			m: &zm,
		},
		grid: initGrid(len(xs), len(ys), len(zs)),
	}
}

func (c *cuboid) volume() (volume int) {
	for x := 0; x < len(*c.xx.s)-1; x++ {
		for y := 0; y < len(*c.yy.s)-1; y++ {
			for z := 0; z < len(*c.zz.s)-1; z++ {
				if c.grid[x][y][z] {
					volume += ((*c.xx.s)[x+1] - (*c.xx.s)[x]) * ((*c.yy.s)[y+1] - (*c.yy.s)[y]) * ((*c.zz.s)[z+1] - (*c.zz.s)[z])
				}
			}
		}
	}

	return volume
}

func (c *cuboid) setState(s step) {
	xStart, xEnd := (*c.xx.m)[s.minX], (*c.xx.m)[s.maxX+1]-1
	yStart, yEnd := (*c.yy.m)[s.minY], (*c.yy.m)[s.maxY+1]-1
	zStart, zEnd := (*c.zz.m)[s.minZ], (*c.zz.m)[s.maxZ+1]-1

	for x := xStart; x <= xEnd; x++ {
		for y := yStart; y <= yEnd; y++ {
			for z := zStart; z <= zEnd; z++ {
				c.grid[x][y][z] = s.state
			}
		}
	}
}

func initGrid(lenXX, lenYY, lenZZ int) [][][]bool {
	grid := make([][][]bool, lenXX)
	for x := 0; x < lenXX; x++ {
		grid[x] = make([][]bool, lenYY)
		for y := 0; y < lenYY; y++ {
			grid[x][y] = make([]bool, lenZZ)
		}
	}

	return grid
}

func parseInterval(l string) (int, int) {
	numbers := strings.Split(l, "..")

	return util.MustAtoi(numbers[0]), util.MustAtoi(numbers[1])
}

func parseInput(lines []string, ignore func(int, int, int, int, int, int) bool) []step {
	r := regexp.MustCompile(stepRegionRegex)

	steps := make([]step, 0)
	for _, line := range lines {
		result := r.FindAllStringSubmatch(line, -1)

		minX, maxX := parseInterval(result[0][2])
		minY, maxY := parseInterval(result[0][3])
		minZ, maxZ := parseInterval(result[0][4])

		var state bool
		if result[0][1] == "on" {
			state = true
		}

		if !ignore(minX, maxX, minY, maxY, minZ, maxZ) {
			steps = append(steps, step{
				state: state,
				minX:  minX,
				maxX:  maxX,
				minY:  minY,
				maxY:  maxY,
				minZ:  minZ,
				maxZ:  maxZ,
			})
		}
	}

	return steps
}

func exe2(lines []string) (volume int) {
	steps := parseInput(lines, func(x1, x2, y1, y2, z1, z2 int) bool { return false })

	c := newCuboids(steps)

	for _, s := range steps {
		c.setState(s)
	}

	return c.volume()
}

func exe1(lines []string) (volume int) {
	steps := parseInput(lines,
		func(x1, x2, y1, y2, z1, z2 int) bool {
			if x1 >= gridMinLimit && x2 <= gridMaxLimit &&
				y1 >= gridMinLimit && y2 <= gridMaxLimit &&
				z1 >= gridMinLimit && z2 <= gridMaxLimit {
				return false
			}
			return true
		},
	)

	c := newCuboids(steps)

	for _, s := range steps {
		c.setState(s)
	}

	return c.volume()
}
