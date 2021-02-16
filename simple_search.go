package algorithm

// 简单查找
func simpleSearch(list []int, target int) int {
	for index, val := range list {
		if val == target {
			return index
		}
	}
	return -1
}
