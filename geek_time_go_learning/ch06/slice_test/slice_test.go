package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s1 := append(s0, 1)
	t.Log(len(s1), cap(s1))

	s2 := []int{1, 2, 3, 4}
	t.Log(len(s2), cap(s2))

	//make([]type, len, cap)
	//len个元素会被初始化为默认零值，未初始化元素不可以访问
	s3 := make([]int, 3, 5)
	t.Log(len(s3), cap(s3))
	t.Log(s3[0], s3[1], s3[2])
	s3 = append(s3, 1)
	t.Log(s3[0], s3[1], s3[2], s3[3])
	t.Log(len(s3), cap(s3))
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(i, len(s), cap(s))
	}
}

func TestSliceShareMem(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar",
					"Apr", "May", "Jun",
					"Jul", "Aug", "Sep",
					"Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknown"
	t.Log(q2)
	t.Log(year)
}

func TestSliceCmp(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//切片只能和nil比较
	//if a == b {
	//	t.Log("equal")
	//}
	t.Log(a, b)
}