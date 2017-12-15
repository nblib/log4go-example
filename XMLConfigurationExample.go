package main

import (
	l4g "github.com/nblib/log4go"
	"time"
)

func main() {
	// Load the configuration (isn't this easy?)
	l4g.LoadConfiguration("example.xml")

	//for true {
	//	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	//	l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	//	l4g.Info("About that time, eh chaps?")
	//	time.Sleep(5 * time.Millisecond)
	//}
	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	l4g.Info("About that time, eh chaps?")
	time.Sleep(1 * time.Second)
	// And now we're ready!

}
