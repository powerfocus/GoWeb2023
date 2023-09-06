package core

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
