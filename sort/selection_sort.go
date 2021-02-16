package sort

// 选择排序
func selectionSort(list []int) []int {
	var res []int
	num := len(list)
	for i := 0; i < num; i++ {
		smallIndex := findSmallest(list)
		// 添加当前最小元素到结果数组中
		res = append(res, list[smallIndex])
		// 删除最小元素
		list = append(list[:smallIndex], list[smallIndex+1:]...)
	}
	return res
}

// 查找数组中的最小元素
func findSmallest(list []int) int {
	small := list[0]
	smallIndex := 0
	for i := 0; i < len(list); i++ {
		if list[i] < small {
			small = list[i]
			smallIndex = i
		}
	}
	return smallIndex
}
