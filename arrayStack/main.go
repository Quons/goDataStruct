package main

import (
	"errors"
	"fmt"
)

func main() {
	arrayStack := CreateStack(4)
	for {
		fmt.Println("show:表示显示栈")
		fmt.Println("exit:退出程序")
		fmt.Println("push:表示添加数据到栈")
		fmt.Println("pop:出栈")
		fmt.Println("请输入你的选择")
		var opt string
		_, err := fmt.Scanf("%v", &opt)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch opt {
		case "show":
			arrayStack.list()
		case "exit":
			return
		case "push":
			fmt.Println("请输入一个数字")
			var i int
			_, err := fmt.Scanf("%v", &i)
			if err != nil {
				fmt.Println(err)
				return
			}
			arrayStack.Push(i)
		case "pop":
			res, err := arrayStack.Pop()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("出栈的数为:%v\n", res)
		default:
			fmt.Println("invalid option")
		}
	}

}

type ArrayStack struct {
	MaxSize int
	Top     int
	Stack   []int
}

func CreateStack(maxSize int) (a ArrayStack) {
	a.MaxSize = maxSize
	a.Top = -1
	a.Stack = make([]int, maxSize, maxSize)
	return
}

func (a *ArrayStack) IsFull() bool {
	return a.Top == a.MaxSize-1
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Top == -1
}

func (a *ArrayStack) Push(value int) {
	if a.IsFull() {
		fmt.Println("stack is full")
		return
	}
	a.Top++
	a.Stack[a.Top] = value
}

func (a *ArrayStack) Pop() (value int, err error) {
	if a.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	result := a.Stack[a.Top]
	a.Top--
	return result, nil
}

func (a *ArrayStack) list() {
	if a.IsEmpty() {
		return
	}
	for i := a.Top; i >= 0; i-- {
		fmt.Printf("stack[%v]:%v\n", i, a.Stack[i])
	}
}
