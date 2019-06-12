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

	notice := `游戏公告`
	notice = base.ConvertString2Byte(notice, base.GBK)
	gameloginr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_SENDNOTICE, 2000, 0, 0, 0, 0, 0, 0, notice)
	a.WriteMsg(gameloginr)
}
func handleLoginnoticeok(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.LoginnoticeokMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	_ = a
	_ = m
	//strusername:="asdfsadf\\\\\\"
	//loginnoticeokr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_USERNAME, 734707136, 255, 0, 0, -1, 0, 0, string(strusername))
	//a.WriteMsg(loginnoticeokr)
	log.Debug("客户端点击公告按钮")
	//发送人物出生地图
	newmapr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_NEWMAP, 0, 649, 626, 0, 0, 0, 0)
	mainmapname := "0" //比奇省
	newmapr.Add_with_line(mainmapname)
	a.WriteMsg(newmapr)

	//logonr:=msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_LOGON,260964368,41026185,7,0,11966752,0,0)
	//
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
