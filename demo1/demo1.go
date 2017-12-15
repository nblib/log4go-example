package demo1

import (
	l4g "github.com/nblib/log4go"
)

type Demo1 struct {

}
var logger l4g.Logger

func init()  {
	logger = make(l4g.Logger)
	logger.LoadConfiguration("demo1example.xml")
}
func (demo1 Demo1) Out()  {
	logger.Info("demo1 info")
	logger.Warn("demo1 warn")

	//logger.Close()
}