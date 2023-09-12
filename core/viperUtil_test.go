package core

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"testing"
)

func TestStdFilePath(t *testing.T) {
	absDir, err := filepath.Abs("../")
	if err != nil {
		panic(err)
	}
	t.Logf("%v", absDir)
}

func TestViperUtil(t *testing.T) {
	absDir, err := filepath.Abs("../")
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("application")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(absDir)
	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("配置文件未找到 ", err)
		} else {
			log.Fatalln(err)
		}
	}
	val := viper.Get("goWeb.dataSource.url")
	t.Logf("%v", val)
}

func TestViperUtil_Get(t *testing.T) {
	viperUtil := NewViperUtil()
	url := viperUtil.Get("goWeb.dataSource.url")
	t.Logf("%v", url)
}
