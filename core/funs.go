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
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(mp).Kind() {
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
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(mp).Kind() {
			IteratorMap(v.(map[string]any))
		} else {
			log.Println(k, " -> ", v)
		}
	}
}

func FindMap(mp map[string]any, tgKey string) []any {
	//var re []any
	re := make([]any, 0)
	for k, v := range mp {
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(mp).Kind() {
			if rv := FindMap(v.(map[string]any), tgKey); rv != nil && len(rv) != 0 {
				re = append(re, rv)
			}
		} else {
			if strings.EqualFold(k, tgKey) {
				log.Println("找到目标 ", v)
				re = append(re, v)
			}
		}
	}
	return re
}

// GetMapVal 如果存在key就返回值
func GetMapVal[K string | int, V any](mp map[K]V, key K) any {
	var re any
	if v, ok := mp[key]; ok {
		re = v
	}
	return re
}

func RGetMapVal(mp map[string]any, key string) []any {
	var re []any
	for k, v := range mp {
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(mp).Kind() {
			if r := RGetMapVal(v.(map[string]any), key); r != nil && len(r) != 0 {
				re = append(re, RGetMapVal(v.(map[string]any), key))
			}
		} else {
			if strings.EqualFold(k, key) {
				re = append(re, v)
			}
		}
	}
	return re
}

func CleanMap[K string | int, V any](mp map[K]V) {
	for k := range mp {
		delete(mp, k)
	}
}
func IterMap(mp map[string]any) {
	for k, v := range mp {
		if reflect.TypeOf(v).Kind() == reflect.TypeOf(mp).Kind() {
			IterMap(v.(map[string]any))
		} else {
			log.Println(k, " -> ", v)
		}
	}
}
