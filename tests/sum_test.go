package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum is incorrect. Expected %d, got %d", 10, total)
	}
}

func TestABunchOfSums(t *testing.T) {
	tables := []struct {
		x   int32
		y   int32
		sum int64
	}{
		{1, 1, 2},
		{2, 4, 6},
		{100, 900, 1000},
		{2000000000, 2000000000, 4000000000},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.sum {
			t.Errorf("Sum of %d and %d was incorrect. Expected %d, got %d", table.x, table.y, table.sum, total)
		}
	}
}
