package sel

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}
func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("selserver,来自客户端连接:%s", a.RemoteAddr())
}
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}
