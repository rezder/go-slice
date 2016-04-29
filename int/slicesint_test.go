package slices

import (
	"testing"
)

func TestSlices(t *testing.T) {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	if !Equal(s1, s2) {
		t.Errorf("Empty failed")
	}
}
