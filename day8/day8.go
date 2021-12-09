package day8

import (
	"fmt"
	"strings"
)

const filePath = "input1.txt"

type signals []string

func (s *signals) findBaseNumbers() map[string]int {
	//base numbers are 1, 4, 7, 8
	digitTranslation := make(map[string]int)
	for _, p := range *s {
		switch len(p) {
		case 2:
			digitTranslation[p] = 1
		case 4:
			digitTranslation[p] = 4
		case 3:
			digitTranslation[p] = 7
		case 7:
			digitTranslation[p] = 8
		}
	}
	return digitTranslation
}

type entry struct {
	i signals
	o signals
}

func toSignals(slice []string) signals {
	i := make(signals, 0)
	for _, s := range slice {
		i = append(i, s)
	}
	return i
}

func parseInput(lines []string) []entry {
	entries := make([]entry, 0)
	for _, line := range lines {
		data := strings.Split(line, "|")
		entries = append(entries,
			entry{
				i: toSignals(strings.Fields(data[0])),
				o: toSignals(strings.Fields(data[1])),
			},
		)
	}

	return entries
}

func getOne(m map[string]int) string {
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return ""
}

func getFour(m map[string]int) string {
	for k, v := range m {
		if v == 4 {
			return k
		}
	}
	return ""
}

func haveSameContent(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, r := range a {
		if !strings.Contains(b, string(r)) {
			return false
		}
	}

	return true
}

func exe2(lines []string) (count int) {
	entries := parseInput(lines)
	for _, e := range entries {
		dt := e.i.findBaseNumbers()
		for _, p := range e.i {
			one := getOne(dt)
			four := getFour(dt)

			l4 := strings.ReplaceAll(four, string(one[0]), "")
			l4 = strings.ReplaceAll(l4, string(one[1]), "")

			if len(p) == 2 || len(p) == 4 || len(p) == 3 || len(p) == 7 {
				continue
			}

			if len(p) == 5 {
				if strings.Contains(p, string(one[0])) && strings.Contains(p, string(one[1])) {
					// its a 3
					dt[p] = 3
				} else if strings.Contains(p, string(l4[0])) && strings.Contains(p, string(l4[1])) {
					// its a 5
					dt[p] = 5
				} else {
					// its a 2
					dt[p] = 2
				}
			} else {
				if !(strings.Contains(p, string(one[0])) && strings.Contains(p, string(one[1]))) {
					// it's a 6
					dt[p] = 6
				} else if strings.Contains(p, string(four[0])) && strings.Contains(p, string(four[1])) && strings.Contains(p, string(four[2])) && strings.Contains(p, string(four[3])) {
					// it's a 9
					dt[p] = 9
				} else {
					// it's a 0
					dt[p] = 0
				}
			}
		}
		decOutput := 0
		for _, p := range e.o {
			for k, v := range dt {
				if haveSameContent(p, k) {
					decOutput = (decOutput * 10) + v
				}
			}
		}

		count += decOutput
	}

	return count
}

func exe1(lines []string) (count int) {
	entries := parseInput(lines)

	/*
		len represents the number of segments that are on so:
			segments    digit
				2         1
				4		  4
				3		  7
				7		  8
	*/
	for _, e := range entries {
		digitTranslation := e.o.findBaseNumbers()
		for _, pattern := range e.o {
			if _, ok := digitTranslation[pattern]; ok {
				count++
			}
		}
	}

	fmt.Println(count)

	return count
}
