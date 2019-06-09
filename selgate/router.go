package selgate

import (
	"leafmir2server/msg"
	"leafmir2server/sel"
)

func init() {
	msg.Processor.SetRouter(&msg.SelResSessionMessage{}, sel.ChanRPC)
	msg.Processor.SetRouter(&msg.QuerychrMessage{}, sel.ChanRPC)
	msg.Processor.SetRouter(&msg.SelchrMessage{}, sel.ChanRPC)
}
