package main

import (
	"fmt"

	"code.vino/GoStudy/Leetcode/DataStructure/Sort/sort"
	"code.vino/GoStudy/Leetcode/DataStructure/Sort/util"
)

func main() {

	fmt.Println(util.GenerateRandomArray(10, 0, 100))

	fmt.Println(util.GenerateNearlyOrderedArray(10, 2))

	s := []rune{1, 2, 3}
	fmt.Sprintln(util.GenerateRandomString(10, s))

	fmt.Println(util.IsSorted(util.GenerateRandomArray(10, 0, 100)))

	//Bubble Sort
	sort.BubbleSort()
}
