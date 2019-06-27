package main

import (
	"errors"
	"fmt"
)

func main() {
	var heroList HeroNodeList
	h1 := HeroNode{1, "林冲", "豹子头", nil}
	h3 := HeroNode{3, "张璐", "玉麒麟", nil}
	h5 := HeroNode{5, "武松", "老虎克星", nil}
	if err := heroList.Add(h1); err != nil {
		fmt.Print(err)
	}
	if err := heroList.Add(h3); err != nil {
		fmt.Print(err)
	}
	if err := heroList.Add(h5); err != nil {
		fmt.Print(err)
	}

	var heroListNew HeroNodeList
	h2 := HeroNode{2, "鲁智深", "花和尚", nil}
	h4 := HeroNode{4, "武松", "老虎克星", nil}
	h6 := HeroNode{6, "武松2", "老虎克星", nil}

	if err := heroListNew.Add(h2); err != nil {
		fmt.Print(err)
	}
	if err := heroListNew.AddByOrder(h4); err != nil {
		fmt.Print(err)
	}
	if err := heroListNew.AddByOrder(h6); err != nil {
		fmt.Print(err)
	}
	//h6.Name = "刘恒越"
	//heroList.Update(h6)
	//heroList.List()
	//fmt.Println(GetLength(*heroList.Head))
	//node, err := GetLastIndexNode(*heroList.Head, 6)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%v\n", node.String())
	//ReverseLinkedList(heroList.Head)
	//heroList.List()
	//fmt.Println("反转打印。。")
	//ReversePrint(heroList.Head)
	resultList := CombineList(*heroList.Head, *heroListNew.Head)
	resultList.List()
}

//获取链表的有效长度，带头结点的链表不包含头结点
func GetLength(head HeroNode) (int, error) {
	if head.next == nil {
		return 0, errors.New("empty list")
	}
	var length int
	tmp := head.next
	for tmp != nil {
		length++
		tmp = tmp.next
	}
	return length, nil

}

//获取单链表的倒数第k个结点
func GetLastIndexNode(head HeroNode, index int) (HeroNode, error) {
	//判断是否为空
	var hero HeroNode
	if head.next == nil {
		return hero, errors.New("empty list")
	}
	//获取有效长度
	length, err := GetLength(head)
	if err != nil {
		return hero, err
	}
	//校验index 是否有效
	if index <= 0 || index > length {
		return hero, errors.New("invalid index")
	}
	tmp := head.next
	for i := 0; i < length-index; i++ {
		tmp = tmp.next
	}
	return *tmp, nil
}

//反转链表
func ReverseLinkedList(head *HeroNode) {
	//判断是否为空
	if head.next == nil || head.next.next == nil {
		return
	}
	//定义临时的reverseHead
	reverseHead := HeroNode{0, "", "", nil}
	//定义当前的节点
	cur := head.next
	var next *HeroNode
	for cur != nil {
		//将临时链表的下一个节点拼到当前节点的后面，并将当前节点拼到临时链表的头结点后面
		next = cur.next
		cur.next = reverseHead.next
		reverseHead.next = cur
		cur = next
	}
	head.next = reverseHead.next
}

//反转打印
func ReversePrint(head *HeroNode) {
	//判断是否为空
	if head.next == nil {
		return
	}
	//先声明一个slice
	stack := make([]*HeroNode, 0, 100)
	cur := head.next
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}

//将两个有序链表合并，最后返回新的有序的链表
func CombineList(first, second HeroNode) (resultList HeroNodeList) {
	//创建head
	resultList.Head = &HeroNode{0, "", "", nil}
	resultCur := resultList.Head
	firstCur := first.next
	secondCur := second.next
	var tmpNode *HeroNode
	for {
		if firstCur == nil {
			resultCur.next = secondCur
			break
		}
		if secondCur == nil {
			resultCur.next = firstCur
			break
		}
		//判断哪个大
		if firstCur.No > secondCur.No {
			tmpNode = secondCur
			secondCur = secondCur.next
		} else {
			tmpNode = firstCur
			firstCur = firstCur.next
		}
		resultCur.next = tmpNode
		resultCur = tmpNode
	}
	return
}

//创建链表结构体
type HeroNodeList struct {
	Head *HeroNode
}

func (heroList *HeroNodeList) Add(hero HeroNode) error {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &HeroNode{0, "", "", nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &hero
	return nil
}

func (heroList *HeroNodeList) Update(hero HeroNode) {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &HeroNode{0, "", "", nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	flag := false
	for {
		if tmp == nil {
			break
		}
		if tmp.No == hero.No {
			flag = true
			break
		}
		tmp = tmp.next
	}
	if flag {
		tmp.Name = hero.Name
		tmp.nickname = hero.nickname
	} else {
		fmt.Println("get no hero")
	}
}
func (heroList *HeroNodeList) Delete(hero HeroNode) {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &HeroNode{0, "", "", nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	flag := false
	for {
		if tmp.next == nil {
			break
		}
		if tmp.No == hero.No {
			flag = true
			break
		}
		tmp = tmp.next
	}
	if flag {
		tmp.next = tmp.next.next
	} else {
		fmt.Println("get no hero")
	}
}

func (heroList *HeroNodeList) AddByOrder(hero HeroNode) error {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &HeroNode{0, "", "", nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	flag := true
	for {
		if tmp.next == nil {
			break
		}
		if tmp.next.No == hero.No {
			flag = false
			break
		}
		if tmp.next.No > hero.No {
			break
		}
		tmp = tmp.next
	}
	if flag {
		//添加
		hero.next = tmp.next
		tmp.next = &hero
	} else {
		return errors.New("标号存在，不能添加")
	}
	return nil
}

func (heroList *HeroNodeList) List() {
	//判断链表是否为空
	if heroList.Head == nil {
		fmt.Println("empty list")
		return
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	for tmp.next != nil {
		tmp = tmp.next
		fmt.Printf("%+v\n", tmp)
	}
}

//创建英雄struct
type HeroNode struct {
	No       int
	Name     string
	nickname string
	next     *HeroNode
}

func (hero *HeroNode) String() string {
	return fmt.Sprintf("no:%v,name:%s,nickname:%s", hero.No, hero.Name, hero.nickname)
}
