package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"leafmir2server/base"
	"os"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	err:=m.Init()
	if err!=nil{
		log.Fatal("加载引擎失败,错误:%s",err.Error())
	}
}
func (m *Module)Init()error{
	_,err:=os.Stat("./!setup.txt")
	if err!=nil{
		return err
	}


	return nil
}

func (m *Module) OnDestroy() {

}
