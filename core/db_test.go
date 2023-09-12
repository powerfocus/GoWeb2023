package core

import (
	"bufio"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"gweb/log"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestRoot(t *testing.T) {
	base := filepath.Dir("")
	abs, _ := filepath.Abs(base)
	t.Logf("%v", base)
	t.Logf("%v", abs)
}

func TestFile(t *testing.T) {
	stat, err := os.Stat("resources/application.yml")
	if os.IsExist(err) {
		t.Logf("%v", "文件存在")
		t.Logf("%v", stat.Name())
	} else {
		t.Logf("%v", "文件不存在")
	}
}

func TestAbsPath(t *testing.T) {
	dir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	t.Logf("%v", dir)
}

func TestPath(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	t.Logf("%s", path)
	dir, _ := os.Getwd()
	t.Logf("%s", dir)
}

func TestOpen(t *testing.T) {
	file, err := os.OpenFile("C:\\Users\\Administrator\\GolandProjects\\gweb\\application.yml", os.O_RDWR, 0)
	Defer(func() {
		t.Logf("%v", "defer操作...")
		err = file.Close()
	})
	ErrPanic(err)
	r := bufio.NewReader(file)
	buf := make([]byte, 128)
	for {
		_, err := r.Read(buf)
		if err != nil {
			break
		}
		t.Logf("%s", string(buf))
	}
}

func TestYaml(t *testing.T) {
	f, err := os.Open("C:\\Users\\Administrator\\GolandProjects\\gweb\\application.yml")
	ErrMsg(err, "打开文件是出现异常")
	defer func(f *os.File) {
		log.Println("defer操作")
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	/*scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		t.Logf("%v", text)
	}*/
	data, err := io.ReadAll(f)
	mp := make(map[string]interface{})
	err = yaml.Unmarshal(data, &mp)
	ErrPanic(err)
	/*for k, v := range mp {
		t.Logf("%v -> %v", k, v)
		if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
			for k, v := range v.(map[string]interface{}) {
				t.Logf("%v -> %v", k, v)
				if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
					for k, v := range v.(map[string]interface{}) {
						t.Logf("||%v -> %v", k, v)
					}
				}
			}
		}
	}*/
	/*t.Logf("%v", strings.Repeat("-", 20))
	IteratorMap(mp)
	t.Logf("%v", strings.Repeat("-", 20))*/
	url := StringKeyToVal(mp, "url")
	for k, v := range url {
		t.Logf("%v: %v\n", k, v)
	}
	t.Logf("%v", strings.Repeat("-", 20))
	Reset()
	host := StringKeyToVal(mp, "host")
	for k, v := range host {
		t.Logf("%v: %v\n", k, v)
	}
}

func TestType(t *testing.T) {
	b := reflect.TypeOf(map[string]any{}) == reflect.TypeOf(map[string]any{})
	t.Logf("%v", b)
	t.Logf("%v, %v", reflect.TypeOf(map[string]any{}).Kind(), reflect.TypeOf(map[string]any{}).Kind())
}

func TestMap(t *testing.T) {
	m := map[string]any{}

	l1 := map[string]any{}
	l1["name"] = "win10"
	l1["address"] = "localhost"

	l2 := map[string]any{}
	l2["username"] = "nacos"
	l2["password"] = "123"
	l1["l2"] = l2

	l3 := map[string]any{}
	l3["address"] = "127.0.0.1"
	l3["port"] = "8080"

	m["l1"] = l1
	m["l2"] = l2
	m["l3"] = l3

	/*a := m["l1"]
	t.Logf("%v", strings.Repeat("-", 20))
	t.Logf("%v type: %v, kind: %v", a, reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	t.Logf("name: %v, String: %v", reflect.TypeOf(a).Name(), reflect.TypeOf(a).String())
	t.Logf("%v", reflect.TypeOf(a).Kind().String())

	t.Logf("%v, %v", a, reflect.TypeOf(a))
	v := a.(map[string]any)["name"]
	t.Logf("%v", v)

	t.Logf("%v", strings.Repeat("-", 20))*/
	//IterMap(m)

	v := FindMap(m, "address")
	t.Logf("%v", v)
	t.Logf("%v", strings.Repeat("-", 20))
	if v, ok := m["l1"]; ok {
		t.Logf("%v", v)
	} else {
		t.Logf("%v", "未找到")
	}
	t.Logf("%v", strings.Repeat("-", 20))
	val := RGetMapVal(m, "address")
	t.Logf("%v", val)
}

func TestSlice(t *testing.T) {
	var lst []any
	t.Logf("%v, %v", len(lst), cap(lst))
	lst2 := make([]any, 10, 20)
	t.Logf("%v, %v", len(lst2), cap(lst2))
}

func TestPathWalk(t *testing.T) {
	dir, err := os.ReadDir("C:\\Users\\Administrator\\Documents\\codes\\gin")
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() {
			t.Logf("[%v]", file.Name())
		} else {
			t.Logf("%v", file.Name())
		}
	}
}

func TestFilePath(t *testing.T) {
	/*if dir, err := filepath.Abs("."); err != nil {
		panic(err)
	} else {
		t.Logf("%v", dir)
	}*/
	err := filepath.Walk("C:\\Users\\Administrator\\Documents\\codes\\gin", func(p string, info fs.FileInfo, err error) error {
		stat, err := os.Stat(p)
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			t.Logf("[%v]", filepath.Base(p))
		} else {
			t.Logf("%v", filepath.Base(p))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestViper(t *testing.T) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("C:\\Users\\Administrator\\GolandProjects\\gweb\\")
	dir, _ := os.Getwd()
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v := viper.Get("goWeb.dataSource.url")
	t.Logf("%v", v)
}

func TestF1(t *testing.T) {
	file, err := os.OpenFile("application1.yml", os.O_CREATE, 0777)
	if err != nil {
		t.Logf("%s %v", "打开文件时错误 ", err)
	}
	file.WriteString("hello world")
	t.Logf("%v", file.Name())
}
