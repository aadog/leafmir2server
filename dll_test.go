package main

import (
	"encoding/hex"
	"fmt"
	. "github.com/ying32/govcl/vcl"
	"syscall"
	"testing"
	"unsafe"
)

var (
	c_project             = syscall.NewLazyDLL("./d91m2.dll")
	cDecodeStream_uEDCode = c_project.NewProc("cDecodeStream_uEDCode")
	cEncodeStream_uEDCode = c_project.NewProc("cEncodeStream_uEDCode")
)

func EncodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	cEncodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r, nil
}
func DecodeStream_uEDCode(_in []byte, _key []byte) ([]byte, error) {
	instm := NewMemoryStreamFromBytes(_in)
	defer instm.Free()
	outstm := NewMemoryStream()
	defer outstm.Free()
	cDecodeStream_uEDCode.Call(
		uintptr(unsafe.Pointer(instm.Instance())),
		uintptr(unsafe.Pointer(outstm.Instance())),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	_, r := outstm.Read(int32(outstm.Size()))
	return r, nil
}

func Test_x(t *testing.T) {
	//_path := "C:\\MirServer\\Mir200\\M2Project_Y.m2project"
	//btm2project, err := ioutil.ReadFile(_path)
	//if err != nil {
	//	return err
	//}
	key := []byte("964CF629-E9EC-47B8-AAFE-39172793BEB4-D89A1212-DB56-4858-9ADE-CF19A782E5A3-0B92A08-3578-49B7-BDF6-21B73AC5D9E5\x00")

	encbt, err := hex.DecodeString("b0b4815d61eb7b8197889f")
	//encbt,err:=base.EncodeStream_uEDCode([]byte("111"),key)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString((encbt)))
	dbtstream, err := DecodeStream_uEDCode(encbt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dbtstream))

	dbtstream1, err := DecodeStream_uEDCode(encbt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dbtstream1))

}
