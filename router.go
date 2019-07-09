package main

import (
	"leafmir2server/game"
	"leafmir2server/login"
	"leafmir2server/msg"
	"leafmir2server/sel"
)

func router() {
	//登录网关
	msg.Processor.SetRouter(&msg.ResTokenMessage{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginMessage{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.SelectServerMessage{}, login.ChanRPC)
	//角色网关
	msg.Processor.SetRouter(&msg.SelResSessionMessage{}, sel.ChanRPC)
	msg.Processor.SetRouter(&msg.QuerychrMessage{}, sel.ChanRPC)
	msg.Processor.SetRouter(&msg.SelchrMessage{}, sel.ChanRPC)
	//游戏网关
	msg.Processor.SetRouter(&msg.GameSessionMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GameLoginMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginnoticeokMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.QueryBagitemsMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ClientdataVersionMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.WantMinimapMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HerodurachangeMessage{}, game.ChanRPC)
}
