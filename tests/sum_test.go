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
