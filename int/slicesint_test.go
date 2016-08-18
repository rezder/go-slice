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
func TestAddSort(t *testing.T) {
	acc := []int{1, 2, 3, 4}
	dec := []int{4, 3, 2, 1}

	test := []int{3, 2, 4, 1}
	var res []int
	res = nil
	for _, v := range test {
		res = AddSorted(res, v, true)
	}
	if !Equal(res, dec) {
		t.Errorf("Sort decending fail: %v", res)
	}
	res = nil
	for _, v := range test {
		res = AddSorted(res, v, false)
	}
	if !Equal(res, acc) {
		t.Errorf("Sort decending fail: %v", res)
	}
}
