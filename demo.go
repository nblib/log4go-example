package main

import (
	"github.com/log4go-example/demo1"
	"github.com/log4go-example/demo2"
	"time"
)

func main() {

	demo1 := demo1.Demo1{}
	//for i := 0; i < 10000; i++ {
	demo1.Out()
	//}

	dem2 := demo2.Demo2{}
	//for i := 0; i < 10000; i++ {
	dem2.Out()
	//}
	time.Sleep(1 * time.Second)
}
