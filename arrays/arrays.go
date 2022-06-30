package arrays

import (
	"strconv"
	"strings"
)

/**
提供各种转换操作
*/

func Int2str(arry []int) []string {
	strArray := make([]string, len(arry))
	for i, _ := range arry {
		strArray[i] = strconv.Itoa(arry[i])
	}
	return strArray
}

func Str2int(arry []string) []int {
	intArray := make([]int, len(arry))
	for i, _ := range arry {
		ai, err := strconv.Atoi(arry[i])
		if err != nil {
			panic("转换发生异常：" + arry[i] + "不能转为int值！")
		}
		intArray[i] = ai
	}
	return intArray
}

func IntJoinStrDot(array []int) string {
	return IntJoinStr(array, ",")
}

func IntJoinStr(arry []int, delimiter string) string {
	return strings.Join(Int2str(arry), delimiter)
}
