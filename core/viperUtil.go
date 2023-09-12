package core

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

var (
	configName  string
	configType  string
	configPaths []string
)

type ViperUtil struct {
	ConfigName  string
	ConfigType  string
	ConfigPaths []string
}

// viper 读取配置文件功能封装
func init() {
	configName = "application"
	configType = "yaml"
	absDir, err := filepath.Abs("")
	if err != nil {
		panic(err)
	}
	// 设置配置文件加载路径
	configPaths = append(configPaths, ".")
	configPaths = append(configPaths, "./config")
	configPaths = append(configPaths, absDir)
	configPaths = append(configPaths, filepath.Dir(absDir))

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	for _, v := range configPaths {
		viper.AddConfigPath(v)
	}
	//viper.AddConfigPath(filepath.Join(absDir, "/../"))
	viper.AddConfigPath(absDir)
	viper.AddConfigPath(filepath.Join(absDir, "./config"))

	if err = viper.ReadInConfig(); err != nil {
		if fErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("配置文件未找到 %v", fErr)
		} else {
			log.Printf("读取配置文件时发生错误 %v", fErr)
		}
	}
}

func NewViperUtil() *ViperUtil {
	instance := new(ViperUtil)
	instance.ConfigName = configName
	instance.ConfigType = configType
	copy(instance.ConfigPaths, configPaths)
	return instance
}

func (t *ViperUtil) Get(key string) any {
	return viper.Get(key)
}
