package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	/*
		向通道先后发送三个元素值，通道变量名称、接送操作符、发送元素值三者之间最好用空格进行分割
	*/
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	//将最先进入channel1的元素2接收并赋值给变量elem
	elem := <-ch1
	fmt.Printf("The first element received from ch1: %v\n",
		elem)

	ch2 := make(chan int, 8)
	for i := 0; i < 8; i++ {
		ch2 <- i
	}
	fmt.Println(len(ch2))
	close(ch2)
	// 清空ch2
	for range ch2 {
	}
	fmt.Println(len(ch2))
}
