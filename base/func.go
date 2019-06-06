package base

import (
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
	"golang.org/x/crypto/blowfish"
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
func D91_Base64Encode(_in []byte) {

}

func BlowfishDecrypt(et, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(et))
	mode.CryptBlocks(pt, et)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}

//
//func BlowfishDecrypt(value, key []byte) ([]byte, error) {
//	//deval,err:=base64.StdEncoding.DecodeString(string(value))
//	//if err!=nil{
//	//	return nil,err
//	//}
//	//value=deval
//	log.Println(string(value))
//
//	dcipher, err := blowfish.NewCipher(key)
//	if (err != nil) { return nil, err }
//	div := value[:blowfish.BlockSize]
//	decrypted := value[blowfish.BlockSize:]
//	if len(decrypted)%blowfish.BlockSize != 0 { return nil, err }
//	dcbc := cipher.NewCBCDecrypter(dcipher, div)
//	dcbc.CryptBlocks(decrypted, decrypted)
//	return decrypted, nil
//}
