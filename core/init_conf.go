package core

import (
	"blogx_server/conf"
	"blogx_server/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadConf() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误 %s", err))
	}
	fmt.Printf("配置文件读取成功 %s 成功\n", flags.FlagOptions.File)
	return
}
