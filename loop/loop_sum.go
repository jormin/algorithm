package loop

// 递归版求和
func loopSum(list []int) int {
	switch len(list) {
	case 0:
		return 0
	case 1:
		return list[0]
	default:
		return list[0] + loopSum(list[1:])
	}
}
