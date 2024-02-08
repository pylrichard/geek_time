package main

import (
	"strings"
)

func main() {
	// 示例1
	var builder1 strings.Builder
	builder1.Grow(1)

	f1 := func(b strings.Builder) {
		// 这里会引发panic
		//b.Grow(1)
	}
	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1
	builder2 := <-ch1
	// 这里会引发panic
	//builder2.Grow(1)
	_ = builder2

	builder3 := builder1
	// 这里会引发panic
	//builder3.Grow(1)
	_ = builder3

	// 示例2
	f2 := func(bp *strings.Builder) {
		// 这里虽然不会引发panic，但不是并发安全的
		(*bp).Grow(1)
		builder4 := *bp
		// 这里会引发panic
		//builder4.Grow(1)
		_ = builder4
	}
	f2(&builder1)

	builder1.Reset()
	builder5 := builder1
	// 这里不会引发panic
	builder5.Grow(1)
}
