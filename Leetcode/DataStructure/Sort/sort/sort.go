package sort

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

//参考网址: https://zhuanlan.zhihu.com/p/40695917

// BubbleSort 冒泡排序
func BubbleSort() {
	array := util.GenerateRandomArray(4, 0, 5)
	len := len(array)
	exchange := false
	fmt.Print("begin: ")
	fmt.Println(array)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				exchange = true
				fmt.Println(array)
			}
		}
		if !exchange {
			fmt.Println("break")
			break
		}
	}
	fmt.Println(array)
}

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
