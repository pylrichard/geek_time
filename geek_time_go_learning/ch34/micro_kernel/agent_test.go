package micro_kernel

import (
	"fmt"
	"testing"
	"time"
)

func TestAgent(t *testing.T) {
	a := NewAgent(100)
	c1 := NewDemoCollector("c1", "1")
	c2 := NewDemoCollector("c2", "2")
	a.RegisterCollector("c1", c1)
	a.RegisterCollector("c2", c2)
	if err := a.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	fmt.Println(a.Start())
	time.Sleep(1 * time.Second)
	a.Stop()
	a.Destroy()
}