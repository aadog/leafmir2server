package base

import (
	"errors"
	"github.com/name5566/leaf/log"
	. "github.com/ying32/govcl/vcl"
	"syscall"
	"unsafe"
)

var (
	c_project              = syscall.NewLazyDLL("./d91m2.dll")
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
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	keystm := NewMemoryStreamFromBytes(_key)
	defer keystm.Free()
	instm.SetPosition(0)
	keystm.SetPosition(0)
	cEncodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(keystm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func DecodeString_uEDCode(_in []byte, _key []byte) []byte {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	keystm := NewMemoryStreamFromBytes(_key)
	defer keystm.Free()
	instm.SetPosition(0)
	keystm.SetPosition(0)
	cDecodeString_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(keystm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func EncodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	keystm := NewMemoryStreamFromBytes(_key)
	defer keystm.Free()
	instm.SetPosition(0)
	keystm.SetPosition(0)
	cEncodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(keystm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r, nil
}
func DecodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	keystm := NewMemoryStreamFromBytes(_key)
	defer keystm.Free()
	instm.SetPosition(0)
	keystm.SetPosition(0)

	cDecodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(keystm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r, nil
}
func EncodeString_EDCode(_in []byte) []byte {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cEncodeString_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func DecodeString_EDCode(_in []byte) []byte {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cDecodeString_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func Base64DecodeEx_EDcode(_in []byte, _len int) []byte {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cBase64DecodeEx_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(_len),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func SetPassWord_EDcode(_in []byte) {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	instm.SetPosition(0)
	cSetPassWord_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
	)
}
func DecryptAES_EDcode(_in []byte) []byte {
	if len(_in) != 16 {
		panic(errors.New("数组长度必须是16位"))
	}
	instm := NewMemoryStreamFromBytes(_in[0:16])
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cDecryptAES_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func EncryptAES_EDcode(_in []byte) []byte {
	if len(_in) != 16 {
		panic(errors.New("数组长度必须是16位"))
	}
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cEncryptAES_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func Base64Encode_EDcode(_in []byte, _len int) []byte {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	instm.SetPosition(0)
	cBase64Encode_EDcode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(_len),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r
}
func Crc32(_in []byte) uint32 {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	instm.SetPosition(0)
	r, _, _ := cCrc32.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
	)
	return uint32(r)
}

func init() {
	err := c_project.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	//初始化后直接设置密码
	SetPassWord_EDcode([]byte("glaciersoftware@126.com"))
}
