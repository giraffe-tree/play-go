package main

import "fmt"

/**
https://leetcode-cn.com/problems/zigzag-conversion/description/
将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：

P   A   H   N
A P L S I I G
Y   I   R
之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

实现一个将字符串进行指定行数变换的函数:

string convert(string s, int numRows);
示例 1:

输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
示例 2:

输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
func main() {
	fmt.Println(convertCopy("PAYPALISHIRING", 4))
	// PIN  ALSIGYAHR  PI
	// PIN  PINALSIGYAHRPI  PI
}

/**
copy from leetcode  12ms
好吧, 它直接建立了一个二维数组...

 */
func convertCopy(s string, numRows int) string {
	var z [][]byte
	if numRows == 1 {
		return s
	}
	strlen := len(s)
	z = make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		z[i] = make([]byte, 0)
	}
	row := 0
	currentFlag := true
	for i := 0; i < strlen; i++ {
		z[row] = append(z[row], byte(s[i]))
		if currentFlag == true && row == numRows-1 {
			currentFlag = false
		}
		if currentFlag == false && row == 0 {
			currentFlag = true
		}
		if currentFlag {
			row++
		} else {
			row--
		}
	}
	var result []byte
	result = make([]byte, 0, strlen)
	for i := 0; i < numRows; i++ {
		result = append(result, z[i]...)
	}
	return string(result)
}

/**
整理规律, 然后转化
24ms
 */
func convert(s string, numRows int) string {
	// 行数为1, 直接返回
	if numRows < 2 {
		return s
	}
	// 假设总行数大于等于3
	// 1. 第一行
	c := ""
	index := 0
	a := 0
	for index <= len(s)-1 {
		c += string(s[index])
		a++
		index = (2*numRows - 2) * a
	}
	// 2. 中间行 √  ALSIGYAHR  × ALLIIYAARR
	// 重置index
	for i := 2; i <= numRows-1; i++ {
		// 从第2行遍历到倒数 第二行
		// 重置 a , 在每一行的位置
		a = 0
		index = i - 1
		// 每一行都再添加一遍
		for index <= len(s)-1 {
			if a%2 == 0 {
				// a 为偶数
				c += string(s[index])
				a++
				index = (2*numRows-2)*(a+1)/2 - i + 1
			} else {
				// a 为奇数时
				c += string(s[index])
				a++
				index = (2*numRows-2)*a/2 + i - 1
			}
		}
	}
	// 3. 最后一行
	a = 0
	index = numRows - 1
	for index <= len(s)-1 {
		c += string(s[index])
		a++
		index = (2*numRows-2)*a + numRows - 1
	}
	return c
}

/**
numrows 总行数
row 当前行数
 */
func getCount(len, numRows, row int) int {
	//
	//count := 0
	//if row > numRows {
	//	return 0
	//}
	return 0
}

/**
暴力法
生成 numRows 个string
 */
func convert2(s string, numRows int) string {
	// 行数为1, 直接返回
	if numRows < 2 {
		return s
	}
	return ""
}
