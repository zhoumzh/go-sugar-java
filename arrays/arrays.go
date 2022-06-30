package arrays

import (
	"strconv"
	"strings"
)

/**
提供各种转换操作
*/

func int2str(arry []int) []string {
	strArray := make([]string, len(arry))
	for i, _ := range arry {
		strArray[i] = strconv.Itoa(arry[i])
	}
	return strArray
}

func str2int(arry []string) []int {
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

func intJoinStrDot(array []int) string {
	return intJoinStr(array, ",")
}

func intJoinStr(arry []int, delimiter string) string {
	return strings.Join(int2str(arry), delimiter)
}
