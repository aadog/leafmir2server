package gamegate

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/patrickmn/go-cache"
	"leafmir2server/conf"
	"leafmir2server/game"
	"leafmir2server/msg"
	"time"
)

var ca *cache.Cache

func init() {
	ca = cache.New(time.Hour*1000, time.Hour*1000)
}

type Module struct {
	*gate.Gate
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
		TCPAddr:         fmt.Sprintf("%s:%d", conf.Server.TcpAddr, conf.Server.GameTcpPort),
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
		MsgParser:       &MsgParser{},
	}

}
