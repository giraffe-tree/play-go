package main

import "fmt"

/**
https://leetcode-cn.com/problems/container-with-most-water/description/

11. 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

示例:
输入: [1,8,6,2,5,4,8,3,7]
输出: 49
 */
func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea3(array))

}

/**
参照 leetcode 阅读理解中的做法 时间复杂度 O(n)

 */
func maxArea3(height []int) int {
	length := len(height)
	l := 0
	r := length - 1
	maxArea := 0
	var current int
	for i := 0; i < length && l < r; i++ {
		if height[l] > height[r] {
			current = height[r]
			area := current * (r-l)
			if area > maxArea {
				maxArea = area
			}
			r--
		} else {
			current = height[l]
			area := current * (r-l)
			if area > maxArea {
				maxArea = area
			}
			l++
		}
	}
	return maxArea
}

/**
优化失败了...
 */
func maxArea2(height []int) int {
	length := len(height)
	maxArea := 0
	for i := 0; i < length; i++ {
		i2 := height[i]
		maxHeight := 0
		for j := i + 1; j < length; j++ {
			current := height[j]
			if current > i2 {
				current = i2
			}
			if current > maxHeight {
				maxHeight = current
			} else {
				continue
			}
			area := maxHeight * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

/**
暴力计算
时间复杂度 O(n^2)
 */
func maxArea(height []int) int {
	length := len(height)
	maxArea := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			i2 := height[i]
			if height[j] < i2 {
				i2 = height[j]
			}
			area := i2 * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
