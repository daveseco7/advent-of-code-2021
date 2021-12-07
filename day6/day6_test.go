package day6

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2021/util"
)

func TestExe1(t *testing.T) {
	const expected = 353274

	// some method implements now exercise 2, so ignore it on this test
	t.Skip()

	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}

func TestExe2(t *testing.T) {
	const expected = 1609314870967

	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}
