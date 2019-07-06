package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	csl := CircleLinkedList{}
	err := csl.AddBoy(125)
	if err != nil {
		fmt.Println(err)
		return
	}
	//csl.ShowBoy()
	csl.CountBoy(10, 20, 125)
}

type CircleLinkedList struct {
	First *Boy
}

func (c *CircleLinkedList) AddBoy(num int) error {
	//检查num的有效性
	if num < 1 {
		return errors.New("invalid num")
	}
	//先处理第一个节点
	c.First = &Boy{1, nil}
	c.First.Next = c.First
	currBoy := c.First
	//从第2个节点开始遍历
	for i := 2; i <= num; i++ {
		//创建临时的节点
		b := &Boy{i, nil}
		//之后的节点
		b.Next = c.First
		currBoy.Next = b
		currBoy = b
	}
	return nil
}

func (c *CircleLinkedList) ShowBoy() {
	var currBoy = c.First
	for {
		fmt.Printf("boy:%v\n", currBoy.No)
		currBoy = currBoy.Next
		if currBoy.No == c.First.No {
			break
		}
	}
}

func (c *CircleLinkedList) CountBoy(startNum, countNum, totalNum int) {
	//校验参数
	if countNum < 0 || startNum < 0 || startNum > totalNum {
		fmt.Println("invalid param")
		return
	}
	//将helper和starter移动到指定的位置
	starter := c.First
	for i := 0; i < startNum-1; i++ {
		starter = starter.Next
	}
	//获取helper
	var helper *Boy
	helper = starter
	for {
		if helper.Next.No == starter.No {
			break
		}
		helper = helper.Next
	}
	//遍历
	for {
		if helper.No == starter.No {
			fmt.Printf("No:%v\n", helper.No)
			break
		}
		//数数
		for i := 0; i < countNum-1; i++ {
			helper = starter
			starter = starter.Next
		}
		//出队
		fmt.Printf("No:%v\n", starter.No)
		helper.Next = starter.Next
		starter = helper.Next
	}
}

type Boy struct {
	No   int
	Next *Boy
}
