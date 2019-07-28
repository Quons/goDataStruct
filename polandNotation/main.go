package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//逆波兰表达式 3 4 - 5 * 6 -
	//先将运算符拆分，得到一个slice
	expression := "3 4 + 5 * 6 -"
	list := strings.Split(expression, " ")
	//将slice传入到一个方法中
	res, err := Calculate(list)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func Calculate(list []string) (int64, error) {
	//创建栈
	stack := CreateStack(10)
	for _, value := range list {
		//判断是否是数字，如果是数字，就直接入栈，
		match, err := regexp.MatchString("\\d+", value)
		if err != nil {
			return 0, err
		}
		if match {
			//入栈
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return 0, err
			}
			stack.Push(v)
		} else {
			//运算符，先出栈两个数，然后计算，再入栈
			num2, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			num1, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			var res int64
			switch value {
			case "*":
				res = num1 * num2
			case "/":
				res = num1 / num2
			case "+":
				res = num1 + num2
			case "-":
				res = num1 - num2
			default:
				return 0, errors.New("invalid operator")
			}
			stack.Push(res)
		}
	}
	res, err := stack.Pop()
	if err != nil {
		return 0, err
	}
	return res, nil
}

type ArrayStack struct {
	MaxSize int
	Top     int
	Stack   []int64
}

func CreateStack(maxSize int) (a ArrayStack) {
	a.MaxSize = maxSize
	a.Top = -1
	a.Stack = make([]int64, maxSize, maxSize)
	return
}

func (a *ArrayStack) IsFull() bool {
	return a.Top == a.MaxSize-1
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Top == -1
}

func (a *ArrayStack) Push(value int64) {
	if a.IsFull() {
		fmt.Println("stack is full")
		return
	}
	a.Top++
	a.Stack[a.Top] = value
}

func (a *ArrayStack) Pop() (value int64, err error) {
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
