package day3

import (
	"log"
	"math"
	"strconv"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type counter map[int]int

func (c *counter) inc(key int) {
	if c == nil {
		*c = make(map[int]int)
	}
	(*c)[key]++
}

func (c *counter) getBit(dominant bool) int {
	if dominant {
		if (*c)[0] > (*c)[1] {
			return 0
		}
		return 1
	}

	if (*c)[0] > (*c)[1] {
		return 1
	}
	return 0
}

func calcGamma(columnLength int, lines []string) (gamma int) {
	idxSet := getIndexes(lines)

	for i := 0; i < columnLength; i++ {
		bit, _, _ := calcNextBit(lines, i, true, idxSet)
		gamma = (gamma << 1) | bit
	}

	return
}

func calcNextBit(lines []string, column int, isDominant bool, idxSet util.IntSet) (bit int, idx0, idx1 []int) {
	c := counter{}
	for _, idx := range idxSet.GetValues() {
		// int() directive will interpret the value as the rune ASCII value, subtract 0 ASCII value.
		// we may have flaky behaviour for unexpected string inputs
		v := int(lines[idx][column] - '0')
		c.inc(v)

		if v == 0 {
			idx0 = append(idx0, idx)
		} else {
			idx1 = append(idx1, idx)
		}
	}

	return c.getBit(isDominant), idx0, idx1
}

func getIndexes(lines []string) (idxSet util.IntSet) {
	for i := 0; i < len(lines); i++ {
		idxSet.AddValue(i)
	}
	return
}

func deleteUnnecessaryNums(idxSet util.IntSet, bit int, idx0, idx1 []int) {
	if bit == 0 {
		idxSet.Delete(idx1...)
	} else {
		idxSet.Delete(idx0...)
	}
}

func calculateBits(lines []string, numOfBits int, isDominant bool, idxSet util.IntSet) (int, error) {
	result := 0
	for i := 0; i < numOfBits; i++ {
		if len(idxSet) == 1 {
			// found value, stop and return such value
			r, err := strconv.ParseInt(lines[idxSet.GetValues()[0]], 2, 64)
			if err != nil {
				return 0, err
			}

			return int(r), nil
		}

		bit, idx0, idx1 := calcNextBit(lines, i, isDominant, idxSet)
		result = (result << 1) | bit
		deleteUnnecessaryNums(idxSet, bit, idx0, idx1)
	}
	return result, nil
}

func exe2(lines []string) int {
	numOfBits := len(lines[0])

	oxygenGeneratorRating, err := calculateBits(lines, numOfBits, true, getIndexes(lines))
	if err != nil {
		log.Fatalf("error in exer2: %v", err)
	}

	CO2ScrubberRating, err := calculateBits(lines, numOfBits, false, getIndexes(lines))
	if err != nil {
		log.Fatalf("error in exer2: %v", err)
	}

	return oxygenGeneratorRating * CO2ScrubberRating
}

func exe1(lines []string) int {
	numOfBits := len(lines[0])
	gamma := calcGamma(numOfBits, lines)

	// since gamma is the most common bit and epsilon the least common, they are each other's complement.
	// calculate gamma's complement
	maxUnsignedNumber := math.Pow(2, float64(numOfBits)) - 1
	epsilon := int(maxUnsignedNumber) - gamma

	return gamma * epsilon
}
