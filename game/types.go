package game

import "leafmir2server/msg"

const (
	TCharDescSize = 43
)

type TStatusSec uint16
type TStatusTime int64
type ShortInt uint8

type TItemPointPropType int
type TAddPoint [6]TAddPointItem
type TAddLevel [5]TAddLevelItem
type TAddHold [2]int32
type TDateTime float64

type TGList struct {
}
type TAddLevelItem struct {
	State     byte               //0:没启用 1:已启用没激活 2:已激活
	ValueType TItemPointPropType //1:吸收伤害 2:伤害反弹 3：伤害加成 4：致命一击 5：会心一击 6:体力恢复 7：魔法恢复 8:经验加成 9:宝物获得几率增加 10：金币获取增加百分比 11：命中  12：敏捷 13：体力 14：魔法值
	Value     uint16
}
type TAddPointItem struct {
	ValueType TItemPointPropType
	Value     uint16
}
type TAddProperty struct {
	Upgrade    byte
	MaxUpgrade byte
	Color      byte
	Light      byte
	CustomName bool
	SoulLevel  byte
	Effect     uint16
	Undefined1 uint16
	Undefined2 uint16
	Undefined3 uint16
	Speed      uint16
	Luck       uint16
}

type TRunAbility struct { //Size 40
	Level            int32 //
	ACMin            int32
	ACMax            int32
	MACMin           int32
	MACMax           int32
	DCMin            int32
	DCMax            int32
	MCMin            int32
	MCMax            int32
	SCMin            int32
	SCMax            int32
	TcMin            int32
	TcMax            int32
	PcMin            int32
	PcMax            int32
	WcMin            int32
	WcMax            int32
	HP               int32    //21亿血
	MP               int32    //21亿魔法
	MaxHP            int32    //21亿血
	MaxMP            int32    //21亿魔法
	Exp              int64    //
	MaxExp           int64    //经验修改为Int64
	Weight           uint16   //
	MaxWeight        uint16   // 背包
	WearWeight       uint16   //
	MaxWearWeight    uint16   //负重
	HandWeight       uint16   //
	MaxHandWeight    uint16   //腕力
	Alcohol          uint16   //酒量
	MaxAlcohol       uint16   //酒量上限
	WineDrinkValue   uint16   //醉酒度
	MedicineValue    uint16   //当前药力值
	MaxMedicineValue uint16   //药力值上限
	NG               uint16   //当前内力值
	MaxNG            uint16   //内力值上限
	Reserver         [20]byte //预留20个字节
}
type TAbility struct {
	Level            int32 //
	ACMin            int32
	ACMax            int32
	MACMin           int32
	MACMax           int32
	DCMin            int32
	DCMax            int32
	MCMin            int32
	MCMax            int32
	SCMin            int32
	SCMax            int32
	TcMin            int32
	TcMax            int32
	PcMin            int32
	PcMax            int32
	WcMin            int32
	WcMax            int32
	HP               int32    //21亿血
	MP               int32    //21亿魔法
	MaxHP            int32    //21亿血
	MaxMP            int32    //21亿魔法
	Exp              int64    //
	MaxExp           int64    //经验修改为Int64
	Weight           uint16   //
	MaxWeight        uint16   // 背包
	WearWeight       uint16   //
	MaxWearWeight    uint16   //负重
	HandWeight       uint16   //
	MaxHandWeight    uint16   //腕力
	Alcohol          uint16   //酒量
	MaxAlcohol       uint16   //酒量上限
	WineDrinkValue   uint16   //醉酒度
	MedicineValue    uint16   //当前药力值
	MaxMedicineValue uint16   //药力值上限
	NG               uint16   //当前内力值
	MaxNG            uint16   //内力值上限
	Reserver         [20]byte //预留20个字节
}
type TNakedAbility struct { //Size 20
	DC    uint16
	MC    uint16
	SC    uint16
	TC    uint16
	PC    uint16
	WC    uint16
	AC    uint16
	MAC   uint16
	HP    int32 //21亿血
	MP    int32 //21亿魔法
	Hit   uint16
	Speed uint16
	X2    uint16
}
type TBatterZhuiXinMessage struct {
	msg.Mir2Message
	Desc TCharDesc
	X    uint16
	Y    uint16
	Dir  byte
}
type TCharDesc struct {
	Feature              int32 //  种族 是否摆摊 性别 和 头发
	Status               int32 //状态
	HP                   int32
	MaxHP                int32
	wChangeToMonsterAppr uint16
	wChangetoMonsterRace uint16
	BodyBlendColor       byte  //混合绘制的身体 颜色
	WeaponBlendColor     byte  //混合绘制的武器颜色
	Properties           int32 //职业 性别 坐骑
	DressWepon           int32 //武器和衣服
	DressWeponEffect     int32 //武器和衣服的特效
	HorseAndLeft         int32 //坐骑 与左手(盾)
	Level                int32
	GroupMember          byte
}
type StdItem struct {
	Name             string
	Idx              string
	Weight           string
	Anicount         string
	Looks            string
	DuraMax          string
	Ac               string
	Price            string
	Color            string
	State            string
	AppendProperties string
}
type MapItem struct {
	Name             string
	MapName          string
	MapFileName      string
	MiniMapIndex     string
	Extend           string
	NewMapProperties string
	MineDrop         string
}
