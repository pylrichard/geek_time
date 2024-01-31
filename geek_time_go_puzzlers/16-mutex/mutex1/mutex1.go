package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
)

const (
	// 代表启用的goroutine数量
	GO_NUM = 5
	// 代表每个goroutine需要写入的数据块数量
	DATA_BLOCK_NUM = 10
	// 代表每个数据块中需要有多少个重复数字
	REPEAT_NUM_IN_DATA_BLOCK_NUM = 10
)

// protecting 用于指示是否使用互斥锁来保护数据写入
// 若值等于0则表示不使用，若值大于0则表示使用
// 改变该变量的值，然后多运行几次程序，并观察程序打印的内容
var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1,
		"It indicates whether to use a mutex to protect data writing.")
}

func main() {
	flag.Parse()

	var buffer bytes.Buffer
	var mu sync.Mutex
	sign := make(chan struct{}, GO_NUM)

	for i := 1; i <= GO_NUM; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= DATA_BLOCK_NUM; j++ {
				// 准备数据
				header := fmt.Sprintf("\n[id: %d, iteration: %d]",
					id, j)
				data := fmt.Sprintf(" %d", id*j)
				// 写入数据
				if protecting > 0 {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				for k := 0; k < REPEAT_NUM_IN_DATA_BLOCK_NUM; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}

	for i := 0; i < GO_NUM; i++ {
		<-sign
	}
	data, err := io.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}
