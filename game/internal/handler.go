package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/msg"
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.QuerychrMessage{}, handleQuerychr)
	handleMsg(&msg.SelchrMessage{}, handleSelchr)
	handleMsg(&msg.GameLoginMessage{}, handleGamelogin)
	handleMsg(&msg.LoginnoticeokMessage{}, handleLoginnoticeok)
	handleMsg(&msg.QueryBagitemsMessage{}, handleQuerybagitems)
}
func handleQuerychr(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.QuerychrMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("查询角色,账号:%s cert:%s", m.Lines[0], m.Lines[1])

	player1 := []string{"测试", "0", "2", "1000", "0"}
	player2 := []string{"测试1", "0", "2", "1000", "0"}
	players := []string{}
	players = append(players, player1...)
	players = append(players, player2...)
	querychrr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_QUERYCHR, 2, 0, 0, 0, 0, 0, 0, players...)
	a.WriteMsg(querychrr)
}
func handleSelchr(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.SelchrMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	log.Debug("选择角色,账号:%s 角色名称:%s", m.Lines[0], m.Lines[1])

	selchrr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_STARTPLAY, 0, 0, 0, 0, 0, 0, 0, "127.0.0.1", "7004")
	a.WriteMsg(selchrr)
}
func handleGamelogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.GameLoginMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("登录游戏服务器请求 账号:%s 名字:%s 版本:%s", m.Lines[0], m.Lines[1], m.Lines[2])

	notice := `测试公告`
	gameloginr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_SENDNOTICE, int32(len(notice)), 0, 0, 0, 0, 0, 0, notice)
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
