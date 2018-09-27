package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/longest-palindromic-substring/description/
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
func main() {
	//var s = "ddabadd"
	//fmt.Println(check(s, 3, 1))
	//
	//fmt.Println(checkOu(s, 0, 1))
	//s := " "
	//fmt.Println(s[0])
	palindrome2 := longestPalindromeCopy("abbca")
	fmt.Println(palindrome2)

	//s := check2("abcba", 2, 3)
	//fmt.Println(s)
}

/**
copy from leetcode 预期 0ms
这种方法区分开了 aba 和 abba , 虽然跟我的思路目的是一样的, 但简洁了不止一点
*/
func longestPalindromeCopy(s string) string {
	if len(s) < 2 {
		return s
	}
	begin, maxLen := 0, 1

	for i := 0; i < len(s); {
		// 当剩余的可遍历长度不够时
		if len(s)-i <= maxLen/2 {
			break
		}

		b, e := i, i

		// 有多少重复的
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}

		newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}

/**
思路2:
将原字符串, 假设每两个字符之间插入一个相同的字符,这样回文字符串必定为奇数
已通过: 28ms
*/
func longestPalindrome2(s string) string {
	oldLen := len(s)
	if oldLen <= 1 {
		return s
	}
	var x = ""
	for j := 0; j < oldLen; j++ {
		x += string(s[j])
		if j < oldLen-1 {
			x += " "
		}
	}
	maxLen := 1
	maxString := string(s[0])
	for i := 0; i < len(x); i++ {
		b := check2(x, i, maxLen+1)
		for b != "" {
			if maxLen == len(b) {
				b = check2(x, i, maxLen+4)
				if b != "" {
					maxString = b
				}
				continue
			}
			maxLen = len(b)
			maxString = b
			b = check2(x, i, maxLen+2)
		}
	}
	return strings.Replace(maxString, " ", "", -1)
}

func check2(s string, index, maxLen int) string {
	// 处理加了空格的字符串
	count := maxLen / 2
	for i := count; i >= 1; i-- {
		if index-i < 0 {
			return ""
		}
		if index+i > len(s)-1 {
			return ""
		}
		if i == count && s[index-i] == 32 && s[index+i] == 32 {
			return s[index-count+1 : index+count]
		}
		if s[index-i] != s[index+i] {
			return ""
		}
	}

	return s[index-count : index+count+1]
}

/**
简单思路:
使用一个 maxLen 记录最长回文的长度, 遍历原始string, 时间复杂度为 O(n) ?
1. 回文数为奇数
2. 回文数为偶数

太复杂...先尝试其他方案..
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	var maxCount = 1
	var maxLen = 1
	for i := 1; i < length-1; i++ {
		// 检查奇数回文
		maxCount = getCount(maxLen)
		for check(s, i, maxCount+1) {
			maxLen = 2*maxCount - 1
			maxCount = getCount(maxLen)
		}

		for checkOu(s, i, maxCount) {
			maxLen = 2 * maxCount
			maxCount = getCount(maxLen)
		}
	}
	return ""
}

func getCount(maxLen int) int {
	return (maxLen + 1) / 2
}

/**
检查多少位之内是否为回文
奇数
*/
func check(s string, index, count int) bool {
	i2 := index - count
	if i2 < 0 {
		return false
	}
	i3 := index + count
	if i3 > len(s)-1 {
		return false
	}
	for i := 1; i <= count; i++ {
		if s[index-i] != s[index+i] {
			return false
		}
	}
	return true
}

func checkOu(s string, index, count int) bool {

	i2 := index - count + 1
	if i2 < 0 {
		return false
	}
	i3 := index + count
	if i3 > len(s)-1 {
		return false
	}
	for i := 1; i <= count; i++ {
		if s[index-i+1] != s[index+i] {
			return false
		}
	}
	return true
}
