package main

import "fmt"

func main() {
	arr := []int{7, 3, 8, 2, 7, 9, 10, 1, 7, 54}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	//设定中轴值
	l, r := left, right
	pivot := arr[(left+right)/2]
	fmt.Println(pivot)
	//第一层计算，将值挪动到中轴值得两侧，
	for l < r { //直邀左索引小于右索引就一直进行
		//查到左边大于等于中轴值的位置
		for arr[l] < pivot { //找到第一个就行，只要比pivot小就一直右挪
			l++
		}
		//查到右边第一个小于等于中轴值得位置
		for arr[r] > pivot { //只要比中轴值大就一直左挪
			r--
		}
		//挪完之后，有可能l已经>=r了，说明左右两边已经分好了，这时候就没有必要再进行交换位置了。
		// 但是这里有等于的情况，如果等于的话，递归就无法停止。所以这种情况得单独处理下
		if l >= r {
			break
		}
		//找到了，交换位置
		arr[l], arr[r] = arr[r], arr[l]
		fmt.Printf("l:%v, r:%v,pivot:%v, arr:%v\n", l, r, pivot, arr)
		//如果交换完后，发现这个arr[l]==pivot值相等,r--，前移
		/*if arr[l] == arr[r] {
			r--
			l++
		}*/
		//如果交换完后，发现这个arr[l]==pivot值相等,r--，前移
		if arr[l] == pivot {
			r--
		}
		//如果交换完后，发现这个arr[r]==pivot值相等,l++，前移
		if arr[r] == pivot {
			l++
		}
	}
	//如果l==r，就都挪动一个,打破无限递归
	if l == r {
		l++
		r--
	}
	if left < r {
		quickSort(arr, left, r)
	}
	if right > l {
		quickSort(arr, l, right)
	}
}
