package sort

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

// BubbleSort sad
func BubbleSort() {
	array := util.GenerateRandomArray(10, 0, 5)
	len := len(array)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
