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




