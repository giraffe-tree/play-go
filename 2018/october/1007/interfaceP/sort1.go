package interfaceP

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5484, 7586}
	//conversion to type IntArray from package sort
	slices := sort.IntSlice{1, 4, 2, 3, 10, 8, 5, 6, 9}
	sort.Ints(data)
	sort.Sort(slices)
	fmt.Println(data)
	fmt.Println(slices)
}
