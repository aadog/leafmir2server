package login

import (
	"encoding/hex"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"leafmir2server/conf"
	"leafmir2server/msg"
	"reflect"
	"strconv"
	"strings"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.LoginMessage{}, handleLogin)
	handleMsg(&msg.SelectServerMessage{}, handleSelectserver)
	handleMsg(&msg.ResTokenMessage{}, handleRestoken)
}
func handleSelectserver(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.SelectServerMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	servername := m.Lines[0]
	log.Debug("选择的服务器:%s", base.ConvertByte2String([]byte(servername), base.GBK))

	loginserverip := conf.Server.TcpAddr
	loginserverport := strconv.Itoa(conf.Server.SelTCPPort)
	sessionid := "session"
	selectserverr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_SELECTSERVER_OK, 25, 0, 0, 0, 0, 0, 0,
		loginserverip,
		loginserverport,
		"8",
		sessionid,
	)
	a.WriteMsg(selectserverr)
}
func handleLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.LoginMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	log.Debug("开始登录 客户端版本:%v 账号:%v,密码:%v", m.Recog, m.Lines[0], m.Lines[1])

	//loginerrr:=msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_PASSWD_FAIL,-1,0,0,0,0,0,0)
	//a.WriteMsg(loginerrr)

	nserver := int32(1)
	server1status := "1"
	server1name := base.ConvertString2Byte("龙际网络", base.GBK)
	//// 给发送者回应一个 Hello 消息
	loginr := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_PASSOK_SELECTSERVER, 0, 0, 0, nserver, 0, 0, 0, server1name, server1status)

	a.WriteMsg(loginr)
}

// 响应客户端sessionid请求
func handleRestoken(args []interface{}) {
	fmt.Println(reflect.TypeOf(args))
	// 收到的 Hello 消息
	m := args[0].(*msg.ResTokenMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)

	strsign := string(m.Lines[0])
	//log.Release("收到客户端效验信息:%s",strsign)
	encstrsign := hex.EncodeToString(base.Md5_with([]byte(a.UserData().(string))))
	encstrsign = strings.ToUpper(encstrsign)
	encstrsign = hex.EncodeToString(base.Md5_with([]byte(encstrsign)))
	encstrsign = strings.ToUpper(encstrsign)
	if encstrsign != strsign {
		log.Debug("%s被踢下线,原因:sign效验失败", a.RemoteAddr().String())
		a.Destroy()
	} else {
		log.Debug("%s 客户端数据效验成功", a.RemoteAddr().String())
	}
}
