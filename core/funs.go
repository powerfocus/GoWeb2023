package core

import "gweb/log"

type CommonFunc func()
type OneFunc func(arg any)
type MultiFunc func(args ...any)

// Defer defer()的封装函数
func Defer(f CommonFunc) {
	defer func() {
		f()
	}()
}

// Try recover()的封装函数
func Try(f OneFunc) {
	Defer(func() {
		err := recover()
		f(err)
	})
}

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrPanicMsg(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func ErrMsg(err error, msg string) {
	if err != nil {
		log.Println(msg, " ", err.Error())
	}
}
