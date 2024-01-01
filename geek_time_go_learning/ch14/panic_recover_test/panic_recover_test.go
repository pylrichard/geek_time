package panic_recover_test_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicAndExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("panic"))
	//os.Exit(-1)
	fmt.Println("End")
}