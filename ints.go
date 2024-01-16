package helper

import (
	"strconv"
	"strings"
)

func StrToInts(s string, sep string) []int {
	if s == "" {
		return []int{}
	}

	items := strings.Split(s, sep)
	itemLen := len(items)
	ret := make([]int, 0, itemLen)
	for i := range items {
		item := strings.TrimSpace(items[i])
		if item == "" {
			continue
		}
		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		ret = append(ret, val)
	}
	return ret
}

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
	if num <= 0 {
		return false
	}
	return num&(num-1) == 0
	//return num & -num == num
}

// 超过半数

func MoreThanHalf(num int) int {
	if num <= 0 {
		return 0
	}
	return num/2 + 1
}
