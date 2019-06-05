package internal

import (
	"bufio"
	"github.com/name5566/leaf/network"
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


// goroutine safe
func (p *MsgParser) Read(conn *network.TCPConn) ([]byte, error) {
	rd:=bufio.NewReader(conn)
	//l,err:=rd.ReadByte()
	//if err!=nil{
	//	return nil,err
	//}
	//if l!='#'{
	//	return nil,errors.New(fmt.Sprintf("无法识别头部字符:%c",l))
	//}
	bt,err:=rd.ReadBytes('!')
	if err!=nil{
		return nil,err
	}
	return bt, nil
}

// goroutine safe
func (p *MsgParser) Write(conn *network.TCPConn, args ...[]byte) error {
	var sendbt []byte
	for _,it:=range args{
		sendbt=append(sendbt,it...)
	}
	conn.Write(sendbt)
	return nil
}

