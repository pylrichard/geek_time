package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr2[1] = 5
	t.Log(arr1[1], arr1[2])
	t.Log(arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	for _, elem := range arr {
		t.Log(elem)
	}
}

func TestArraySection(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := arr1[:]
	t.Log(arr2)
}