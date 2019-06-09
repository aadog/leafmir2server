package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"leafmir2server/conf"
	"leafmir2server/msg"
	"reflect"
	"strconv"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.SelResSessionMessage{}, handleResSession)
	handleMsg(&msg.QuerychrMessage{}, handleQuerychr)
	handleMsg(&msg.SelchrMessage{}, handleSelchr)
}

// 响应客户端sessionid请求
func handleResSession(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.SelResSessionMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	_ = m
	clisessoin := m.Lines[0]
	log.Debug("%s:获取到客户端session:%s", a.RemoteAddr().String(), clisessoin)
}

//查询角色
func handleQuerychr(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.QuerychrMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("查询角色,账号:%s cert:%s", m.Lines[0], m.Lines[1])

	player1 := []string{base.ConvertString2Byte("7>*清风", base.GBK), "0", "2", "13", "1"}
	player2 := []string{base.ConvertString2Byte("明月", base.GBK), "0", "2", "25", "1"}
	nplayer := int32(2)
	players := []string{}
	players = append(players, player1...)
	players = append(players, player2...)

	querychrr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_QUERYCHR, nplayer, 3, 1, 0, 0, 0, 0, players...)
	a.WriteMsg(querychrr)
}
func handleSelchr(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.SelchrMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	name := base.ConvertByte2String([]byte(m.Lines[0]), base.GBK)
	servername := base.ConvertByte2String([]byte(m.Lines[1]), base.GBK)
	log.Debug("选择角色,账号:%s 角色名称:%s", name, servername)

	selchrr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_STARTPLAY, 0, 0, 0, 0, 0, 0, 0, conf.Server.TcpAddr, strconv.Itoa(conf.Server.GameTcpPort))
	selchrr.Raw = base.EncodeString_uEDCode([]byte(selchrr.Stringlines()), []byte("tg%^&re3tb)(5G5R$7yg5.,,jIKh"))
	a.WriteMsg(selchrr)
}
