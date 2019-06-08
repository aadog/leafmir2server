package base

import (
	"github.com/name5566/leaf/log"
	"syscall"
	"unsafe"
)

var (
	c_project     = syscall.NewLazyDLL("./project.dll")
	cEncodeStream = c_project.NewProc("cEncodeStream")
	cDecodeStream = c_project.NewProc("cDecodeStream")
)

func EncodeStream(_in []byte, _key []byte) ([]byte, error) {
	var rlen int
	cEncodeStream.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	r := make([]byte, rlen)
	cEncodeStream.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	return r, nil
}
func DecodeStream(_in []byte, _key []byte) ([]byte, error) {
	var rlen int
	cDecodeStream.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(0),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	r := make([]byte, rlen)
	cDecodeStream.Call(
		uintptr(unsafe.Pointer(&_in[0])),
		uintptr(len(_in)),
		uintptr(unsafe.Pointer(&r[0])),
		uintptr(unsafe.Pointer(&rlen)),
		uintptr(unsafe.Pointer(&_key[0])),
	)
	return r, nil
}

func init() {
	err := c_project.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}
