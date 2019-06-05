package msg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type Mir2Processor struct{
	msgInfo      map[int16]*MsgInfo
	msgID        map[reflect.Type]int16
}

// must goroutine safe
func (this* Mir2Processor) Route(msg interface{}, userData interface{}) error{
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
func (this* Mir2Processor) Unmarshal(data []byte) (interface{}, error){
	if strings.HasPrefix(string(data),"#")==false{
		return nil,errors.New("无法识别包头")
	}
	if strings.HasSuffix(string(data),"!")==false{
		return nil,errors.New("无法识别包尾")
	}
	if len(data)<3{
		return nil,errors.New("最短长度为3")
	}
	nseq,err:=strconv.Atoi(string(data[1]))
	if err!=nil{
		return nil,err
	}
	_=nseq   //获取到seq
	//fmt.Println("seq:",nseq)
	encdata:=data[2:len(data)-1]
	dedata:=decode6BitBytes(encdata)

	if strings.HasPrefix(string(dedata),"**"){  //处理特殊登录指令
	    sdedata:=string(dedata[2:])
		spsdedata:=strings.Split(sdedata,"/")
		message:=Mir2Message{}
		for _,it:=range spsdedata{
			message.Lines=append(message.Lines,base.ConvertByte2String([]byte(it),base.GB18030))
		}
		message.Ident=int16(CM_GAMELOGIN)
		// msg
		i := this.msgInfo[message.Ident]
		if i==nil{
			return nil,errors.New(fmt.Sprintf("message id %v not registered", message.Ident))
		}
		it:=reflect.NewAt(i.msgType.Elem(),unsafe.Pointer(&message))
		return it.Interface(), nil
	}
	dedatareader:=bufio.NewReader(bytes.NewReader(dedata))

	message:=Mir2Message{}
	err=binary.Read(dedatareader,binary.LittleEndian,&message.Recog)
	if err!=nil{
		return nil,err
	}
	err=binary.Read(dedatareader,binary.LittleEndian,&message.Ident)
	if err!=nil{
		return nil,err
	}
	err=binary.Read(dedatareader,binary.LittleEndian,&message.Param)
	if err!=nil{
		return nil,err
	}
	err=binary.Read(dedatareader,binary.LittleEndian,&message.Tag)
	if err!=nil{
		return nil,err
	}
	err=binary.Read(dedatareader,binary.LittleEndian,&message.Series)
	if err!=nil{
		return nil,err
	}
	messagebody:=string(dedata[12:len(dedata)])
	spmessagebody:=strings.Split(messagebody,"/")
	for _,body:=range spmessagebody{
		message.Lines=append(message.Lines,base.ConvertByte2String([]byte(body),base.GB18030))
	}

	// msg
	i := this.msgInfo[message.Ident]
	if i==nil{
		return nil,errors.New(fmt.Sprintf("message id %v not registered", message.Ident))
	}
	it:=reflect.NewAt(i.msgType.Elem(),unsafe.Pointer(&message))
	return it.Interface(), nil
}
// must goroutine safe
func (this* Mir2Processor) Marshal(msg interface{}) ([][]byte, error){
	it:=msg.(*Mir2Message)
	wbuf:=&bytes.Buffer{}
	err:=binary.Write(wbuf,binary.LittleEndian,it.Recog)
	if err!=nil{
		return nil,err
	}
	err=binary.Write(wbuf,binary.LittleEndian,it.Ident)
	if err!=nil{
		return nil,err
	}
	err=binary.Write(wbuf,binary.LittleEndian,it.Param)
	if err!=nil{
		return nil,err
	}
	err=binary.Write(wbuf,binary.LittleEndian,it.Tag)
	if err!=nil{
		return nil,err
	}
	err=binary.Write(wbuf,binary.LittleEndian,it.Series)
	if err!=nil{
		return nil,err
	}
	_,err=wbuf.WriteString(it.Stringlines())
	if err!=nil{
		return nil,err
	}
	encdata:=encoder6BitBuf(wbuf.Bytes())


	rbuf:=&bytes.Buffer{}
	_,err=rbuf.WriteString("#")
	if err!=nil{
		return nil,err
	}
	_,err=rbuf.Write(encdata)
	if err!=nil{
		return nil,err
	}
	_,err=rbuf.WriteString("!")
	if err!=nil{
		return nil,err
	}
	return [][]byte{rbuf.Bytes()},nil
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (this* Mir2Processor) SetRouter(msg interface{}, msgRouter *chanrpc.Server) {
	msgType := reflect.TypeOf(msg)
	id, ok := this.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	this.msgInfo[id].msgRouter = msgRouter
}
// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (this* Mir2Processor) Register(msg interface{},_type int16) int16 {
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
	this.msgInfo[_type] =i

	this.msgID[msgType] = _type
	return _type
}

func NewMir2Processor() *Mir2Processor {
	p := new(Mir2Processor)
	p.msgID = make(map[reflect.Type]int16)
	p.msgInfo=make(map[int16]*MsgInfo)
	return p
}


type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}
type MsgHandler func([]interface{})



func decode6BitBytes(src []byte) []byte {
	var decode6BitMask = [...]byte{0xfc, 0xf8, 0xf0, 0xe0, 0xc0}
	var size = len(src)
	var dest = make([]byte, size*3/4)
	var destPos = 0
	var bitPos uint = 2
	var madeBit uint = 0

	var ch byte = 0
	var chCode byte = 0
	var tmp byte = 0

	for i := 0; i < size; i++ {
		if (src[i] - 0x3c) >= 0 {
			ch = byte(src[i] - 0x3c)
		} else {
			destPos = 0
			break
		}

		if destPos >= len(dest) {
			break
		}

		if madeBit+6 >= 8 {
			chCode = byte(tmp | ((ch & 0x3f) >> uint(6-bitPos)))

			dest[destPos] = chCode
			destPos += 1

			madeBit = 0
			if bitPos < 6 {
				bitPos += 2
			} else {
				bitPos = 2
				continue
			}
		}

		tmp = (byte)((ch << bitPos) & decode6BitMask[bitPos-2])

		madeBit += 8 - bitPos
	}

	return dest
}
func encoder6BitBuf(src []byte) []byte {
	var size = len(src)
	var destLen = (size/3)*4 + 10
	var dest = make([]byte, destLen)
	var destPos = 0
	var resetCount = 0

	var chMade, chRest byte = 0, 0

	for i := 0; i < size; i++ {
		if destPos >= destLen {
			break
		}

		chMade = (byte)((chRest | ((src[i] & 0xff) >> uint(2+resetCount))) & 0x3f)
		chRest = (byte)((((src[i] & 0xff) << uint(8-(2+resetCount))) >> uint(2)) & 0x3f)

		resetCount += 2
		if resetCount < 6 {
			dest[destPos] = (byte)(chMade + 0x3c)
			destPos += 1
		} else {
			if destPos < destLen-1 {
				dest[destPos] = (byte)(chMade + 0x3c)
				destPos += 1
				dest[destPos] = (byte)(chRest + 0x3c)
				destPos += 1
			} else {
				dest[destPos] = (byte)(chMade + 0x3c)
				destPos += 1
			}

			resetCount = 0
			chRest = 0
		}
	}
	if resetCount > 0 {
		dest[destPos] = (byte)(chRest + 0x3c)
		destPos += 1
	}

	dest[destPos] = 0
	return dest[:destPos]

}