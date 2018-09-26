package logs

import (
	"log"
	"os"
)

//Instance ...
type Instance struct {
	//L ...
	L *log.Logger
}

//Inst ...
var Inst Instance

//NewInstance ...
func NewInstance() *Instance {
	l := log.New(os.Stdout, "", log.LstdFlags)
	Inst.L = l
	return &Instance{
		L: l,
	}
}

//GetInstance ...
func GetInstance() *log.Logger {
	if Inst == (Instance{}) {
		return NewInstance().L
	}
	return Inst.L
}
