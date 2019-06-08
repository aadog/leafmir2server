package internal

import (
	"bytes"
	"compress/zlib"
	"errors"
	"github.com/beevik/etree"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"io"
	"io/ioutil"
	"leafmir2server/base"
	"os"
)

var (
	skeleton    = base.NewSkeleton()
	ChanRPC     = skeleton.ChanRPCServer
	stditemlist map[string]*StdItem
	mapitemlist map[string]*MapItem
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	stditemlist = make(map[string]*StdItem)
	mapitemlist = make(map[string]*MapItem)
	m.Skeleton = skeleton

	err := m.Init()
	if err != nil {
		log.Fatal("加载引擎失败,错误:%s", err.Error())
	}
}

func (m *Module) Loadmapitemlist_with_element(_element *etree.Element) error {
	log.Release("开始加载地图数据")
	eles := _element.SelectElements("TuMapItem")
	for _, ele := range eles {
		it := MapItem{}
		it_name := ele.SelectElement("Name")
		if it_name != nil {
			it.Name = it_name.Text()
		}
		it_mapname := ele.SelectElement("MapName")
		if it_mapname != nil {
			it.MapName = it_mapname.Text()
		}
		it_mapfilename := ele.SelectElement("MapFileName")
		if it_mapfilename != nil {
			it.MapName = it_mapfilename.Text()
		}
		it_minimapindex := ele.SelectElement("MiniMapIndex")
		if it_minimapindex != nil {
			it.MiniMapIndex = it_minimapindex.Text()
		}
		it_extend := ele.SelectElement("Extend")
		if it_extend != nil {
			it.Extend = it_extend.Text()
		}
		it_newmapproperties := ele.SelectElement("NewMapProperties")
		if it_newmapproperties != nil {
			it.NewMapProperties = it_newmapproperties.Text()
		}
		it_minedrop := ele.SelectElement("MineDrop")
		if it_minedrop != nil {
			it.MineDrop = it_minedrop.Text()
		}
		mapitemlist[it_name.Text()] = &it
	}
	log.Release("加载地图数据完成:%d/条", len(mapitemlist))
	return nil
}

func (m *Module) Loadstditemlist_with_element(_element *etree.Element) error {
	log.Release("开始加载物品数据")
	eles := _element.SelectElement("Items").SelectElements("TuStdItem")
	for _, ele := range eles {
		it := StdItem{}
		it_name := ele.SelectElement("Name")
		if it_name != nil {
			it.Name = it_name.Text()
		}
		it_idx := ele.SelectElement("Idx")
		if it_idx != nil {
			it.Idx = it_idx.Text()
		}
		it_Weight := ele.SelectElement("Weight")
		if it_Weight != nil {
			it.Weight = it_Weight.Text()
		}
		it_Anicount := ele.SelectElement("Anicount")
		if it_Anicount != nil {
			it.Anicount = it_Weight.Text()
		}
		it_Looks := ele.SelectElement("Looks")
		if it_Looks != nil {
			it.Looks = it_Looks.Text()
		}
		it_DuraMax := ele.SelectElement("DuraMax")
		if it_DuraMax != nil {
			it.DuraMax = it_DuraMax.Text()
		}
		it_Ac := ele.SelectElement("Ac")
		if it_Ac != nil {
			it.Ac = it_Ac.Text()
		}
		it_Price := ele.SelectElement("Price")
		if it_Price != nil {
			it.Price = it_Price.Text()
		}
		it_Color := ele.SelectElement("Color")
		if it_Color != nil {
			it.Color = it_Color.Text()
		}
		it_State := ele.SelectElement("State")
		if it_State != nil {
			it.State = it_State.Text()
		}
		it_AppendProperties := ele.SelectElement("AppendProperties")
		if it_AppendProperties != nil {
			it.AppendProperties = it_AppendProperties.Text()
		}

		stditemlist[it_idx.Text()] = &it
	}
	log.Release("加载物品数据成功:%d/条", len(stditemlist))
	return nil
}
func (m *Module) Init() error { //"C:\\MirServer\\Mir200\\M2Project_Y.m2project"
	m2projectpath := "C:\\MirServer\\Mir200\\M2Project_Y.m2project"

	_, err := os.Stat("./!setup.txt")
	if err != nil {
		return err
	}
	btm2project, err := DecodeM2project_with_path(m2projectpath)
	if err != nil {
		return err
	}
	log.Release("解密成功,版本文件:%s", m2projectpath)

	xmldoc := etree.NewDocument()
	err = xmldoc.ReadFromBytes(btm2project)
	if err != nil {
		return err
	}

	log.Release("%s", xmldoc.FindElement("//TM2Project/Name").Text())
	log.Release("脚本语言:%s", xmldoc.FindElement("//TM2Project/ScriptLanguage").Text())

	elestditemlist := xmldoc.FindElement("//TM2Project/GameData")
	if elestditemlist == nil {
		return errors.New("找不到物品数据库在m2project")
	}
	err = m.Loadstditemlist_with_element(elestditemlist)
	if err != nil {
		return err
	}
	elemapitemlist := xmldoc.FindElement("//TM2Project/Maps")
	err = m.Loadmapitemlist_with_element(elemapitemlist)
	if err != nil {
		return err
	}

	return nil
}
func DecodeM2project_with_path(_path string) ([]byte, error) {
	btm2project, err := ioutil.ReadFile(_path)
	if err != nil {
		return nil, err
	}
	dbtstream, err := base.DecodeStream(btm2project, []byte("964CF629-E9EC-47B8-AAFE-39172793BEB4-D89A1212-DB56-4858-9ADE-CF19A782E5A3-0B92A08-3578-49B7-BDF6-21B73AC5D9E5\x00"))
	if err != nil {
		return nil, err
	}

	debtm2project, err := DoZlibUnCompress(dbtstream)
	if err != nil {
		return nil, err
	}
	return debtm2project, nil
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) ([]byte, error) {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func (m *Module) OnDestroy() {

}
