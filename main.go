package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"leafmir2server/conf"
	"leafmir2server/game"
	"leafmir2server/gamegate"
	"leafmir2server/login"
	"leafmir2server/logingate"
	"leafmir2server/sel"
	"leafmir2server/selgate"
	"log"
)

func main() {

	lconf.LogFlag = log.Lshortfile
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	//初始化路由
	router()
	leaf.Run(
		&sel.Module{},
		&selgate.Module{},
		&logingate.Module{},
		&login.Module{},
		&gamegate.Module{},
		&game.Module{},
	)

}
