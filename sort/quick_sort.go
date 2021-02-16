package sort

// 快速排序
func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]
	var less []int
	var greater []int
	for _, v := range list {
		if v < pivot {
			less = append(less, v)
		}
		if v > pivot {
			greater = append(greater, v)
		}
	}
	var res []int
	res = append(res, quickSort(less)...)
	res = append(res, pivot)
	res = append(res, quickSort(greater)...)
	return res
}
