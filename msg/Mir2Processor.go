package msg

import (
	"errors"
	"fmt"
	"github.com/a97077088/leaf/chanrpc"
	"github.com/a97077088/leaf/log"
	"math"
	"reflect"
)

type Mir2Processor struct {
	msgInfo map[int16]*MsgInfo
	msgID   map[reflect.Type]int16
}

// must goroutine safe
func (this *Mir2Processor) Route(msg interface{}, userData interface{}) error {
	msgType := reflect.ValueOf(msg).Type()
	id, ok := this.msgID[msgType]
	if !ok {
		return fmt.Errorf("message %s not registered", msgType)
	}
	i := this.msgInfo[id]
	if i.msgHandler != nil {
		i.msgHandler([]interface{}{msg, userData})
	}

	if i.msgRouter != nil {
		i.msgRouter.Go(msgType, msg, userData)
	}
	return nil
}

// must goroutine safe
func (this *Mir2Processor) Unmarshal(data []byte) (interface{}, error) {
	message, err := DecodeMir2Message_with_bytes(data)
	if err != nil {
		return nil, err
	}
	// msg
	i := this.msgInfo[int16(message.Ident)]
	if i == nil {
		return nil, errors.New(fmt.Sprintf("message id %v not registered", message.Ident))
	}
	it := reflect.New(i.msgType.Elem())
	pit := it.Interface().(IMessage)
	pit.SetMir2Message(message)
	return it.Interface(), nil
}

// must goroutine safe
func (this *Mir2Processor) Marshal(msg interface{}) ([][]byte, error) {
	it := msg.(*Mir2Message)
	msgbt, err := it.EncodeBytes()
	if err != nil {
		return nil, err
	}
	return [][]byte{msgbt}, nil
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (this *Mir2Processor) SetRouter(msg interface{}, msgRouter *chanrpc.Server) {
	msgType := reflect.TypeOf(msg)
	id, ok := this.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	this.msgInfo[id].msgRouter = msgRouter
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (this *Mir2Processor) Register(msg interface{}, _type int16) int16 {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("json message pointer required")
	}
	if _, ok := this.msgID[msgType]; ok {
		log.Fatal("message %s is already registered", msgType)
	}
	if len(this.msgInfo) >= math.MaxUint16 {
		log.Fatal("too many protobuf messages (max = %v)", math.MaxUint16)
	}

	i := new(MsgInfo)
	i.msgType = msgType
	this.msgInfo[_type] = i

	this.msgID[msgType] = _type
	return _type
}

func NewMir2Processor() *Mir2Processor {
	p := new(Mir2Processor)
	p.msgID = make(map[reflect.Type]int16)
	p.msgInfo = make(map[int16]*MsgInfo)
	return p
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}
type MsgHandler func([]interface{})
