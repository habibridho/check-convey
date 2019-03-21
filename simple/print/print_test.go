package print

import (
	"testing"
)

func TestPrint(t *testing.T) {
	s := "some string"
	result := Prt(s)
	if result != s {
		t.Error("result should be equal with s")
	}
}
