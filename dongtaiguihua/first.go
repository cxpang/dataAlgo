package main

import (
	"fmt"
)

func main() {
	mat:=[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(len(mat))
	count:=CountPath(mat)
	fmt.Println(count)
}
func CountPath(a [][]int) int  {
	row := len(a)
	col := len(a[1])
	fmt.Println(row,col)
	var opt  [8][8]int

	for i:=row-1;i>=0 ; i-- {
		for j:=col-1;j>=0 ;j--  {
			if a[i][j] == 1 {
				opt[i][j] = 0
			}else if i==row-1 && j==col-1  {
				opt[i][j] = 0
			}else if i==row-1 || j==col-1 {
				opt[i][j] = 1

			}else {
				opt[i][j] = opt[i+1][j] + opt[i][j+1]
			}
		}
	}
	return opt[0][0]
}
