package flags

import "flag"

type Options struct {
	File string
	DB   bool
	Version bool
}

var FlagOptions *Options

func Parse() {
	// 初始化 FlagOptions 避免空指针解引用
	FlagOptions = new(Options)
	
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本号")
	flag.Parse()
}