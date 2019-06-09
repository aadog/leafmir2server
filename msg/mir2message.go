package msg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type IMessage interface {
	SetMir2Message(_message *Mir2Message)
}
type BaseMsg struct {
	Mir2Message
}

func (this *BaseMsg) SetMir2Message(_message *Mir2Message) {
	this.Mir2Message = *_message
}

type Mir2Message struct {
	Ident      int32  //2
	Recog      int32  //4
	Param      int32  //4
	Tag        int32  //4
	Series     int32  //4
	NSessionID int32  //4
	NToken     int32  //4
	CRC        uint32 //4
	//Align int16          //delphi松散模式有自动32位需要补码两位
	Lines []string
}

func (this *Mir2Message) EncodeBytes() ([]byte, error) {
	it := this
	wbuf := &bytes.Buffer{}
	err := binary.Write(wbuf, binary.LittleEndian, it.Ident)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.Recog)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.Param)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.Tag)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.Series)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.NSessionID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.NToken)
	if err != nil {
		return nil, err
	}
	err = binary.Write(wbuf, binary.LittleEndian, it.CRC)
	if err != nil {
		return nil, err
	}
	//err=binary.Write(wbuf,binary.LittleEndian,it.Align)
	//if err!=nil{
	//	return nil,err
	//}
	_, err = wbuf.WriteString(it.Stringlines())
	if err != nil {
		return nil, err
	}

	return wbuf.Bytes(), nil
}
func (this *Mir2Message) Stringlines() string {
	r := ""
	for i, s := range this.Lines {
		if i != 0 {
			r = fmt.Sprintf("%s/", r)
		}
		r = fmt.Sprintf("%s%s", r, s)
	}
	return r
}
func (this *Mir2Message) Add_with_line(_line string) {
	this.Lines = append(this.Lines, _line)
}
func NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(_msg int32, _recog, _param, _tag, _series, _NSessionID, _NToken int32, _CRC uint32, _lines ...string) *Mir2Message {
	lines := []string{}
	for _, line := range _lines {
		lines = append(lines, line)
	}
	m := &Mir2Message{_msg, _recog, _param, _tag, _series, _NSessionID, _NToken, _CRC, lines}
	return m
}
func DecodeMir2Message_with_bytes(_bytes []byte) (*Mir2Message, error) {
	var message Mir2Message
	dedatareader := bufio.NewReader(bytes.NewReader(_bytes))
	err := binary.Read(dedatareader, binary.LittleEndian, &message.Ident)
	if err != nil {
		return nil, err
	}
	message.Ident = int32(int16(message.Ident))
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Recog)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Param)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Tag)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Series)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.NSessionID)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.NToken)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.CRC)
	if err != nil {
		return nil, err
	}
	//err=binary.Read(dedatareader,binary.LittleEndian,&message.Align)
	//if err!=nil{
	//	return nil,err
	//}
	messagebody := string(_bytes[32:len(_bytes)])
	spmessagebody := strings.Split(messagebody, "/")
	for _, body := range spmessagebody {
		message.Lines = append(message.Lines, body)
	}
	return &message, nil
}
func DecodeMir2Message_with_Txtbytes(_bytes []byte) (*Mir2Message, error) {
	var message Mir2Message
	dedatareader := bufio.NewReader(bytes.NewReader(_bytes))
	err := binary.Read(dedatareader, binary.LittleEndian, &message.Ident)
	if err != nil {
		return nil, err
	}
	message.Ident = int32(int16(message.Ident))
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Recog)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Param)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Tag)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.Series)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.NSessionID)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.NToken)
	if err != nil {
		return nil, err
	}
	err = binary.Read(dedatareader, binary.LittleEndian, &message.CRC)
	if err != nil {
		return nil, err
	}
	//err=binary.Read(dedatareader,binary.LittleEndian,&message.Align)
	//if err!=nil{
	//	return nil,err
	//}

	messagebody := string(_bytes[32:len(_bytes)])
	message.Lines = append(message.Lines, messagebody)
	return &message, nil
}
