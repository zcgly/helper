package helper

func UpperTo2Power(num int) int {
	//return 1 << uint(math.Ceil(math.Log2(float64(num))))
	num--
	num |= num >> 1
	num |= num >> 2
	num |= num >> 4
	num |= num >> 8
	num |= num >> 16
	return num + 1
}

func Is2Power(num int) bool {
	return num&(num-1) == 0
	//return num & -num == num
}

// 超过半数
func MoreThanHalf(num int) int {
	return num/2 + 1
}
