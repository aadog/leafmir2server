package logingate

import (
	"fmt"
	"github.com/a97077088/leaf/gate"
	"github.com/patrickmn/go-cache"
	"leafmir2server/conf"
	"leafmir2server/login"
	"leafmir2server/msg"
	"time"
)

type Module struct {
	*gate.Gate
}

var ca *cache.Cache

func init() {
	ca = cache.New(time.Hour*1000, time.Hour*1000)
}
func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         fmt.Sprintf("%s:%d", conf.Server.TcpAddr, conf.Server.LoginTCPPort),
		Processor:       msg.Processor,
		AgentChanRPC:    login.ChanRPC,
		MsgParser:       &MsgParser{},
	}

}
