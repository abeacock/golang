package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{2, 3, 5},
		{5, 3, 8},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, expected: %d", table.x, table.y, total, table.n)
		}
	}
}