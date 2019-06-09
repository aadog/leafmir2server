package login

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	uuid "github.com/satori/go.uuid"
	"leafmir2server/base"
	"leafmir2server/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}
func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("loginserver,来自客户端连接:%s", a.RemoteAddr())
	clitoken := uuid.NewV4().String()
	clitoken = "{3BCA3DFF-966D-48FA-A891-3E43D247EA99}"
	enctoken := base.EncodeString_uEDCode([]byte(clitoken), []byte("#$Ggy%(*^45fghj@@#sqw[]KHG%^&UHBR#$ty"))
	smsg := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.SM_REQTOKEN, 0, 0, 0, 0, 0, 0, 0)
	smsg.Add_with_line(string(enctoken))
	a.SetUserData(clitoken)
	a.WriteMsg(smsg)
}
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	fmt.Println("断开连接")
}
