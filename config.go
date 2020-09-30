package workflow

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// SetServerOptions 读取配置文件
func SetServerOptions(configsPath string) {
	if configsPath == "" {
		setDefaultOptions()
		return
	}
	filePaths := strings.Split(configsPath, "/")
	file := filePaths[len(filePaths)-1]
	fileDir := strings.Replace(configsPath, file, "", -1)
	fileNames := strings.SplitN(file, ".", 2)
	fileName := fileNames[0]
	fileType := fileNames[1]

	viper.SetConfigName(fileName)
	//supports JSON, TOML, YAML, HCL, INI, envfile and Java Properties files
	if fileType != "" {
		viper.SetConfigType(fileType)
	}
	viper.AddConfigPath(fileDir)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		setDefaultOptions()
	}
}

// setDefaultOptions 生成默认配置
func setDefaultOptions() {
	// 读取之前生成的默认配置
	viper.SetConfigName("sflow")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err == nil {
		return
	}
	// 生成默认配置
	viper.SetDefault("mode", "debug")
	viper.SetDefault("dsn", "root:123456@tcp(127.0.0.1:3306)/sflow?charset=utf8")
	viper.SetDefault("port", "6062")
	viper.SafeWriteConfig()
}
