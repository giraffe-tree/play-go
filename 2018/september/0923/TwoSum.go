package main

import "fmt"

/**
https://leetcode-cn.com/problems/two-sum/description/
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

 */
func main() {
	// 数组赋值
	var nums = []int{3,2,4}
	var target = 6
	var x = twoSum2(nums, target)
	fmt.Println(x)
	var m = make(map[int]int, 3)
	for i := 0; i < 3; i++ {
		m[nums[i]] = i
	}
	var a,b =m[3]
	fmt.Println(a,b)

}

func twoSum(nums []int, target int) []int {
	var count = len(nums)

	// for 循环
	for i := 0; i < count; i++ {
		for j := i+1; j < count; j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}

	}
	return nil
}


func twoSum2(nums []int, target int) []int {
	l := len(nums)
	var m = make(map[int]int, l)
	for i := 0; i < l; i++ {
		m[nums[i]] = i
	}
	for i := 0; i < l; i++ {
		ni := nums[i]
		t := target - ni

		// 在第四种情况中，if 可以包含一个初始化语句（如：给一个变量赋值）。
		// 这种写法具有固定的格式（在初始化语句后方必须加上分号)
		if v, ok := m[t]; ok {
			if v != i {
				return []int{i, v}
			}
		}
	}

	return []int{}
}