package loop

// 长方形分割成最大正方形
func dividing_to_max_square(length int, width int) int {
	if length%width == 0 {
		return width
	} else {
		return dividing_to_max_square(width, length%width)
	}
}
