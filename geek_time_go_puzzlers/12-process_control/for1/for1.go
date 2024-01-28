package main

import (
	"fmt"
)

func main() {
	// 示例1
	fmt.Println("Example 1")
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	// [1 2 3 7 5 6]
	fmt.Println(numbers1)
	fmt.Println()

	/*
		示例2
		被迭代对象是range表达式结果值的副本而不是原值
		第一次迭代时e为1，改变numbers2第二个元素的值，新值为3
		但被迭代对象(复制的numbers2副本)的第二个元素没有改变，还是2
		第二次迭代时e为2，不为3
	*/
	fmt.Println("Example 2")
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	// [7 3 5 7 9 11]
	fmt.Println(numbers2)
	fmt.Println()

	/*
		示例3
		注意：切片是引用类型，数组是值类型
		切片内部结构为struct slice{*point, len, cap}。
		数据部分是一个指针，复制对象时把指针值复制了
		遍历切片之前，务必检查切片是否为空或nil
	*/
	fmt.Println("Example 3")
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1
	// 第二次迭代时e为3，修改第三个元素的值为6
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	// [22 3 6 10 15 21]
	fmt.Println(numbers3)
	fmt.Println()

	// 示例4
	fmt.Println("Example 4")
	numbers4 := [6]int{1, 2, 3, 4, 5, 6}
	for _, e := range numbers4 {
		//e是复制对象中的值，不会改变numbers4数组元素的值
		e += 1
	}
	// [1, 2, 3, 4, 5, 6]
	fmt.Println(numbers4)
	fmt.Println()

	numbers5 := []int{1, 2, 3, 4, 5, 6}
	for _, e := range numbers5 {
		e += 1
	}
	// [1, 2, 3, 4, 5, 6]
	fmt.Println(numbers5)
	fmt.Println()

	// 示例5
	fmt.Println("Example 5")
	m := make(map[int]*int)
	// for k, v := range numbers5 {
	// 	fmt.Println(v, &v)
	// 	m[k] = &v
	// }
	// 见标准库\for.md
	for k, v := range numbers5 {
		n := v
		m[k] = &n
	}
	// Go哈希表遍历是无序的，Go团队在设计哈希表的遍历操作时特意引入了随机数以保证遍历的随机性
	// 见标准库\for.md
	for k, v := range m {
		fmt.Printf("map[%v]=%v\n", k, *v)
	}
	fmt.Println(m)
	fmt.Println()

	// 示例6
	fmt.Println("Example 6")
	len := 8
	ch := make(chan []int, len)
	for i := 0; i < len; i++ {
		s := make([]int, 1)
		s[0] = i
		ch <- s
	}
	// 关闭通道避免遍历死锁
	close(ch)
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	// 被编译器转换为等价三段for循环的形式：使用<-ch阻塞通道并获值，根据返回的布尔值hb判断当前通道是否有值
	// 如果有值则赋值hv并执行循环体，而后重新陷入阻塞等待新数据
	for hv, hb := <-ch; hb; hv, hb = <-ch {
		v := hv
		fmt.Println(hb)
		// nil要赋值给切片等集合类型
		hv = nil
		fmt.Println(v)
	}
	v, ok := <-ch
	fmt.Println(v, ok)
}
