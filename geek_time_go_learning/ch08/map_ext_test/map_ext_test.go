package map_ext_test_test

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](2), m[2](2))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	//添加元素
	set[1] = true
	n := 3
	/*
		判断元素是否存在
	 */
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	set[3] = true
	//获取元素个数
	t.Log(len(set))
	//删除元素
	delete(set, 1)
	n = 1
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}