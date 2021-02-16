package search

// 二分查找
func binarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		val := list[mid]
		if val == target {
			return mid
		}
		if val < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
