package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/name5566/leaf/log"
	"leafmir2server/base"
	"leafmir2server/msg"

	"strconv"
	"strings"
	"testing"
)

func Test_decmsg(t *testing.T) {

	enc := []byte("#sPkL~gmQpTQIjgU[fnK@WokTwr}}dU>ijK}t&i>GucO=NBvUcRvORxu?NBu?NC@BNhvNNCq?NBu?NBvBNBu?NBu?NBu?NBu?Nhu?NBu?NBu?NC@JPSDNbS&BXRvENBu?NCDjNBu?NEC?NBu?QRu?NBu{bRu?NBvrNBu?NCW?NBu?OUS?NBu?WRu?NBvHNBu?NB&zNBu?NEy?NBu?Phu?NBv>XRu?NBu}NBu?NC]?NBu?PS??NBvBaRu?NB&INBu?NCbzNBu?OSq?NBu?Rhu?NBvAbRu?NBvRNBu?NCO?NBu?OlS?NBu?Nxu?NBvMNBu?NCHzNBu?O[C?NBu?Phu?NBv>XRu?NBvrNhvNOxvBOUC?NBvBNhvNNB]{Q[vhcSjfNBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBu?NBvNOx&JXRu?NC@BNhu?NBv>Zhu{NBu?XRvzNBu?Rxu@!")
	debt, err := DecodesrvMsg(enc)
	if err != nil {
		log.Fatal(err.Error())
	}
	//fmt.Println(base.ConvertByte2String([]byte(debt.Lines[0]),base.GBK))
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
func Test_C(t *testing.T) {
	key := []byte("964CF629-E9EC-47B8-AAFE-39172793BEB4-D89A1212-DB56-4858-9ADE-CF19A782E5A3-0B92A08-3578-49B7-BDF6-21B73AC5D9E5")
	//tr,err:=base.EncodeStream_uEDCode([]byte("test111111"),key)
	//if err!=nil{
	//	log.Fatal(err.Error())
	//}
	//log.Debug("加密后的结果:%v",hex.EncodeToString(tr))
	tr, err := hex.DecodeString("5a9b10410adf342b5f5ed58f711639943acb")
	if err != nil {
		log.Fatal(err.Error())
	}

	dtr, err := base.DecodeStream_uEDCode(tr, key)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Debug("解密后的结果:%v", string(dtr))

}
