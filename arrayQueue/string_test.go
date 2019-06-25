package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	t.Log(IntSlice2String(s))

}

//将[]int 转换成逗号分隔的字符串
func IntSlice2String(s interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", ",", -1)
}
