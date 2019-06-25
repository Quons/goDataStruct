package main

import "fmt"

func main() {
	//	创建二维数组
	//给二维数组赋值
	//遍历二维数组
	sparseArray1 := [11][11]int{}
	sparseArray1[1][2] = 1
	sparseArray1[2][4] = 2
	for _, row := range sparseArray1 {
		for _, v := range row {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
	//查询有多少个有效值
	sum := 0
	for _, row := range sparseArray1 {
		for _, v := range row {
			if v != 0 {
				sum++
			}
		}
	}
	//创建相应的稀疏数组
	chessArray := make([][3]int, sum+1)
	chessArray[0][0] = len(sparseArray1)
	chessArray[0][1] = len(sparseArray1)
	chessArray[0][2] = 2
	count := 0
	for i, row := range sparseArray1 {
		for j, v := range row {
			if v != 0 {
				count++
				chessArray[count][0] = i
				chessArray[count][1] = j
				chessArray[count][2] = v
			}
		}
	}
	//输出
	for _, row := range chessArray {
		fmt.Printf("%v\t%v\t%v\n", row[0], row[1], row[2])
	}
	//恢复成原始数组
	//创建数组
	originArray := make([][]int, chessArray[0][0])
	for i := range originArray {
		originArray[i] = make([]int, chessArray[0][1])
	}
	for i, row := range chessArray {
		if i == 0 {
			continue
		}
		originArray[row[0]][row[1]] = row[2]
	}
	fmt.Println("origin array~~")
	for _, row := range originArray {
		for _, v := range row {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}
