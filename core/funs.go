package core

import (
	"gweb/log"
	"reflect"
	"strings"
)

var (
	val map[string]any
	nav string
)

type CommonFunc func()
type OneFunc func(arg any)
type MultiFunc func(args ...any)

func init() {
	val = make(map[string]any)
}

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

// Reset 多次使用时，必须调用来清楚旧数据
func Reset() {
	CleanMap(val)
	nav = ""
}

func StringKeyToVal(mp map[string]any, key any) map[string]any {
	for k, v := range mp {
		nav += k
		nav += "/"
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(map[string]any{}).Kind() {
			StringKeyToVal(v.(map[string]any), key)
			nav = ""
		} else {
			if k == key {
				val[nav] = v
				nav = strings.Replace(nav, "url/", "", 1)
			}
		}
	}
	return val
}

func IteratorMap(mp map[string]any) {
	for k, v := range mp {
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(map[string]any{}).Kind() {
			IteratorMap(v.(map[string]any))
		} else {
			log.Println(k, " -> ", v)
		}
	}
}

func FindMap(mp map[string]any) any {
	var re any

	return re
}

func CleanMap[K string | int, V any](mp map[K]V) {
	for k, _ := range mp {
		delete(mp, k)
	}
}
