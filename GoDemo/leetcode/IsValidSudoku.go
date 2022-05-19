package leetcode

func IsValidSudoku(board [][]byte) bool {
    rows := make([][]int, 10)
	columns := make([][]int, 10)
	boxs := make([][]int, 10)

	for i := 1; i < 10; i++ {
		rows[i] = make([]int, 10)
		columns[i] = make([]int, 10)
		boxs[i] = make([]int, 10)
	}

    for i := 0;i<9;i++ {
        for j := 0;j<9;j++ {
            if(board[i][j] == '.') {
                continue;
            }

            currentValue := board[i][j]-'0'
            if(rows[i][currentValue] == 1) {
                return false
            }
            if(columns[j][currentValue] == 1) {
                return false
            }
            if(boxs[j/3 + (i / 3) * 3][currentValue] == 1) {
                return false
            }

            //遍历byte数组把数字放进对应的数组中
            rows[i][currentValue] = 1
            columns[j][currentValue] = 1
            boxs[j/3 + (i / 3) * 3][currentValue] = 1
        }
    }
    return true
}