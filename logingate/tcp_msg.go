package logingate

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/name5566/leaf/network"
	"leafmir2server/base"
	"leafmir2server/msg"
	"strconv"
	"strings"
)

// --------------
// | len | data |
// --------------
type MsgParser struct {
}

func (p *MsgParser) DecodeAesMessage_with_bytes(_in []byte) (*msg.Mir2Message, error) {
	if len(_in) != 44 {
		return nil, errors.New("长度必须为44")
	}
	return nil, errors.New("实现不完整")
}
func (p *MsgParser) Read(conn *network.TCPConn) ([]byte, error) {
	nresseq, _ := ca.Get(base.Reskey(conn.RemoteAddr().String()))
	defer func() {
		ca.SetDefault(base.Reskey(conn.RemoteAddr().String()), nresseq.(int)+1)
	}()

	rd := bufio.NewReader(conn)
	bt, err := rd.ReadBytes('!')
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(string(bt), "#") == false {
		return nil, errors.New("无法识别包头")
	}
	if strings.HasSuffix(string(bt), "!") == false {
		return nil, errors.New("无法识别包尾")
	}
	if len(bt) < 3 {
		return nil, errors.New("最短长度为3")
	}
	if nresseq == 0 { //此时还未认证，开始解密认证信息,的token效验

		nseq, err := strconv.Atoi(string(bt[1]))
		if err != nil {
			return nil, err
		}
		if nseq != 0 {
			return nil, errors.New("认证失败，req错误")
		}
		encdata := bt[2 : len(bt)-1]

		decdata := base.DecodeString_EDCode(encdata)
		decdata = base.DecodeString_uEDCode(decdata, []byte("^(YEFGH%tgb$r64ib5"))
		tokenid := decdata
		m := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.CM_RESTOKEN, 0, 0, 0, 0, 0, 0, 0)
		m.Add_with_line(string(tokenid))
		encbt, err := m.EncodeBytes()
		if err != nil {
			return nil, err
		}
		return encbt, nil
	} else { //已经认证
		nseq, err := strconv.Atoi(string(bt[1]))
		if err != nil {
			return nil, err
		}
		_ = nseq
		encdata := bt[2 : len(bt)-1]

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
	return bt, nil
}
func (p *MsgParser) Write(conn *network.TCPConn, args ...[]byte) error {
	nreqseq, _ := ca.Get(base.Reqkey(conn.RemoteAddr().String()))
	defer func() {
		ca.SetDefault(base.Reqkey(conn.RemoteAddr().String()), nreqseq.(int)+1)
	}()
	var bt []byte
	for _, it := range args {
		bt = append(bt, it...)
	} //合并所有byte
	var wbuf bytes.Buffer
	if nreqseq == 0 { //此时还未认证，开始解密认证信息,的token效验

		message, err := msg.DecodeMir2Message_with_Txtbytes(bt)
		if err != nil {
			return err
		}

		wbuf.WriteString("#")
		wbuf.Write(base.EncodeString_EDCode([]byte(message.Lines[0])))
		wbuf.WriteString("!")
		//ti.ReqSession = string(base.DecodeString_uEDCode([]byte(message.Lines[0]), []byte("#$Ggy%(*^45fghj@@#sqw[]KHG%^&UHBR#$ty")))
		conn.Write(wbuf.Bytes())
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
	return nil
}
func (p *MsgParser) Conn(conn *network.TCPConn) {
	ca.SetDefault(base.Reqkey(conn.RemoteAddr().String()), 0)
	ca.SetDefault(base.Reskey(conn.RemoteAddr().String()), 0)
}
func (p *MsgParser) Close(conn *network.TCPConn) {
	ca.Delete(base.Reqkey(conn.RemoteAddr().String()))
	ca.Delete(base.Reskey(conn.RemoteAddr().String()))
}
