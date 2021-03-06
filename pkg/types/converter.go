package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

//Uint64ToString 将uint64转换为string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num,10)
}

//StringToInt 将字符串转换为int
func StringToInt(str string) int {
	//转化为整型
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}