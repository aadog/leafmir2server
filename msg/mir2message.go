package msg

import (
	"fmt"
	"leafmir2server/base"
)

type Mir2Message struct{
	Recog   int32       //4
	Ident   int16          //2
	Param   int16          //2
	Tag     int16         //2
	Series  int16          //2
	Lines []string
}

func (this* Mir2Message) Stringlines()string{
	r:=""
	for i,s:=range this.Lines{
		if i!=0{
			r=fmt.Sprintf("%s/",r)
		}
		r=fmt.Sprintf("%s%s",r,s)
	}
	return base.ConvertString2Byte(r,base.GB18030)
}
func NewMir2Message_with_msg_soul_wparam_atag_nseries(_msg int16,_soul int32,_wparam,_atag,_nseries int16,_lines ...string)*Mir2Message{
	lines:=[]string{}
	for _,line:=range _lines{
		lines=append(lines,line)
	}
	m:=&Mir2Message{_soul,_msg,_wparam,_atag,_nseries,lines}
	return m
}

