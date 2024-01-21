package lib

import (
	in "go/geek_time/geek_time_go_puzzlers/2-lib/lib3/lib/internal"
	"os"
)

func SayHello(name string) {
	in.SayHello(os.Stdout, name)
}
