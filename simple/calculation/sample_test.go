package calculation

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		x := 2
		y := 5
		result := Add(x, y)
		if result != 7 {
			t.Error("result should be 7")
		}
	})

	t.Run("failed", func(t *testing.T) {
		x := 2
		_ = 5
		result := Add(x)
		if result != 7 {
			t.Errorf("result should be 7")
		}
	})
}

func TestSub(t *testing.T) {
	x := 5
	y := 2
	result := Sub(x, y)
	if result != 3 {
		t.Error("result should be 3")
	}
}
