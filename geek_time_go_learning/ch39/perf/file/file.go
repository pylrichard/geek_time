package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	//行数和列数少了，prof没有信息
	row = 10000
	col = 10000

)

func fillMatrix(x *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			x[i][j] = s.Intn(10000)
		}
	}
}

func calculate(x *[row][col]int) {
	for i := 0; i < row; i++ {
		sum := 0
		for j := 0; j < col; j++ {
			sum += x[i][j]
		}
	}
}

/*
	go build -o file.exe file.go && ./file.exe
	go tool pprof file.exe cpu.prof
	top -cum
	list fillMatrix
	go tool pprof -http=":8081" file.exe cpu.prof
 */
func main() {
	f1, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create cpu profile: ", err)
	}

	if err := pprof.StartCPUProfile(f1); err != nil {
		log.Fatal("could not start cpu profile: ", err)
	}
	defer pprof.StopCPUProfile()

	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f2, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	//GC降低内存占用
	runtime.GC()
	if err := pprof.WriteHeapProfile(f2); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f2.Close()

	f3, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile: ", err)
	}
	//支持多种profile，在runtime/pprof.go中查看
	if gprof := pprof.Lookup("goroutine"); gprof == nil {
		log.Fatal("could not write goroutine profile: ")
	} else {
		gprof.WriteTo(f3, 0)
	}
	f3.Close()
}