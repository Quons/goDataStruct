package main

import (
	"errors"
	"fmt"
)

func main() {
	var heroList DoubleHeroNodeList
	h1 := DoubleHeroNode{1, "林冲", "豹子头", nil, nil}
	h3 := DoubleHeroNode{3, "张璐", "玉麒麟", nil, nil}
	h5 := DoubleHeroNode{5, "武松", "老虎克星", nil, nil}

	if err := heroList.Add(h1); err != nil {
		fmt.Print(err)
	}

	if err := heroList.Add(h3); err != nil {
		fmt.Print(err)
	}

	if err := heroList.Add(h5); err != nil {
		fmt.Print(err)
	}

	var heroListNew DoubleHeroNodeList
	h2 := DoubleHeroNode{2, "鲁智深", "花和尚", nil, nil}
	h4 := DoubleHeroNode{4, "武松", "老虎克星", nil, nil}
	h6 := DoubleHeroNode{6, "武松2", "老虎克星", nil, nil}

	/*if err := heroListNew.Add(h2); err != nil {
		fmt.Print(err)
	}*/
	if err := heroListNew.AddByOrder(h4); err != nil {
		fmt.Print(err)
	}
	if err := heroListNew.AddByOrder(h6); err != nil {
		fmt.Print(err)
	}
	heroListNew.Delete(h2)
	heroListNew.List()
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
}

//创建链表结构体
type DoubleHeroNodeList struct {
	Head *DoubleHeroNode
}

func (heroList *DoubleHeroNodeList) Add(hero DoubleHeroNode) error {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &DoubleHeroNode{0, "", "", nil, nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &hero
	hero.Pre = tmp
	return nil
}

func (heroList *DoubleHeroNodeList) Update(hero DoubleHeroNode) {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &DoubleHeroNode{0, "", "", nil, nil}
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
		tmp = tmp.Next
	}
	if tmp != nil && flag {
		tmp.Name = hero.Name
		tmp.Nickname = hero.Nickname
	} else {
		fmt.Println("get no hero")
	}
}
func (heroList *DoubleHeroNodeList) Delete(hero DoubleHeroNode) {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &DoubleHeroNode{0, "", "", nil, nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	var flag = false
	for {
		if tmp == nil {
			break
		}
		if tmp.No == hero.No {
			flag = true
			break
		}
		tmp = tmp.Next
	}
	if tmp != nil && flag {
		tmp.Pre.Next = tmp.Next
		if tmp.Next != nil {
			tmp.Next.Pre = tmp.Pre
		}
	} else {
		fmt.Println("get no hero")
	}
}

func (heroList *DoubleHeroNodeList) AddByOrder(hero DoubleHeroNode) error {
	if heroList.Head == nil {
		//判断链表是空，就初始化一个
		heroList.Head = &DoubleHeroNode{0, "", "", nil, nil}
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	flag := true
	for {
		if tmp.Next == nil {
			break
		}
		if tmp.Next.No == hero.No {
			flag = false
			break
		}
		if tmp.Next.No > hero.No {
			break
		}
		tmp = tmp.Next
	}
	if flag {
		//添加
		hero.Next = tmp.Next
		if tmp.Next != nil {
			tmp.Next.Pre = &hero
		}
		tmp.Next = &hero
		hero.Pre = tmp
	} else {
		return errors.New("标号存在，不能添加")
	}
	return nil
}

func (heroList *DoubleHeroNodeList) List() {
	//判断链表是否为空
	if heroList.Head == nil {
		fmt.Println("empty list")
		return
	}
	//临时变量，用于遍历
	tmp := heroList.Head
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Printf("%+v\n", tmp)
	}
}

//创建英雄struct
type DoubleHeroNode struct {
	No       int
	Name     string
	Nickname string
	Next     *DoubleHeroNode
	Pre      *DoubleHeroNode
}

func (hero *DoubleHeroNode) String() string {
	return fmt.Sprintf("no:%v,name:%s,Nickname:%s", hero.No, hero.Name, hero.Nickname)
}
