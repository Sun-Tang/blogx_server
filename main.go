package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"

	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	logrus.Warnf("blogx_server xxxx")
	logrus.Debug("blogx_server yyyy")
	logrus.Error("blogx_server zzzz")
	logrus.Infof("blogx_server aaaa")
}
