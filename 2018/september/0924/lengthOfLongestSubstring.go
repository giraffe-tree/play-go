package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*/
func main() {
	fmt.Println(lengthOfLongestSubstring3("bbbb"))
	fmt.Println(lengthOfLongestSubstring3("axasd"))
	fmt.Println(lengthOfLongestSubstring3("bbbb"))
}

/**
version 3.0
	用时 8 ms -> 4 ms
	数组版本
*/
func lengthOfLongestSubstring3(s string) int {
	var max = 0
	var end = 0
	array := [256]int{0}
	for start := 0; start < len(s); start++ {
		// uint8
		char := s[start]
		value := array[char]
		array[char] = start + 1
		if value == 0 || value-1 < end {
			// 如果不存在, 加入 map 中

			if start-end+1 > max {
				max = start - end + 1
			}
			continue
		}
		//如果存在, 拿到 value 即重复的 序列号
		end = value
	}
	return max
}

/**
 version2.0
	使用 map , 不用 delete 元素
	用时 12ms
*/
func lengthOfLongestSubstring2(s string) int {
	var length = len(s)
	var maxCount = 0
	var end = 0
	hashMap := make(map[uint8]int)
	for start := 0; start < length; start++ {
		// uint8
		char := s[start]

		value, ok := hashMap[char]
		if !ok || value < end {
			// 如果不存在, 加入 map 中
			hashMap[char] = start
			if start-end+1 > maxCount {
				maxCount = start - end + 1
			}
			continue
		}
		//如果存在, 拿到 value 即重复的 序列号
		if value >= end {
			end = value + 1
		}

		hashMap[char] = start
	}
	//fmt.Printf("start:%d end:%d\n", start, end)
	return maxCount
}

/**
简单思路, 设定一个 start , 一个end 指针, 指向最长的子串
最初版本
*/
func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	start, end := 0, 0
	var maxCount = 0
	hashMap := make(map[uint8]int)
	for index := 0; index < length; index++ {
		// uint8
		char := s[start]
		start++
		value, ok := hashMap[char]
		if !ok {
			// 如果不存在, 加入 map 中
			hashMap[char] = index
			// 使用 start -end 消耗资源更少
			var size = start - end
			//var size = len(hashMap)
			if size > maxCount {
				maxCount = size
			}
			continue
		}
		//如果存在, 拿到 value 即重复的 序列号
		for j := end; j < value+1; j++ {
			// 如何去除 map 中的元素
			delete(hashMap, s[j])
		}
		end = value + 1
		hashMap[char] = index
	}
	fmt.Printf("start:%d end:%d\n", start, end)
	return maxCount
}

/**

 */
func lengthOfLongestSubstringCopy2(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		char := s[i]
		loc := location[char]
		fmt.Printf("s[%d]: %d location[%d]:%d \n ", i, char, char, location[char])

		if loc >= left {
			left = loc + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[char] = i
		fmt.Print()
	}

	return maxLen
}

/**
copy from leetcode
*/
func lengthOfLongestSubstringCopy(s string) int {
	m := [256]int{0}
	var res = 0
	var left = 0
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 || m[s[i]] < left {
			if res < i-left+1 {
				res = i - left + 1
			}
		} else {
			left = m[s[i]]
		}
		m[s[i]] = i + 1
	}
	return res
}

/**
不能用
*/
func lengthOfLongestSubstringOld(s string) int {
	length := len(s)
	// 数组声明 https://stackoverflow.com/questions/38362631/go-error-non-constant-array-bound
	// 原代码: var array [length]int
	// You can't instantiate an array like that with a value calculated at runtime.
	// Instead use make to initialize a slice with the desired length. It would look like this;
	array := make([]int, length)

	hashMap := make(map[uint8]int)
	for i := 0; i < length; i++ {
		char := s[i]
		index, ok := hashMap[char]
		if !ok {
			// 不存在, 则放进 map 中
			hashMap[char] = i
		} else {
			// 存在, 则放进 array 中
			array[i] = 1
			array[index] = 1
		}
	}
	var maxLen = 0
	var current = 0
	for j := 0; j < length; j++ {
		if array[j] == 0 {
			current++
		} else {
			current = 0
		}
		if current >= maxLen {
			maxLen = current
		}
	}
	return maxLen
}
