package base

import "golang.org/x/text/encoding/simplifiedchinese"

type Charset string
const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)
func ConvertString2Byte(_str string,charset Charset)string{
	var rstr string
	switch charset{
	case GB18030:
		var encodebyte,_=simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(_str))
		rstr=string(encodebyte)
	default:
		rstr=_str
	}
	return rstr
}
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
