package demo2

import (
	l4g "github.com/nblib/log4go"
)

type Demo2 struct {

}
var logger l4g.Logger

func init()  {
	logger = make(l4g.Logger)
	logger.LoadConfiguration("demo2example.xml")
}
func (demo2 Demo2) Out()  {
	logger.Info("demo2 info")
	logger.Warn("demo2 warn")
}