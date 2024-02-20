package main

func main() {
	//示例1
	ch1 := make(chan int, 1)
	ch1 <- 1
	//通道已满，会造成阻塞
	// channel1 <- 2

	//示例2
	ch2 := make(chan int, 1)
	//通道已空，会造成阻塞
	// element, ok := <-channel2
	//避免变量未使用报错
	// _, _ = element, ok
	ch2 <- 1

	//示例3
	var ch3 chan int
	//通道值为nil，会造成永久阻塞
	// channel3 <- 1
	//通道值为nil，会造成永久阻塞
	// <-channel3
	_ = ch3
}
