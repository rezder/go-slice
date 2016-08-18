package slices

func Contain(list []int, value int) (con bool) {
	for _, v := range list {
		if v == value {
			con = true
			break
		}
	}
	return con
}

//Equal test for equal slices
// This is a test for elements is in the same order and equal amounts in both exist.
// Cap. does not matter and empty and nil is equal because the have the same numbers of elements
// and
func Equal(list1 []int, list2 []int) (equal bool) {
	if len(list1) == len(list2) {
		equal = true
		if len(list1) > 0 {
			for i, v := range list1 {
				if v != list2[i] {
					equal = false
					break
				}
			}
		}
	}
	return equal
}
func WithOutNew(list []int, out []int) (res []int) {
	res = make([]int, 0, len(list))
	for _, v := range list {
		if !Contain(out, v) {
			res = append(res, v)
		}
	}
	return res
}
func WithOut(list []int, out []int) (res []int) {
	res = list[:]
	for i := 0; i < len(res); {
		if Contain(out, res[i]) {
			res = res[:i+copy(res[i:], res[i+1:])]
		} else {
			i++
		}
	}
	return res
}
func Remove(list []int, out int) (res []int) {
	res = make([]int, 0, len(list))
	for _, v := range list {
		if out != v {
			res = append(res, v)
		}
	}
	return res
}

//AddSorted add a element to a sorted list keeping the sort
//dec true if decending sort
func AddSorted(list []int, in int, dec bool) []int {
	if len(list) == 0 {
		list = append(list, in)
	} else {
		index := len(list)
		if dec {
			for i, v := range list {
				if in >= v {
					index = i
					break
				}
			}
		} else {
			for i, v := range list {
				if in <= v {
					index = i
					break
				}
			}
		}
		list = append(list, 0)
		copy(list[index+1:], list[index:])
		list[index] = in
	}
	return list
}
