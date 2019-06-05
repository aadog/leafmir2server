package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leafmir2server/msg"
	"reflect"
	"strconv"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.LoginMessage{},handleLogin)
	handleMsg(&msg.SelectServerMessage{},handleSelectserver)
}

func handleSelectserver(args []interface{}){
	// 收到的 Hello 消息
	m := args[0].(*msg.SelectServerMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)
	servername:= m.Lines[0]
	log.Debug("选择的服务器:%s",servername)
	Cert:=10
	loginserverip:="127.0.0.1"
	loginserverport:="7004"
	loginservercert:=strconv.Itoa(Cert)
    selectserverr:=msg.NewMir2Message_with_msg_soul_wparam_atag_nseries(msg.SM_SELECTSERVER_OK,int32(Cert),0,0,0,
    	loginserverip,
    	loginserverport,
    	loginservercert,
    	)
    a.WriteMsg(selectserverr)
}

func handleLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.LoginMessage)
	// 消息的发送者
	a := args[1].(gate.Agent)


	// 输出收到的消息的内容
	log.Debug("开始登录 客户端版本:%v 账号:%v,密码:%v", m.Recog,m.Lines[0],m.Lines[1])

	//// 给发送者回应一个 Hello 消息
	loginr:=msg.NewMir2Message_with_msg_soul_wparam_atag_nseries(msg.SM_PASSOK_SELECTSERVER,0,0,0,1,"清风","1")
	a.WriteMsg(loginr)
}


