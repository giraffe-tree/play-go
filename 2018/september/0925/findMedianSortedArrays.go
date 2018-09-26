package main

func main() {

}

/**
暴力法
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	length1 := len(nums1)
	length2 := len(nums2)
	length := length1 + length2
	count := length2
	flag := false
	if length1 > length2 {
		count = length1
		flag = true
	}
	for i := 0; i < count; i++ {
		num1 := nums1[i]

	}

}
