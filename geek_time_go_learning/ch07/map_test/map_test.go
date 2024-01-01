package map_test

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	t.Log(m1[2])
	t.Logf("len of m1 = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len of m2 = %d", len(m2))

	//10代表Initial Capacity，map无法像slice那样为len个元素设置零值
	m3 := make(map[int]int, 10)
	t.Logf("len of m3 = %d", len(m3))
}

func TestMapExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	//访问的key不存在时，仍会返回零值，不能通过返回nil判断元素是否存在
	//m[3] = 0
	if v, ok := m[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[int]int{1:1, 2:4, 3:9}
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestForLoop(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	var dst []*int
	for _, i := range src {
		dst = append(dst, &i)
	}
	for _, p := range dst {
		fmt.Println(*p)
	}

	/*
		for range迭代逻辑
	 */
	var n int
	for j := 0; j < len(src); j++ {
		/*
			遍历过程中并没有返回集合元素，而是将元素值复制给一个固定变量
			这个复制操作就是for range存在性能问题的原因
		 */
		n = src[j]
		dst = append(dst, &j)
		fmt.Println(n)
	}
	/*
		以为的迭代逻辑
		for j := 0; j < len(src); j++ {
			dst = append(dst, &src[j])
		}
	 */
}