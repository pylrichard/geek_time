package main

import (
	"context"
	"fmt"
	"time"
)

type myKey int

func main() {
	keys := []myKey{
		myKey(20),
		myKey(30),
		myKey(60),
		myKey(61),
	}
	values := []string{
		"value in node2",
		"value in node3",
		"value in node6",
		"value in node7",
	}

	rootNode := context.Background()
	node1, cancelFunc1 := context.WithCancel(rootNode)
	defer cancelFunc1()

	// 示例1
	node2 := context.WithValue(node1, keys[0], values[0])
	node3 := context.WithValue(node2, keys[1], values[1])
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[0], node3.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[1], node3.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[2], node3.Value(keys[2]))
	fmt.Println()

	// 示例2
	node4, _ := context.WithCancel(node3)
	node5, _ := context.WithTimeout(node4, time.Hour)
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[0], node5.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[1], node5.Value(keys[1]))
	fmt.Println()

	// 示例3
	node6 := context.WithValue(node5, keys[2], values[2])
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[0], node6.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[2], node6.Value(keys[2]))
	fmt.Println()

	// 示例4
	node7 := context.WithValue(node5, keys[3], values[3])
	fmt.Printf("The value of the key %v found in the node7: %v\n",
		keys[1], node7.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node7: %v\n",
		keys[2], node7.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node7: %v\n",
		keys[3], node7.Value(keys[3]))
	fmt.Println()

	// 示例5
	node8, _ := context.WithCancel(node6)
	node9, _ := context.WithTimeout(node8, time.Hour)
	fmt.Printf("The value of the key %v found in the node9: %v\n",
		keys[1], node9.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node9: %v\n",
		keys[2], node9.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node9: %v\n",
		keys[3], node9.Value(keys[3]))
}
