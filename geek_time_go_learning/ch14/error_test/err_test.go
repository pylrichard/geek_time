package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/*
	定义不同的错误变量，以便于判断错误类型
 */
var LessThan2Error = errors.New("param should be not less than 2")
var LargerThan100Error = errors.New("param should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThan2Error
	}
	if n > 100 {
		return nil, LargerThan100Error
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i - 2] + fibList[i - 1])
	}

	return fibList, nil
}

/*
	及早失败，避免嵌套
 */
func GetFibonacci1(s string) {
	var (
		i		int
		err		error
		list	[]int
	)
	if i, err = strconv.Atoi(s); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func GetFibonacci2(s string) {
	var (
		i		int
		err		error
		list	[]int
	)
	if i, err = strconv.Atoi(s); err != nil {
		fmt.Println(err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}

func TestGetFibonacci(t *testing.T) {
	if list, err := GetFibonacci(1); err != nil {
		if err == LessThan2Error {
			t.Error("need a larger num")
		}
		if err == LargerThan100Error {
			t.Error("need a less num")
		}
	} else {
		t.Log(list)
	}
	var s string = "a"
	GetFibonacci1(s)
	GetFibonacci2(s)
}