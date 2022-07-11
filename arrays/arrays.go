package arrays

import (
	"strconv"
	"strings"
)

/**
提供各种数组操作
*/

func Int2Str(array32 []int) []string {
	return manyInt2Str(nil, nil, array32, nil, nil)
}

func Int82Str(array8 []int8) []string {
	return manyInt2Str(array8, nil, nil, nil, nil)
}

func Int162Str(array16 []int16) []string {
	return manyInt2Str(nil, array16, nil, nil, nil)
}

func Int322Str(array32 []int32) []string {
	return manyInt2Str(nil, nil, nil, array32, nil)
}

func Int642str(array []int64) []string {
	return manyInt2Str(nil, nil, nil, nil, array)
}

func manyInt2Str(i8 []int8, i16 []int16, ia []int, i32 []int32, i64 []int64) []string {
	_len := len(i8) + len(i16) + len(ia) + len(i32) + len(i64)
	strArray := make([]string, _len)
	for i, _ := range ia {
		strArray[i] = strconv.Itoa(ia[i])
	}
	for i, _ := range i8 {
		strArray[i] = strconv.Itoa(int(i8[i]))
	}
	for i, _ := range i16 {
		strArray[i] = strconv.FormatInt(int64(i16[i]), 10)
	}
	for i, _ := range i32 {
		strArray[i] = strconv.FormatInt(int64(i32[i]), 10)
	}
	for i, _ := range i64 {
		strArray[i] = strconv.FormatInt(i64[i], 10)
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
	return strings.Join(Int2Str(arry), delimiter)
}
