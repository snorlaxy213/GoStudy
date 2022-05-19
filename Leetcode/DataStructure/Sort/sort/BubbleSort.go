package sort

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

//参考网址: https://zhuanlan.zhihu.com/p/40695917

// BubbleSort 冒泡排序
func BubbleSort() {
	array := util.GenerateRandomArray(4, 0, 5)
	length := len(array)
	fmt.Print("begin: ")
	fmt.Println(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Print("end: ")
	fmt.Println(array)
}

//优化点：定一个变量exchange记录当某一一轮是否发生交换行为，若为发生则判定已经排序成功，跳出循环即可。
func BubbleSortV2() {
	array := util.GenerateRandomArray(10, 0, 100)
	len := len(array)
	fmt.Print("begin: ")
	fmt.Println(array)
	for i := 0; i < len - 1; i++ {
		exchange := false
		for j := 0; j < len -1 - i; j++ {
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
	fmt.Print("end: ")
	fmt.Println(array)
}

//优化点：记录一轮交换后的最终索引，通过观察发现这个索引后的数字都是有序的，那么以后就不用比较这个索引后面的数字，即内层循环的边界是前面所说的最终索引。
//当这个最终索引为0的时候（即前面没有数字的时候）排序就结束了。
func BubbleSortV3() {
	array := util.GenerateRandomArray(10, 0, 100)
	n := len(array) - 1
	fmt.Print("begin: ")
	fmt.Println(array)
	for(true) {
		// 记录最后一次交换索引位置
        last := 0;
		for j := 0; j < n; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				last = j;
				fmt.Println(array)
			}
		}
		n = last
		if n == 0 {
			break
		}
	}
	fmt.Print("end: ")
	fmt.Println(array)
}




