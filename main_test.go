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

func Test_decsnappy(t *testing.T) {

}
func Test_decmsg(t *testing.T) {
	enc := []byte("#2IgnfKRGxKRGxJ~DxXEXxWUPkJxO~JxiwKBCzKRGvKg?uJ{cHgMYapBmeKh<vLw]uMQ]uLg?yOyXCJSW&MSWrPhjBPAy[PBT>JyHVRBHvTxTFT~PWMRjraCC@NQ@AJzbnZkPtb[KeLw>RXVH~YULjGD>fW~qeKQ<mTkTwa~jtZg<~JhCqGCHzYUviGBa~KBCqGBWyJUHnbA>DXEjyYU@sIQ?oIgmoIgmoIgmo!")
	debt, err := DecodecliMsg(enc)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(debt)
}
func Test_decstr(t *testing.T) {
	enc := []byte("#2IgnfKRGxKRGxJ~DxXEXxWUPkJxO~JxiwKBCzKRGvKg?uJ{cHgMYapBmeKh<vLw]uMQ]uLg?yOyXCJSW&MSWrPhjBPAy[PBT>JyHVRBHvTxTFT~PWMRjraCC@NQ@AJzbnZkPtb[KeLw>RXVH~YULjGD>fW~qeKQ<mTkTwa~jtZg<~JhCqGCHzYUviGBa~KBCqGBWyJUHnbA>DXEjyYU@sIQ?oIgmoIgmoIgmo!")
	debt, err := DecodecliString(enc)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(debt))
}

func DecodecliMsg(bt []byte) (*msg.Mir2Message, error) {
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
func DecodesrvMsg(bt []byte) (*msg.Mir2Message, error) {
	if strings.HasPrefix(string(bt), "#") == false {
		return nil, errors.New("无法识别包头")
	}
	if strings.HasSuffix(string(bt), "!") == false {
		return nil, errors.New("无法识别包尾")
	}
	if len(bt) < 44+2 {
		return nil, errors.New("最短长度为45")
	}
	encdata := bt[1 : len(bt)-1]

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
func DecodecliString(bt []byte) ([]byte, error) {
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
		return []byte(string(encdata[2:])), nil
	} else {

		return []byte(string(encdata)), nil
	}
}
