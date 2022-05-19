package sort

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

// SelectionSort 选择排序
func SelectionSort() {
	array := util.GenerateRandomArray(10, 0, 10)
	len := len(array)

	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if array[minIndex] > array[j] {
				array[minIndex], array[j] = array[j], array[minIndex]
			}
		}
	}
	fmt.Println(array)
}