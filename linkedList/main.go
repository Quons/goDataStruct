package main

import (
	"errors"
	"fmt"
)

func main() {
	var heroList HeroNodeList
	h1 := HeroNode{1, "林冲", "豹子头", nil}
	h2 := HeroNode{2, "鲁智深", "花和尚", nil}
	h3 := HeroNode{3, "张璐", "玉麒麟", nil}
	h4 := HeroNode{4, "武松", "老虎克星", nil}

	h5 := HeroNode{5, "武松", "老虎克星", nil}
	h6 := HeroNode{6, "武松2", "老虎克星", nil}
	if err := heroList.Add(h1); err != nil {
		fmt.Print(err)
	}
	if err := heroList.Add(h2); err != nil {
		fmt.Print(err)
	}
	if err := heroList.Add(h3); err != nil {
		fmt.Print(err)
	}
	if err := heroList.Add(h4); err != nil {
		fmt.Print(err)
	}
	if err := heroList.AddByOrder(h6); err != nil {
		fmt.Print(err)
	}
	if err := heroList.AddByOrder(h5); err != nil {
		fmt.Print(err)
	}
	h6.Name = "刘恒越"
	heroList.Update(h6)
	heroList.List()
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
