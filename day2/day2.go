package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day2/input1.txt"

// string mapping

type direction string

const (
	forward direction = "forward"
	down              = "down"
	up                = "up"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *position) product() int {
	return p.horizontal * p.depth
}

func (p *position) move(dir direction, value int) error {
	switch dir {
	case forward:
		p.horizontal += value
		p.depth += p.aim * value
	case down:
		p.aim += value
	case up:
		p.aim -= value
	default:
		return fmt.Errorf("invalid move direction: %s", dir)
	}
	return nil
}

func parseMoveOperation(line string) (direction, int, error) {
	slice := strings.Split(line, " ")
	if len(slice) != 2 {
		return "", 0, fmt.Errorf(" invalid move direction: %s", line)
	}
	n, err := strconv.Atoi(slice[1])
	if err != nil {
		return "", 0, fmt.Errorf(" invalid move intensity: %s", slice[1])
	}

	return direction(slice[0]), n, nil
}

func exe1(lines []string) int {
	p := position{0, 0, 0}

	for _, l := range lines {
		dir, intensity, err := parseMoveOperation(l)
		if err != nil {
			log.Fatalf("error in exer1: %v", err)
		}
		if err := p.move(dir, intensity); err != nil {
			log.Fatalf("error in exer1: %v", err)
		}
	}

	return p.product()
}

func Run() {
	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		panic(err)
	}

	//exe1 1604850
	//exe2 1685186100
	fmt.Println(exe1(lineArray))
}
