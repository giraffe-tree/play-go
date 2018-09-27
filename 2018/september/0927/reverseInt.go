package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
https://leetcode-cn.com/problems/reverse-integer/description/
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
 */
func main() {
	fmt.Println(reverseCopy(123))
}

/**
copy from leetcode
 */
func reverseCopy(x int) int {
	var reply string
	z:= int(math.Abs(float64(x)))
	// 这里不能使用 string(z) z会被转为 ascii 码
	temp := strconv.Itoa(z)
	var last string
	for i := len(temp) - 1; i >= 0; i-- {
		v := string(temp[i])
		if v == "0" && last == "" {
			continue
		}
		last = v
		reply = reply + v
	}
	r, _ := strconv.Atoi(reply)
	if x < 0 {
		r = -r
		if r < -2147483648 {
			return 0
		}
	} else {
		if r > 2147483647 {
			return 0
		}
	}
	return r
}

/**
取余反转 8ms
 */
func reverse(x int) int {
	flag := true
	if x < 0 {
		flag = false
		x = -x
	}

	length := 0
	arr := make([]int, 16)
	for x > 0 {
		i := x % 10
		arr[length] = i
		length++
		x = x / 10
	}
	result := 0
	for j := 0; j < length; j++ {
		result = result*10 + arr[j]
	}
	if flag && result <= 2147483647 {
		return result
	} else if !flag && result <= 2147483648 {

		return -result
	}
	return 0
}

// 2147483648
// 9646324351
