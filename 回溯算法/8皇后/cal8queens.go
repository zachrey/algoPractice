package al39cal8queens

import (
	"fmt"
)

var rst = make([]int, 8)

func cal8queens(row int) {
	if row >= 8 {
		printRst(rst)
		return
	}

	for i := 0; i < 8; i++ {
		if isOk(row, i) {
			rst[row] = i
			cal8queens(row + 1)
		}
	}
}

func isOk(row, column int) bool {
	leftUp, rightUp := column-1, column+1 // 左上角，右上角
	for i := row - 1; i >= 0; i-- {
		if rst[i] == column {
			return false // 正上方有放棋子
		}
		if leftUp >= 0 && rst[i] == leftUp {
			return false // 左斜方向有棋子
		}
		if rightUp < 8 && rst[i] == rightUp {
			return false // 右斜方向有棋子
		}
		leftUp--
		rightUp++
	}
	return true
}

func printRst(rst []int) {
	for row := 0; row < 8; row++ {
		for i := 0; i < 8; i++ {
			if rst[row] == i {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
