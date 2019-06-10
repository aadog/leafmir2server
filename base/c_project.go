package base

import (
	"errors"
	"github.com/name5566/leaf/log"
	"syscall"
	"unsafe"
)

var (
	c_project              = syscall.NewLazyDLL("./project.dll")
	cEncodeStream_uEDCode  = c_project.NewProc("cEncodeStream_uEDCode")
	cDecodeStream_uEDCode  = c_project.NewProc("cDecodeStream_uEDCode")
	cEncodeString_uEDCode  = c_project.NewProc("cEncodeString_uEDCode")
	cDecodeString_uEDCode  = c_project.NewProc("cDecodeString_uEDCode")
	cEncodeString_EDcode   = c_project.NewProc("cEncodeString_EDcode")
	cDecodeString_EDcode   = c_project.NewProc("cDecodeString_EDcode")
	cBase64DecodeEx_EDcode = c_project.NewProc("cBase64DecodeEx_EDcode")
	cSetPassWord_EDcode    = c_project.NewProc("cSetPassWord_EDcode")
	cDecryptAES_EDcode     = c_project.NewProc("cDecryptAES_EDcode")
	cEncryptAES_EDcode     = c_project.NewProc("cEncryptAES_EDcode")
	cBase64Encode_EDcode   = c_project.NewProc("cBase64Encode_EDcode")
	cCrc32                 = c_project.NewProc("cCrc32")
)

func EncodeString_uEDCode(_in []byte, _key []byte) []byte {
	_in = append(_in, 0x00)
	_key = append(_key, 0x00)
	var rlen int
	cEncodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&_key[0])),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
	)

	r := make([]byte, rlen)
	cEncodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&_key[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
	)
	return r
}
func DecodeString_uEDCode(_in []byte, _key []byte) []byte {
	_in = append(_in, 0x00)
	_key = append(_key, 0x00)
	var rlen int
	cDecodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&_key[0])),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
	)
	r := make([]byte, rlen)
	cDecodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&_key[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
	)
	return r
}
func EncodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	_key = append(_key, 0x00)
	var rlen int
	cEncodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	r := make([]byte, rlen)
	cEncodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	return r, nil
}
func DecodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	_key = append(_key, 0x00)
	var rlen int
	cDecodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	r := make([]byte, rlen)
	cDecodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	return r, nil
}
func EncodeString_EDCode(_in []byte) []byte {
	_in = append(_in, 0x00)
	var rlen int
	cEncodeString_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
	)

	r := make([]byte, rlen)
	cEncodeString_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
	)
	return r
}
func DecodeString_EDCode(_in []byte) []byte {
	_in = append(_in, 0x00)
	var rlen int
	cDecodeString_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
	)
	r := make([]byte, rlen)
	cDecodeString_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
	)
	return r
}
func Base64DecodeEx_EDcode(_in []byte, _len int) []byte {
	_in = append(_in, 0x00)
	r := make([]byte, _len)
	cBase64DecodeEx_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(_len),
	)
	return r
}
func SetPassWord_EDcode(_in []byte) {
	cSetPassWord_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
	)
}
func DecryptAES_EDcode(_in []byte) []byte {
	if len(_in) != 16 {
		panic(errors.New("长度必须为16位"))
	}
	r := make([]byte, 16)
	cDecryptAES_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
	)
	return r
}
func EncryptAES_EDcode(_in []byte) []byte {
	if len(_in) != 16 {
		panic(errors.New("长度必须为16位"))
	}
	r := make([]byte, 16)
	cEncryptAES_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
	)
	return r
}
func Base64Encode_EDcode(_in []byte, _len int) []byte {
	_in = append(_in, 0x00)
	r := make([]byte, _len)
	cBase64Encode_EDcode.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(_len),
	)
	return r
}
func Crc32(_in []byte) int {
	r, _, _ := cCrc32.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
	)
	return int(r)
}

func init() {
	err := c_project.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	//初始化后直接设置密码
	SetPassWord_EDcode([]byte("glaciersoftware@126.com"))
}
