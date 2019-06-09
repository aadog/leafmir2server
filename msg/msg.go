package msg

var Processor = NewMir2Processor()

func init() {
	Processor.Register(&LoginMessage{}, CM_IDPASSWORD)
	Processor.Register(&SelectServerMessage{}, CM_SELECTSERVER)
	Processor.Register(&QuerychrMessage{}, CM_QUERYCHR)
	Processor.Register(&SelchrMessage{}, CM_SELCHR)
	Processor.Register(&GameLoginMessage{}, CM_GAMELOGIN)
	Processor.Register(&LoginnoticeokMessage{}, CM_LOGINNOTICEOK)
	Processor.Register(&QueryBagitemsMessage{}, CM_QUERYBAGITEMS)
	Processor.Register(&ReqTokenMessage{}, SM_REQTOKEN)
	Processor.Register(&ResTokenMessage{}, CM_RESTOKEN)
	Processor.Register(&SelResSessionMessage{}, CM_SELRESSESSION)
}

type LoginMessage struct {
	BaseMsg
}
type SelectServerMessage struct {
	BaseMsg
}
type QuerychrMessage struct {
	BaseMsg
}
type SelchrMessage struct {
	BaseMsg
}
type GameLoginMessage struct {
	BaseMsg
}
type LoginnoticeokMessage struct {
	BaseMsg
}
type QueryBagitemsMessage struct {
	BaseMsg
}
type ReqTokenMessage struct {
	BaseMsg
}
type ResTokenMessage struct {
	BaseMsg
}
type SelResSessionMessage struct {
	BaseMsg
}
