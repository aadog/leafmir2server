package base

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	GBK = Charset("GBK")
)

func ConvertString2Byte(_str string, charset Charset) string {

	var rstr string
	switch charset {
	case GBK:
		var encodebyte, _ = simplifiedchinese.GBK.NewEncoder().Bytes([]byte(_str))
		rstr = string(encodebyte)
	default:
		rstr = _str
	}
	return rstr
}
func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GBK:

		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	default:
		str = string(byte)
	}
	return str
}
func Decode6BitBytes(src []byte) []byte {
	var decode6BitMask = [...]byte{0xfc, 0xf8, 0xf0, 0xe0, 0xc0}
	var size = len(src)
	var dest = make([]byte, size*3/4)
	var destPos = 0
	var bitPos uint = 2
	var madeBit uint = 0

	var ch byte = 0
	var chCode byte = 0
	var tmp byte = 0

	for i := 0; i < size; i++ {
		if (src[i] - 0x3c) >= 0 {
			ch = byte(src[i] - 0x3c)
		} else {
			destPos = 0
			break
		}

		if destPos >= len(dest) {
			break
		}

		if madeBit+6 >= 8 {
			chCode = byte(tmp | ((ch & 0x3f) >> uint(6-bitPos)))

			dest[destPos] = chCode
			destPos += 1

			madeBit = 0
			if bitPos < 6 {
				bitPos += 2
			} else {
				bitPos = 2
				continue
			}
		}

		tmp = (byte)((ch << bitPos) & decode6BitMask[bitPos-2])

		madeBit += 8 - bitPos
	}

	return dest
}
func Encoder6BitBuf(src []byte) []byte {
	var size = len(src)
	var destLen = (size/3)*4 + 10
	var dest = make([]byte, destLen)
	var destPos = 0
	var resetCount = 0

	var chMade, chRest byte = 0, 0

	for i := 0; i < size; i++ {
		if destPos >= destLen {
			break
		}

		chMade = (byte)((chRest | ((src[i] & 0xff) >> uint(2+resetCount))) & 0x3f)
		chRest = (byte)((((src[i] & 0xff) << uint(8-(2+resetCount))) >> uint(2)) & 0x3f)

		resetCount += 2
		if resetCount < 6 {
			dest[destPos] = (byte)(chMade + 0x3c)
			destPos += 1
		} else {
			if destPos < destLen-1 {
				dest[destPos] = (byte)(chMade + 0x3c)
				destPos += 1
				dest[destPos] = (byte)(chRest + 0x3c)
				destPos += 1
			} else {
				dest[destPos] = (byte)(chMade + 0x3c)
				destPos += 1
			}

			resetCount = 0
			chRest = 0
		}
	}
	if resetCount > 0 {
		dest[destPos] = (byte)(chRest + 0x3c)
		destPos += 1
	}

	dest[destPos] = 0
	return dest[:destPos]

}
func Md5_with(_in []byte) []byte {
	h := md5.New()
	h.Write(_in)
	return h.Sum(nil)
}
func Reskey(_s string) string {
	return fmt.Sprintf("%s_Res", _s)
}
func Reqkey(_s string) string {
	return fmt.Sprintf("%s_Req", _s)
}
