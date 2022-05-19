package sort

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

// InsertionSort 插入排序
func InsertionSort() {
	array := util.GenerateRandomArray(10, 1, 50)
	// array := []int{4, 1, 2}
	len := len(array)

	for i := 1; i < len; i++ {
		currentValue := array[i]
		position := i
		for position > 0 && array[position-1] > currentValue {
			array[position] = array[position-1]
			position = position - 1
		}
		array[position] = currentValue
	}

	fmt.Println(array)
}