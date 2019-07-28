package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("sfsfsdf")
	})
	fmt.Println("aaaa")
	time.Sleep(10 * time.Second)
}
