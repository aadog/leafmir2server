package msg
var Processor=NewMir2Processor()

func init() {
	Processor.Register(&LoginMessage{},CM_IDPASSWORD)
	Processor.Register(&SelectServerMessage{},CM_SELECTSERVER)
	Processor.Register(&QuerychrMessage{},CM_QUERYCHR)
	Processor.Register(&SelchrMessage{},CM_SELCHR)
	Processor.Register(&GameLoginMessage{},CM_GAMELOGIN)
	Processor.Register(&LoginnoticeokMessage{},CM_LOGINNOTICEOK)
	Processor.Register(&QueryBagitemsMessage{},CM_QUERYBAGITEMS)
}


type LoginMessage struct {
	Mir2Message
}
type SelectServerMessage struct{
	Mir2Message
}
type QuerychrMessage struct{
	Mir2Message
}
type SelchrMessage struct{
	Mir2Message
}
type GameLoginMessage struct {
	Mir2Message
}
type LoginnoticeokMessage struct{
	Mir2Message
}
type QueryBagitemsMessage struct{
	Mir2Message
}