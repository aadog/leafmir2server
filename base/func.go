package base

import (
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

var (
	Base64_CodeTable = []byte{
		'<', '>', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z', '~', '[', ']', '&', '^', '|', '?', '@', '{', '}', '=',
	}
)

func ConvertString2Byte(_str string, charset Charset) string {
	var rstr string
	switch charset {
	case GB18030:
		var encodebyte, _ = simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(_str))
		rstr = string(encodebyte)
	default:
		rstr = _str
	}
	return rstr
}
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
