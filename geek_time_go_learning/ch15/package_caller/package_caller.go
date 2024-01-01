package package_caller

import (
	"fmt"
	pc "go/geek_time/geek_time_go_learning/ch15/package_callee"
)

func PackageCaller() {
	fmt.Println(pc.Square(5))
	fmt.Println(pc.Sum(1, 5))
}