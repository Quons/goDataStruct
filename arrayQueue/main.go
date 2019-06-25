package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	//创建队列
	queue := CreateQueue(4)
	loop := true
	opt := ""
	for loop {
		fmt.Println("s(show):显示队列")
		fmt.Println("e(exit):退出程序")
		fmt.Println("a(add):添加数据到队列")
		fmt.Println("g(get):从队列中数据")
		fmt.Println("h(head):查看队列的头数据")
		_, err := fmt.Scanf("%s\n", &opt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch opt {
		case "s":
			err = queue.ShowQueue()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "a":
			fmt.Println("请输入数字")
			var i int
			_, err = fmt.Scanf("%v\n", &i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err = queue.AddQueue(i); err != nil {
				fmt.Println(err)
				continue
			}
		case "g":
			i, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(i)
		case "h":
			i, err := queue.headQueue()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(i)
		case "e":
			loop = false
		default:
			fmt.Println("invalid opt")
		}
	}
	fmt.Println("exit..")
}

//创建数组队列结构体
type ArrayQueue struct {
	MaxSize int
	Front   int
	Rear    int
	Arr     []int
}

//判断队列是否已满
func CreateQueue(size int) ArrayQueue {
	var queue ArrayQueue
	queue.MaxSize = size
	queue.Front = 0
	queue.Rear = 0
	queue.Arr = make([]int, size)
	return queue
}

//判断队列是否已满
func (a *ArrayQueue) IsFull() bool {
	//如果尾部等于容量-1
	return (a.Rear+1)%a.MaxSize == a.Front
}

//判断队列是否为空
func (a *ArrayQueue) IsEmpty() bool {
	return a.Rear == a.Front
}

//添加数据
func (a *ArrayQueue) AddQueue(d int) error {
	//先判断是否已满
	if a.IsFull() {
		return errors.New("queue is full")
	}
	a.Arr[a.Rear] = d
	a.Rear = (a.Rear + 1) % a.MaxSize
	return nil
}

//取数据
func (a *ArrayQueue) GetQueue() (int, error) {
	//先判断是否为空
	if a.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	tmp := a.Arr[a.Front]
	a.Front = (a.Front + 1) % a.MaxSize
	return tmp, nil
}

//遍历队列
func (a *ArrayQueue) ShowQueue() error {
	//先判断是否为空
	if a.IsEmpty() {
		return errors.New("empty queue")
	}
	for i := a.Front; i < a.Front+a.size(); i++ {
		fmt.Println(a.Arr[i%a.MaxSize])
	}
	return nil
}

//取头部数据
func (a *ArrayQueue) headQueue() (int, error) {
	//先判断是否为空
	if a.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return a.Arr[a.Front+1], nil
}

//取头部数据
func (a *ArrayQueue) size() int {
	//先判断是否为空
	return (a.Rear + a.MaxSize - a.Front) % a.MaxSize
}
func Test() {
var a =23
fmt.Println(a)

}