package main

import (
	"github.com/nblib/log4go-example/demo2"
	"time"
	"fmt"
)

func main() {

	//demo1 := demo1.Demo1{}
	//for i := 0; i < 10000; i++ {
	//demo1.Out()
	//}

	/*for i := 0; i < 1000; i++ {
		go func() {
			dem2 := demo2.Demo2{}
			for true {
				dem2.Out()

			}
		}()
	}*/
	dem2 := demo2.Demo2{}
	for true {

		dem2.Out()
		fmt.Println(".")
	}

	time.Sleep(200 * time.Second)
}
