package leetcode

import "fmt"

func MinPathSum(grid [][]int) int {
    width, high := len(grid[0]), len(grid);
     
    //如果长高为0，直接返回0
    if(width == 0 || high == 0) {
        return 0;
    }
    //初始化汇总网格
    for i:=1;i<high;i++{
        grid[i][0] += grid[i - 1][0];
    }
    for i:=1;i<width;i++ {
        grid[0][i] += grid[0][i - 1];
    }
    for i:=1;i <high;i++ {
        for j:=1;j <width;j++ {
			fmt.Println(grid[i][j]);
            grid[i][j] += min(grid[i - 1][j], grid[i][j - 1]);
        }
    }
    return grid[high - 1][width - 1];            
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}