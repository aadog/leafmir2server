package gamegate

import (
	"leafmir2server/game"
	"leafmir2server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.GameLoginMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginnoticeokMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.QueryBagitemsMessage{}, game.ChanRPC)
}
