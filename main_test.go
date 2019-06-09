package main

import (
	"errors"
	"fmt"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"leafmir2server/msg"
	"strconv"
	"strings"
	"testing"
)

func Test_enc(t *testing.T) {
	enc := []byte("#kJ]sA[BasUB>v&XldGJyAvkMriRtRfDax<xVTI]Ibe]=YKAV0lLQUYoEXHY5sTipFOR5M2t0jzLks95FbaxRaUEItG8dGOWRKdIanEKQ0/iejSO1qhunNgOcoWvkNzUd/bf/3hG5BA==!")
	debt, err := DecodeMsg(enc)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(debt)
}

func DecodeMsg(bt []byte) (*msg.Mir2Message, error) {
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
	return rmsg, nil
}

func DecodeString(bt []byte) ([]byte, error) {
	if strings.HasPrefix(string(bt), "#") == false {
		return nil, errors.New("无法识别包头")
	}
	if strings.HasSuffix(string(bt), "!") == false {
		return nil, errors.New("无法识别包尾")
	}
	nseq, err := strconv.Atoi(string(bt[1]))
	if err != nil {
		return nil, err
	}
	_ = nseq

	encdata := base.DecodeString_EDCode(bt[2 : len(bt)-1])
	if encdata[0] == '*' && encdata[1] == '*' {
		m := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.CM_GAMELOGIN, 0, 0, 0, 0, 0, 0, 0)
		m.Add_with_line(string(encdata[2:]))
		encbt, err := m.EncodeBytes()
		if err != nil {
			return nil, err
		}
		return encbt, nil
	} else {
		m := msg.NewMir2Message_with_msg_recog_param_tag_series_nsessionid_ntoken_ctc_lines(msg.CM_GAMESESSION, 0, 0, 0, 0, 0, 0, 0)
		m.Add_with_line(string(encdata))
		encbt, err := m.EncodeBytes()
		if err != nil {
			return nil, err
		}
		return encbt, nil
	}
}
