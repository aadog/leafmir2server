package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"leafmir2server/conf"
	"leafmir2server/game"
	"leafmir2server/gamegate"
	"leafmir2server/gate"
	"leafmir2server/login"
	"log"
)

func main() {
	lconf.LogFlag = log.Lshortfile
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		gate.Module,
		gamegate.Module,
		login.Module,
		game.Module,
	)

}
