package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
}
