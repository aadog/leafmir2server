package gate

import (
	"leafmir2server/login"
	"leafmir2server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.LoginMessage{},login.ChanRPC)
	msg.Processor.SetRouter(&msg.SelectServerMessage{},login.ChanRPC)
}
