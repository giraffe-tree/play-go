package main

import "fmt"

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
你可以假设 nums1 和 nums2 不同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]
中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
中位数是 (2 + 3)/2 = 2.5
*/
func main() {
	ints1 := []int{1, 2}
	ints2 := []int{1, 2}

	findMedianSortedArrays(ints1, ints2)
}

/**
暴力法
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	length1 := len(nums1)
	length2 := len(nums2)
	length := length1 + length2
	fmt.Println(length)

	return 0.0
}

/**
copy
从 leetcode 上的 java 实现转义成 go 实现

*/
func findMedianSortedArraysCopy(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		temp := nums1
		nums1 = nums2
		nums2 = temp
		tmp := m
		m = n
		n = tmp
	}
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			// i is too small
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			// i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums2[j-1]
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else if nums2[j] > nums1[i] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return (float64(maxLeft) + float64(minRight)) / 2
		}
	}
	return 0.0
}

/**
java 实现

```java
public double findMedianSortedArrays(int[] A, int[] B) {
        int m = A.length;
        int n = B.length;
        if (m > n) { // to ensure m<=n
            int[] temp = A; A = B; B = temp;
            int tmp = m; m = n; n = tmp;
        }
        int iMin = 0, iMax = m, halfLen = (m + n + 1) / 2;
        while (iMin <= iMax) {
            int i = (iMin + iMax) / 2;
            int j = halfLen - i;
            if (i < iMax && B[j-1] > A[i]){
                iMin = i + 1; // i is too small
            }
            else if (i > iMin && A[i-1] > B[j]) {
                iMax = i - 1; // i is too big
            }
            else { // i is perfect
                int maxLeft = 0;
                if (i == 0) { maxLeft = B[j-1]; }
                else if (j == 0) { maxLeft = A[i-1]; }
                else { maxLeft = Math.max(A[i-1], B[j-1]); }
                if ( (m + n) % 2 == 1 ) { return maxLeft; }

                int minRight = 0;
                if (i == m) { minRight = B[j]; }
                else if (j == n) { minRight = A[i]; }
                else { minRight = Math.min(B[j], A[i]); }

                return (maxLeft + minRight) / 2.0;
            }
        }
        return 0.0;
    }
```
*/
func arrayRange(nums []int) {
	// 遍历数组
	for index, value := range nums {
		fmt.Println(index, value)
	}

}
