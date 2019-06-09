package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"leafmir2server/msg"
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

	handleMsg(&msg.GameSessionMessage{}, handleGameSession)
	handleMsg(&msg.GameLoginMessage{}, handleGamelogin)
	handleMsg(&msg.LoginnoticeokMessage{}, handleLoginnoticeok)
	handleMsg(&msg.QueryBagitemsMessage{}, handleQuerybagitems)
}

func handleGameSession(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.GameSessionMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	strsession := m.Lines[0]
	log.Debug("%s:发送来的session:%s", a.RemoteAddr().String(), strsession)
}

func handleGamelogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.GameLoginMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	strname := base.ConvertByte2String([]byte(m.Lines[1]), base.GBK)
	log.Debug("%s:登录游戏服务器请求 账号:%s 名字:%s 版本:%s", a.RemoteAddr().String(), m.Lines[0], strname, m.Lines[2])

	notice := `ABody '\\\\　　　　　 健康游戏忠告 \\\　　注意自身保护　谨防受骗上当\\　　适度游戏益脑　沉迷游戏伤身\      \　　合理安排时间　享受健康生活\'`
	gameloginr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_SENDNOTICE, 2000, 0, 0, 0, 0, 0, 0, notice)
	a.WriteMsg(gameloginr)
}
func handleLoginnoticeok(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.LoginnoticeokMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	_ = m
	btbuf := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	loginnoticeokr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_LOGON, 1000, 273, 590, ((1 << 8) | 1), 0, 0, 0, string(btbuf))
	a.WriteMsg(loginnoticeokr)

}
func handleQuerybagitems(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.QueryBagitemsMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	_ = a

	log.Debug("%v", m)
	querybagitemsr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_BAGITEMS, 1, 0, 0, 0, 0, 0, 0)
	a.WriteMsg(querybagitemsr)
}
