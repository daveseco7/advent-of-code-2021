package day12

import (
	"strings"
)

const filePath = "input1.txt"

const (
	Start = "start"
	End   = "end"
)

type path []string

func (p *path) currentPath(cave string) path {
	currentPath := make(path, 0)
	for _, v := range *p {
		currentPath = append(currentPath, v)
	}
	currentPath = append(currentPath, cave)

	return currentPath
}

type paths map[string]path

func (p *paths) findAllPaths(canDoubleVisit bool, start, end string, visited path) []path {
	if start == end {
		return []path{{start}}
	}

	currentPath := visited.currentPath(start)

	isSecondVisit := strings.ToLower(start) == start && hasBeenVisited(start, visited)

	visitable := make([]string, 0)
	if p, ok := (*p)[start]; ok {
		for _, v := range p {
			if v == Start {
				continue
			}

			if strings.ToUpper(v) == v || !hasBeenVisited(v, visited) || (!isSecondVisit && canDoubleVisit) {
				visitable = append(visitable, v)
			}
		}
	}

	allPaths := make([]path, 0)
	for _, v := range visitable {
		solutions := p.findAllPaths(!isSecondVisit && canDoubleVisit, v, end, currentPath)
		for _, s := range solutions {
			allPaths = append(allPaths, s)
		}
	}

	return allPaths
}

func hasBeenVisited(s string, visited []string) bool {
	for _, i := range visited {
		if i == s {
			return true
		}
	}

	return false
}

func parseInput(lines []string) paths {
	p := make(paths)
	for _, line := range lines {
		s := strings.Split(line, "-")
		from := s[0]
		to := s[1]
		if _, ok := p[from]; !ok {
			p[from] = make([]string, 0, 1)
		}

		if _, ok := p[to]; !ok {
			p[to] = make([]string, 0, 1)
		}

		p[from] = append(p[from], to)
		p[to] = append(p[to], from)
	}

	return p
}

func exe2(lines []string) int {
	p := parseInput(lines)

	solutions := p.findAllPaths(true, Start, End, make([]string, 0))

	return len(solutions)
}

func exe1(lines []string) int {
	p := parseInput(lines)

	solutions := p.findAllPaths(false, Start, End, make([]string, 0))

	return len(solutions)
}
