package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	expression := "70+2*6-3"
	//创建两个堆栈
	numStack := CreateStack(10)
	operStack := CreateStack(10)
	//定义相关变量
	num1 := 0
	num2 := 0
	res := 0
	oper := 0
	keepNum := ""
	index := 0
	var err error
	for {
		//获取到index所在的字符
		ch := int(expression[index])
		//判断当前是数字还是运算符
		if IsOper(ch) {
			//运算符
			if operStack.IsEmpty() {
				//为空，直接入栈
				operStack.Push(ch)
			} else {
				//不为空，判断当前运算符和栈顶运算符的优先级，如果小于等于，则先计算，再入栈，
				if Priority(ch) <= Priority(operStack.Peep()) {
					//出栈两个数，计算，入栈，
					num1, err = numStack.Pop()
					if err != nil {
						fmt.Println(err)
						return
					}
					num2, err = numStack.Pop()
					if err != nil {
						fmt.Println(err)
						return
					}
					oper, err = operStack.Pop()
					if err != nil {
						fmt.Println(err)
						return
					}
					res = Cal(num1, num2, oper)
					operStack.Push(ch)
					numStack.Push(res)
				} else {
					//大于之前的运算符，直接入栈
					operStack.Push(ch)
				}
			}
		} else {
			keepNum += string(ch)
			//判断是不是最后一个字符
			if index == len(expression)-1 {
				n, err := strconv.ParseInt(keepNum, 10, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				numStack.Push(int(n))
			} else {
				if IsOper(int(expression[index+1])) {
					n, err := strconv.ParseInt(keepNum, 10, 64)
					if err != nil {
						fmt.Println(err)
						return
					}
					numStack.Push(int(n))
					keepNum = ""
				}
			}
		}
		index++
		if index >= len(expression) {
			break
		}
	}
	//将最后栈中的元素进行计算
	for {
		if operStack.IsEmpty() {
			break
		}
		//出栈两个数，计算，入栈，
		num1, err = numStack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err = numStack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		oper, err = operStack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		res = Cal(num1, num2, oper)
		numStack.Push(res)
	}
	res, err = numStack.Pop()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("expression:%s,result:%v\n", expression, res)

}

type CalculateStack struct {
	MaxSize int
	Top     int
	Stack   []int
}

func CreateStack(maxSize int) (a CalculateStack) {
	a.MaxSize = maxSize
	a.Top = -1
	a.Stack = make([]int, maxSize, maxSize)
	return
}

func (a *CalculateStack) IsFull() bool {
	return a.Top == a.MaxSize-1
}

func (a *CalculateStack) IsEmpty() bool {
	return a.Top == -1
}

func (a *CalculateStack) Push(value int) {
	if a.IsFull() {
		fmt.Println("stack is full")
		return
	}
	a.Top++
	a.Stack[a.Top] = value
}

func (a *CalculateStack) Pop() (value int, err error) {
	if a.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	result := a.Stack[a.Top]
	a.Top--
	return result, nil
}

//获取栈顶元素
func (a *CalculateStack) Peep() int {
	return a.Stack[a.Top]
}

//返回运算符的优先级
func Priority(oper int) int {
	if oper == '*' || oper == '/' {
		return 1
	} else if oper == '+' || oper == '-' {
		return 0
	} else {
		return -1
	}
}

//判断是否是运算符
func IsOper(oper int) bool {
	return oper == '*' || oper == '/' || oper == '+' || oper == '-'
}

//计算两个值
func Cal(num1, num2, val int) int {
	if val == '*' {
		return num1 * num2
	} else if val == '/' {
		return num2 / num1
	} else if val == '+' {
		return num1 + num2
	} else if val == '-' {
		return num2 - num1
	} else {
		return 0
	}
}
