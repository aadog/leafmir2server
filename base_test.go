package main

import (
	"fmt"
	"leafmir2server/base"
	"testing"
)

func TestBlowfishDecrypt(t *testing.T) {

	c := ("nimhJaix1yjuBzW0XWdKBZrLY5sL58Hsngtz2ud0uCJOtV2aQhFf6PRlUYNV+Q7OBSW2dy2vaxYZngIl7xekDAXdSUghiaxC90xJFcTR2uU3DT8S")
	r := base.BlowfishDecrypt([]byte(c), []byte("#$Ggy%(*^45fghj@@#sqw[]KHG%^&UHBR#$ty"))
	//if err!=nil{
	//	log.Fatal(err.Error())
	//}
	fmt.Println(string(r))
}
