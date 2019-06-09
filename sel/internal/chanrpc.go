package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"sync"
)

var tcpInfo sync.Map

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}
func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	sti := base.TcpInfo{
		a.RemoteAddr().String(),
		"",
		"",
	} //建立客户端tcpinfo
	SetTcpInfo(a.RemoteAddr().String(), sti)
	log.Debug("selserver,来自客户端连接:%s", a.RemoteAddr())
}
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	//清除tcpinfo
	tcpInfo.Delete(a.RemoteAddr().String())
	//log.Debug("清空gettcpinfo")
	//清除完成
}
func GetTcpInfo(_key string) base.TcpInfo {
	ti, ok := tcpInfo.Load(_key)
	if ok == false {
		return base.TcpInfo{}
	}
	return ti.(base.TcpInfo)
}
func SetTcpInfo(_key string, _val base.TcpInfo) {
	tcpInfo.Store(_key, _val)
}
func CopyTcpInfo(_key string) base.TcpInfo {
	return GetTcpInfo(_key)
}
