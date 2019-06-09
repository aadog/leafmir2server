package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/name5566/leaf/network"
	"leafmir2server/base"
	"leafmir2server/login"
	"leafmir2server/msg"
	"strconv"
	"strings"
)

// --------------
// | len | data |
// --------------
type MsgParser struct {
}

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	return p
}

func (p *MsgParser) DecodeAesMessage_with_bytes(_in []byte) (*msg.Mir2Message, error) {
	if len(_in) != 44 {
		return nil, errors.New("长度必须为44")
	}
	return nil, errors.New("实现不完整")
}

// goroutine safe
func (p *MsgParser) Read(conn *network.TCPConn) ([]byte, error) {

	rd := bufio.NewReader(conn)
	bt, err := rd.ReadBytes('!')
	if err != nil {
		return nil, err
	}
	fmt.Println(bt)

	if strings.HasPrefix(string(bt), "#") == false {
		return nil, errors.New("无法识别包头")
	}
	if strings.HasSuffix(string(bt), "!") == false {
		return nil, errors.New("无法识别包尾")
	}
	if len(bt) < 44+2 {
		return nil, errors.New("最短长度为45")
	}
	nseq, err := strconv.Atoi(string(bt[1]))
	if err != nil {
		return nil, err
	}
	_ = nseq
	encdata := bt[2 : len(bt)-1]
	fmt.Println(encdata)

	//解密头部字节
	dechd := base.Base64DecodeEx_EDcode([]byte(string(encdata[:44])), 32) //传递一个拷贝
	dechd1 := base.DecryptAES_EDcode(dechd[:16])
	dechd2 := base.DecryptAES_EDcode(dechd[16:])

	decd2 := base.DecodeString_EDCode([]byte(string(encdata[44:]))) //这里转string可以copy一次
	decbt := append(dechd1, dechd2...)
	decbt = append(decbt, decd2...)
	rmsg, err := msg.DecodeMir2Message_with_bytes(decbt)
	if err != nil {
		return nil, err
	}
	return rmsg.EncodeBytes()
}

// goroutine safe
func (p *MsgParser) Write(conn *network.TCPConn, args ...[]byte) error {
	var bt []byte
	for _, it := range args {
		bt = append(bt, it...)
	} //合并所有byte
	var wbuf bytes.Buffer

	ti := login.GetTcpInfo(conn.RemoteAddr().String())
	if ti.ReqSession == "" { //此时还未认证，开始解密认证信息,的token效验
		message, err := msg.DecodeMir2Message_with_Txtbytes(bt)
		if err != nil {
			return err
		}
		switch message.Ident {
		case msg.SM_REQTOKEN:
			wbuf.WriteString("#")
			wbuf.Write(base.EncodeString_EDCode([]byte(message.Lines[0])))
			wbuf.WriteString("!")
			ti.ReqSession = string(base.DecodeString_uEDCode([]byte(message.Lines[0]), []byte("#$Ggy%(*^45fghj@@#sqw[]KHG%^&UHBR#$ty")))
			conn.Write(wbuf.Bytes())
		}
	} else { //认证完成了
		message, err := msg.DecodeMir2Message_with_bytes(bt)
		if err != nil {
			return err
		}
		var encbuf bytes.Buffer

		//解密头部字节
		decchd := bt[:32] //传递一个拷贝
		enchd1 := base.EncryptAES_EDcode(decchd[:16])
		enchd2 := base.EncryptAES_EDcode(decchd[16:])
		decchd = append(enchd1, enchd2...)

		enchd := base.Base64Encode_EDcode(decchd[:32], 44)
		encbuf.Write(enchd)
		if message.Stringlines() != "" {
			sbuf := base.EncodeString_EDCode([]byte(message.Stringlines()))
			encbuf.Write(sbuf)
		}
		wbuf.WriteString("#")
		wbuf.Write(encbuf.Bytes())
		wbuf.WriteString("!")
		conn.Write(wbuf.Bytes())

	}

	////解密完成
	//ti:=login.GetTcpInfo("aaa")
	//if ti==nil{   //此时还未认证，开始解密认证信息,的token效验
	//   sencdata:=base.EncodeString_EDCode(decdata)
	//   var sencbyte bytes.Buffer
	//   sencbyte.WriteString("#")
	//   sencbyte.Write(sencdata)
	//   sencbyte.WriteString("!")
	//   conn.Write(sencbyte.Bytes())
	//   fmt.Println("消息发送成功")
	//}else{  //发送的是认证过后的包
	//
	//}
	//conn.Write(sendbt)
	return nil
}
