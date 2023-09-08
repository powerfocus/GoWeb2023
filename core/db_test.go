package core

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"reflect"
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
	if err != nil {
		panic(err)
	}
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
	/*if v, ok := mp["goWeb"]; ok {
		t.Logf("%v, %v", reflect.TypeOf(v), v)
	}*/
	for k, v := range mp {
		t.Logf("%v -> %v", k, v)
		for k, v := range v.(map[string]interface{}) {
			t.Logf("%v -> %v", k, v)
			if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
				for k, v := range v.(map[string]interface{}) {
					t.Logf("%v -> %v", k, v)
				}
			}
		}
	}
}
