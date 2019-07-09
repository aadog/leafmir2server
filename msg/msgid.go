package msg

const (
	_JOB_WAR    = 0 //战士
	_JOB_MAG    = 1 //法师
	_JOB_DAO    = 2 //道士
	_JOB_CIK    = 3 //刺客
	_JOB_ARCHER = 4 //弓箭手
	_JOB_SHAMAN = 5 //和尚
	//传奇中人物只有8个方向,但是打符,锲蛾飞行,神鹰都有16方向
	DR_UP        = 0 //北
	DR_UPRIGHT   = 1 //东北向
	DR_RIGHT     = 2 //东
	DR_DOWNRIGHT = 3 //东南向
	DR_DOWN      = 4 //南
	DR_DOWNLEFT  = 5 //西南向
	DR_LEFT      = 6 //西
	DR_UPLEFT    = 7 //西北向

	//装备项目
	U_DRESS     = 0  //衣服
	U_WEAPON    = 1  //武器
	U_RIGHTHAND = 2  //右手 勋章  蜡烛
	U_NECKLACE  = 3  //项链
	U_HELMET    = 4  //头盔
	U_ARMRINGL  = 5  //左手手镯,符
	U_ARMRINGR  = 6  //右手手镯
	U_RINGL     = 7  //左戒指
	U_RINGR     = 8  //右戒指
	U_BUJUK     = 9  //物品(毒符)
	U_BELT      = 10 //腰带
	U_BOOTS     = 11 //鞋
	U_CHARM     = 12 //宝石
	U_ZHULI     = 13 //斗笠
	U_FASHION   = 14 //时装衣服
	U_MOUNT     = 15 //坐骑
	U_SHIED     = 16 //盾牌

	U_JEWELRYITEM1 = 17 //首饰盒1
	U_JEWELRYITEM2 = 18 //首饰盒2
	U_JEWELRYITEM3 = 19 //首饰盒3
	U_JEWELRYITEM4 = 20 //首饰盒4
	U_JEWELRYITEM5 = 21 //首饰盒5
	U_JEWELRYITEM6 = 22 //首饰盒6

	U_ZODIAC1       = 23 // 十二生肖
	U_ZODIAC2       = 24 // 十二生肖
	U_ZODIAC3       = 25 // 十二生肖
	U_ZODIAC4       = 26 // 十二生肖
	U_ZODIAC5       = 27 // 十二生肖
	U_ZODIAC6       = 28 // 十二生肖
	U_ZODIAC7       = 29 // 十二生肖
	U_ZODIAC8       = 30 // 十二生肖
	U_ZODIAC9       = 31 // 十二生肖
	U_ZODIAC10      = 32 // 十二生肖
	U_ZODIAC11      = 33 // 十二生肖
	U_ZODIAC12      = 34 // 十二生肖
	U_RESERVER      = 38
	U_MAXUSEITEMIDX = 38 //最大的使用物品序号

	HAM_ALL      = 0 //[攻击模式: 全体攻击]
	HAM_PEACE    = 1 //[攻击模式: 和平攻击]
	HAM_DEAR     = 2 //[攻击模式: 夫妻攻击]
	HAM_MASTER   = 3 //[攻击模式: 师徒攻击]
	HAM_GROUP    = 4 //[攻击模式: 编组攻击]
	HAM_GUILD    = 5 //[攻击模式: 行会攻击]
	HAM_PKATTACK = 6 //[攻击模式: 红名攻击]
	HAM_NATION   = 7 //[攻击模式：国家攻击]
	HAM_CAMP     = 8 //[攻击模式：阵营攻击]

	//动态的 有时间限制
	STATE_POISON_DECHEALTH   = 0  //中毒类型：绿毒
	STATE_POISON_DAMAGEARMOR = 1  //中毒类型：红毒
	STATE_POISON_LOCKSPELL   = 2  //不能攻击
	STATE_LOCKRUN            = 3  //不能跑动(中蛛网) 20080811
	STATE_POISON_DONTMOVE    = 4  //不能移动
	STATE_POISON_STONE       = 5  //中毒类型:麻痹
	STATE_POISON_TRAP        = 6  //陷阱不能移动
	STATE_ProtectionDEFENCE  = 7  //护体神盾 20080107
	STATE_TRANSPARENT        = 8  //隐身
	STATE_DEFENCEUP          = 9  //神圣战甲术  防御力
	STATE_MAGDEFENCEUP       = 10 //幽灵盾  魔御力
	STATE_BUBBLEDEFENCEUP    = 11 //魔法盾
	STATE_BECOOL             = 12 //冷酷 通常用于刺客
	STATE_ONLYWALK           = 13 //只能行走不允许 跑动 致残毒药
	STATE_HITDRUG            = 14 //致残毒药状态 在这种状态下 有几率 可以对敌人进行致残打击
	MAX_STATUS_ATTRIBUTE     = 15 //状态最大值 增加自定义状态需要修改这个顺序

	//静态玩家手动变更  注意 这些和上面的 组成一个4字节字段
	STATE_STONE_MODE       = 20 //石化状态 主要用于祖玛教主 和  祖玛雕像 这类
	POISON_68              = 21
	STATE_OPENHEATH        = 22 //显示血量
	STATE_ALLOWGROUP       = 23 //允许组队
	STATE_ALLOWGUILD       = 24 //允许行会邀请
	STATE_ALLOWGUILDRECALL = 25 //允许行会合一
	STATE_ALLOWGROUPRECALL = 26 //允许组队合一
	STATE_ALLOWDEAL        = 27 //允许交易
	STATE_ALLOWRELIVE      = 28 //允许复活
	STATE_ALLOWFASHION     = 29 //显示时装
	STATE_Recruit          = 30 //招募队员
	STATE_Frozen           = 31 //冰冻状态

	//绑定  商城里的
	_ITEM_STATE_BIND           = 0  //(绑定)不可交易,不可出售,丢弃消失
	_ITEM_STATE_NEVERDROP      = 1  //永不掉落
	_ITEM_STATE_NOREPAIR       = 2  //不可修理
	_ITEM_STATE_NOTSTORE       = 3  //不可存仓
	_ITEM_STATE_OFFLINEFREE    = 4  //离线消失
	_ITEM_STATE_DIEFREE        = 5  //死亡消失
	_ITEM_STATE_NOTDROP        = 6  //不可丢弃
	_ITEM_STATE_TAKEONBIND     = 7  //穿戴后绑定
	_ITEM_STATE_NOTTAKEOFF     = 8  //穿戴后无法取下
	_ITEM_STATE_DIEFREE_Latent = 9  //死亡消失
	_ITEM_STATE_DIEDROP_Latent = 10 //死亡必爆

	//物品库基础状态
	USERMODE_PLAYGAME = 1
	USERMODE_LOGIN    = 2
	USERMODE_LOGOFF   = 3
	USERMODE_NOTICE   = 4

	RUNGATEMAX = 8

	RUNGATECODE = 0xAA55AA55

	OS_MOVINGOBJECT = 1
	OS_ITEMOBJECT   = 2
	OS_EVENTOBJECT  = 3
	OS_GATEOBJECT   = 4
	OS_SWITCHOBJECT = 5
	OS_MAPEVENT     = 6
	OS_DOOR         = 7
	OS_ROON         = 8

	RC_PLAYOBJECT  = 0   //人物
	RC_PLAYMONSTER = 150 //人形怪物
	RC_GUARD       = 12  //大刀守卫
	RC_PEACENPC    = 15  //固定NPC
	RC_ANIMAL      = 50  //无效
	RC_MONSTER     = 80  //怪物
	RC_NPC         = 10  //商人
	RC_CAMION      = 31  //随从 镖车
	RC_Trainer     = 55  //练功师
	RC_BOX         = 30  //宝箱怪物
	RC_ARCHERGUARD = 112 //NPC 弓箭手
	RC_HEROOBJECT  = 200 //英雄

	RCC_USERHUMAN = RC_PLAYOBJECT
	RCC_GUARD     = RC_GUARD
	RCC_MERCHANT  = RC_ANIMAL

	ISM_WHISPER = 1234
	//STARTMESSAGE
	CM_QUERYCHR     = 100 //登录成功,客户端显出左右角色的那一瞬
	CM_NEWCHR       = 101 //创建角色
	CM_DELCHR       = 102 //删除角色
	CM_SELCHR       = 103 //选择角色
	CM_SELECTSERVER = 104 //服务器,注意不是选区,一区往往有(至多8个??group.dat中是这么写的)不止一个的服务器
	CM_QUERYDELCHR  = 105 //查询删除过的角色信息 20080706
	CM_RESDELCHR    = 106 //恢复删除的角色 20080706

	SM_HORSERUN = 5
	SM_RUSH     = 6 //跑动中改变方向
	SM_RUSHKUNG = 7 //野蛮冲撞
	SM_FIREHIT  = 8 //烈火

	SM_BACKSTEP       = 9  //后退,野蛮效果? //半兽统领公箭手攻击玩家的后退??axemon.pas中procedure   TDualAxeOma.Run
	SM_TURN           = 10 //转向
	SM_WALK           = 11 //走
	SM_SITDOWN        = 12
	SM_RUN            = 13 //跑
	SM_HIT            = 14 //砍
	SM_HEAVYHIT       = 15 //
	SM_BIGHIT         = 16 //
	SM_SPELL          = 17 //使用魔法
	SM_POWERHIT       = 18 //攻杀
	SM_LONGHIT        = 19 //刺杀
	SM_DIGUP          = 20 //挖是一"起"一"坐",这里是挖动作的"起"
	SM_DIGDOWN        = 21 //挖动作的"坐"
	SM_FLYAXE         = 22 //飞斧,半兽统领的攻击方式?
	SM_LIGHTING       = 23 //免蜡开关
	SM_WIDEHIT        = 24 //半月
	SM_CRSHIT         = 25 //抱月刀
	SM_TWINHIT        = 26 //开天斩重击
	SM_ALIVE          = 27 //复活??复活戒指
	SM_MOVEFAIL       = 28 //移动失败,走动或跑动
	SM_HIDE           = 29 //隐身?
	SM_DISAPPEAR      = 30 //地上物品消失
	SM_STRUCK         = 31 //受攻击
	SM_DEATH          = 32 //正常死亡
	SM_SKELETON       = 33 //尸体
	SM_NOWDEATH       = 34 //秒杀?
	SM_CHANGEDIR      = 35 //改变方向 随云
	SM_HEAR           = 40 //有人回你的话
	SM_FEATURECHANGED = 41
	SM_USERNAME       = 42
	SM_WINEXP         = 44 //获得经验
	SM_LEVELUP        = 45 //升级,左上角出现墨绿的升级字样
	SM_DAYCHANGING    = 46 //传奇界面右下角的太阳星星月亮

	SM_LOGON              = 50 //logon
	SM_NEWMAP             = 51 //新地图??
	SM_ABILITY            = 52 //打开属性对话框,F11
	SM_HEALTHSPELLCHANGED = 53 //治愈术使你的体力增加

	SM_CIDHIT   = 57 //龙影剑法
	SM_4FIREHIT = 58 //4级烈火 20080112
	SM_QTWINHIT = 59 //开天斩轻击

	SM_SANJUEHIT  = 60 //三绝杀
	SM_ZHUIXINHIT = 61 //追心刺
	SM_DUANYUEHIT = 62 //断岳斩
	SM_HENGSAOHIT = 63 //横扫千军
	SM_YTPDHIT    = 64 //倚天劈地
	SM_XPYJHIT    = 65 //血魄一击

	SM_4LONGHIT   = 66 //四级刺杀
	SM_YUANYUEHIT = 67 //圆月弯刀

	SM_STRUCK_MISS = 3101 //MISS
	SM_STORM       = 4000

	SM_ACTION_MIN  = SM_HORSERUN
	SM_ACTION_MAX  = SM_WIDEHIT
	SM_ACTION2_MIN = 4001
	SM_ACTION2_MAX = 4002

	SM_GROUPHEALTHSPELLCHANGED = 4003 //队友体力
	SM_GROUPLEVELCHANGED       = 4004 //队员等级
	SM_MAPDESCRIPTION          = 54   //地图描述,行会战地图?攻城区域?安全区域?
	SM_SPELL2                  = 117

	//对话消息
	SM_MOVEMESSAGE       = 99
	SM_DELETEMOVEMESSAGE = 98
	SM_SYSMESSAGE        = 100 //系统消息,盛大一般红字,私服蓝字
	SM_GROUPMESSAGE      = 101 //组内聊天!!
	SM_CRY               = 102 //喊话
	SM_WHISPER           = 103 //私聊
	SM_GUILDMESSAGE      = 104 //行会聊天!~
	SM_NATIONMESSAGE     = 105
	SM_CAMPMESSAGE       = 106

	SM_ADDITEM     = 200
	SM_BAGITEMS    = 201
	SM_DELITEM     = 202
	SM_UPDATEITEM  = 203
	SM_ADDMAGIC    = 210
	SM_SENDMYMAGIC = 211
	SM_DELMAGIC    = 212
	SM_BACKSTEPEX  = 250

	//服务器端发送的命令 SM:server msg,服务端向客户端发送的消息

	//登录、新帐号、新角色、查询角色等
	SM_CERTIFICATION_FAIL    = 501
	SM_ID_NOTFOUND           = 502
	SM_PASSWD_FAIL           = 503 //验证失败,"服务器验证失败,需要重新登录"??
	SM_NEWID_SUCCESS         = 504 //创建新账号成功
	SM_NEWID_FAIL            = 505 //创建新账号失败
	SM_CHGPASSWD_SUCCESS     = 506 //修改密码成功
	SM_CHGPASSWD_FAIL        = 507 //修改密码失败
	SM_GETBACKPASSWD_SUCCESS = 508 //密码找回成功
	SM_GETBACKPASSWD_FAIL    = 509 //密码找回失败

	SM_QUERYCHR       = 520 //返回角色信息到客户端
	SM_NEWCHR_SUCCESS = 521 //新建角色成功
	SM_NEWCHR_FAIL    = 522 //新建角色失败
	SM_DELCHR_SUCCESS = 523 //删除角色成功
	SM_DELCHR_FAIL    = 524 //删除角色失败
	SM_STARTPLAY      = 525 //开始进入游戏世界(点了健康游戏忠告后进入游戏画面)
	SM_STARTFAIL      = 526 ////开始失败,玩传奇深有体会,有时选择角色,点健康游戏忠告后黑屏

	SM_QUERYCHR_FAIL       = 527 //返回角色信息到客户端失败
	SM_OUTOFCONNECTION     = 528 //超过最大连接数,强迫用户下线
	SM_PASSOK_SELECTSERVER = 529 //密码验证完成且密码正确,开始选服
	SM_SELECTSERVER_OK     = 530 //选服成功
	SM_NEEDUPDATE_ACCOUNT  = 531 //需要更新,注册后的ID会发生什么变化?私服中的普通ID经过充值??或者由普通ID变为会员ID,GM?
	SM_UPDATEID_SUCCESS    = 532 //更新成功
	SM_UPDATEID_FAIL       = 533 //更新失败
	SM_SESSIONID           = 539

	SM_QUERYDELCHR       = 534 //返回删除过的角色 20080706
	SM_QUERYDELCHR_FAIL  = 535 //返回删除过的角色失败 20080706
	SM_RESDELCHR_SUCCESS = 536 //恢复删除角色成功 20080706
	SM_RESDELCHR_FAIL    = 537 //恢复删除角色失败 20080706
	SM_NOCANRESDELCHR    = 538 //禁止恢复删除角色,即不可查看 200800706

	SM_DROPITEM_SUCCESS = 600
	SM_DROPITEM_FAIL    = 601

	RM_SPEEDTIME = 10000
	SM_SPEEDTIME = 603

	SM_ITEMSHOW = 610
	SM_ITEMHIDE = 611
	//  SM_DOOROPEN           = 612
	SM_OPENDOOR_OK                 = 612 //通过过门点成功
	SM_OPENDOOR_LOCK               = 613 //发现过门口是封锁的,以前盛大秘密通道去赤月的门要5分钟开一次
	SM_CLOSEDOOR                   = 614 //用户过门,门自行关闭
	SM_TAKEON_OK                   = 615
	SM_TAKEON_FAIL                 = 616
	SM_TAKEOFF_OK                  = 619
	SM_TAKEOFF_FAIL                = 620
	SM_SENDUSEITEMS                = 621
	SM_WEIGHTCHANGED               = 622
	SM_CLEAROBJECTS                = 633
	SM_CHANGEMAP                   = 634 //地图改变,进入新地图
	SM_EAT_OK                      = 635
	SM_EAT_FAIL                    = 636
	SM_BUTCH                       = 637 //野蛮?
	SM_MAGICFIRE                   = 638 //地狱火,火墙??
	SM_MAGICFIRE_FAIL              = 639
	SM_MAGIC_LVEXP                 = 640
	SM_DURACHANGE                  = 642
	SM_MERCHANTSAY                 = 643
	SM_MERCHANTDLGCLOSE            = 644
	SM_SENDGOODSLIST               = 645
	SM_SENDUSERSELL                = 646
	SM_SENDBUYPRICE                = 647
	SM_USERSELLITEM_OK             = 648
	SM_USERSELLITEM_FAIL           = 649
	SM_BUYITEM_SUCCESS             = 650 //?
	SM_BUYITEM_FAIL                = 651 //?
	SM_SENDDETAILGOODSLIST         = 652
	SM_GOLDCHANGED                 = 653
	SM_CHANGELIGHT                 = 654 //负重改变
	SM_LAMPCHANGEDURA              = 655 //蜡烛持久改变
	SM_CHANGENAMECOLOR             = 656 //名字颜色改变,白名,灰名,红名,黄名
	SM_CHARSTATUSCHANGED           = 657
	SM_SENDNOTICE                  = 658 //发送健康游戏忠告(公告)
	SM_GROUPMODECHANGED            = 659 //组队模式改变
	SM_CREATEGROUP_OK              = 660
	SM_CREATEGROUP_FAIL            = 661
	SM_GROUPADDMEM_OK              = 662
	SM_GROUPDELMEM_OK              = 663
	SM_GROUPADDMEM_FAIL            = 664
	SM_GROUPDELMEM_FAIL            = 665
	SM_GROUPCANCEL                 = 666
	SM_GROUPMEMBERS                = 667
	SM_SENDUSERREPAIR              = 668
	SM_USERREPAIRITEM_OK           = 669
	SM_USERREPAIRITEM_FAIL         = 670
	SM_SENDREPAIRCOST              = 671
	SM_DEALMENU                    = 673
	SM_DEALASK                     = 4005
	SM_DEALASKCANCEL               = 4006
	SM_DEALASKOK                   = 4007
	SM_DEALTRY_FAIL                = 674
	SM_DEALADDITEM_OK              = 675
	SM_DEALADDITEM_FAIL            = 676
	SM_DEALDELITEM_OK              = 677
	SM_DEALDELITEM_FAIL            = 678
	SM_DEALCANCEL                  = 681
	SM_DEALREMOTEADDITEM           = 682
	SM_DEALREMOTEDELITEM           = 683
	SM_DEALCHGGOLD_OK              = 684
	SM_DEALCHGGOLD_FAIL            = 685
	SM_DEALREMOTECHGGOLD           = 686
	SM_DEALSUCCESS                 = 687
	SM_DEALMESSAGE                 = 688
	SM_SENDUSERSTORAGEITEM         = 700
	SM_STORAGE_OK                  = 701
	SM_STORAGE_FULL                = 702
	SM_STORAGE_FAIL                = 703
	SM_SAVEITEMLIST                = 704
	SM_TAKEBACKSTORAGEITEM_OK      = 705
	SM_TAKEBACKSTORAGEITEM_FAIL    = 706
	SM_TAKEBACKSTORAGEITEM_FULLBAG = 707
	RM_CHANGEEFFIGYSTATE           = 10001
	SM_CHANGEEFFIGYSTATE           = 4008

	SM_AREASTATE = 708 //周围状态
	SM_MYSTATUS  = 766 //我的状态,最近一次下线状态,如是否被毒,挂了就强制回城

	SM_DELITEMS                 = 709
	SM_READMINIMAP_OK           = 710
	SM_READMINIMAP_FAIL         = 711
	SM_SENDUSERMAKEDRUGITEMLIST = 712
	SM_MAKEDRUG_SUCCESS         = 713
	//  714
	//  716
	SM_MAKEDRUG_FAIL = 4009

	SM_CHANGEGUILDNAME      = 750
	SM_SENDUSERSTATE        = 751 //
	SM_SUBABILITY           = 752 //打开输助属性对话框
	SM_OPENGUILDDLG         = 753 //
	SM_OPENGUILDDLG_FAIL    = 754 //
	SM_SENDGUILDMEMBERLIST  = 756 //
	SM_GUILDADDMEMBER_OK    = 757 //
	SM_GUILDADDMEMBER_FAIL  = 758
	SM_GUILDDELMEMBER_OK    = 759
	SM_GUILDDELMEMBER_FAIL  = 760
	SM_GUILDRANKUPDATE_FAIL = 761
	SM_BUILDGUILD_OK        = 762
	SM_BUILDGUILD_FAIL      = 763
	SM_DONATE_OK            = 764
	SM_DONATE_FAIL          = 765

	SM_MENU_OK             = 767 //?
	SM_GUILDMAKEALLY_OK    = 768
	SM_GUILDMAKEALLY_FAIL  = 769
	SM_GUILDBREAKALLY_OK   = 770 //?
	SM_GUILDBREAKALLY_FAIL = 771 //?
	SM_DLGMSG              = 772 //Jacky
	SM_SPACEMOVE_HIDE      = 800 //道士走一下隐身
	SM_SPACEMOVE_SHOW      = 801 //道士走一下由隐身变为现身
	SM_RECONNECT           = 802 //与服务器重连
	SM_GHOST               = 803 //尸体清除,虹魔教主死的效果?
	SM_SHOWEVENT           = 804 //显示事件
	SM_HIDEEVENT           = 805 //隐藏事件
	SM_SPACEMOVE_HIDE2     = 806
	SM_SPACEMOVE_SHOW2     = 807
	SM_HIDEGATE            = 808
	SM_SHOWGATE            = 809
	SM_TIMECHECK_MSG       = 810 //时钟检测,以免客户端作弊
	SM_ADJUST_BONUS        = 811 //?

	//----SM_消息 从5000开始添加----//
	SM_OPENPULSE_OK          = 4010
	SM_OPENPULSE_FAIL        = 4011
	SM_RUSHPULSE_OK          = 4012
	SM_RUSHPULSE_FAIL        = 4013
	SM_PULSECHANGED          = 4014
	SM_BATTERORDER           = 4015
	SM_CANUSEBATTER          = 4016
	SM_BATTERUSEFINALLY      = 4017
	SM_BATTERSTART           = 4018
	SM_OPENPULS              = 4019 //打开经络
	SM_HEROBATTERORDER       = 4020
	SM_HEROPULSECHANGED      = 4021
	SM_BATTERBACKSTEP        = 4022
	SM_STORMSRATE            = 4023
	SM_STORMSRATECHANGED     = 4024
	SM_HEROSTORMSRATECHANGED = 4025
	SM_OPENPULSENEEDLEV      = 4026
	SM_HEROATTECTMODE        = 4027
	SM_GETASSESSHEROINFO     = 4028
	SM_QUERYASSESSHERO       = 4029
	SM_SHOWASSESSDLG         = 4030
	SM_ISDEHERO              = 4031
	SM_OPENTRAININGDLG       = 4032

	SM_OPENHEALTH  = 1100
	SM_CLOSEHEALTH = 1101

	SM_BREAKWEAPON       = 1102 //武器破碎
	SM_INSTANCEHEALGUAGE = 1103 //实时治愈
	SM_CHANGEFACE        = 1104 //变脸,发型改变?
	SM_VERSION_FAIL      = 1106 //客户端版本验证失败

	SM_ITEMUPDATE = 1500
	SM_MONSTERSAY = 1501

	SM_EXCHGTAKEON_OK   = 4033
	SM_EXCHGTAKEON_FAIL = 4034

	SM_THROW = 4035

	RM_DELITEMS          = 10002 //Jacky
	RM_TURN              = 10003 //转向
	RM_WALK              = 10004
	RM_RUN               = 10005
	RM_HIT               = 10006
	RM_SHOOT             = 10007
	RM_SPELL             = 10008
	RM_SPELL2            = 10009
	RM_PUSH              = 10010
	RM_RUSH              = 10011
	RM_STRUCK            = 10012 //受物理打击
	RM_DEATH             = 10013
	RM_DISAPPEAR         = 10014
	RM_STORM             = 10015
	RM_MAGSTRUCK         = 10016
	RM_MAGHEALING        = 10017
	RM_STRUCK_MAG        = 10018 //受魔法打击
	RM_MAGSTRUCK_MINE    = 10019
	RM_INSTANCEHEALGUAGE = 10020 //jacky
	RM_HEAR              = 10021 //公聊
	RM_WHISPER           = 10022
	RM_CRY               = 10023
	RM_RIDE              = 10024
	RM_WINEXP            = 10025
	RM_USERNAME          = 10026
	RM_LEVELUP           = 10027
	RM_CHANGENAMECOLOR   = 10028
	RM_PUSHEX            = 10029
	RM_SOULEXP           = 10030

	RM_LOGON                   = 10031
	RM_ABILITY                 = 10032
	RM_HEALTHSPELLCHANGED      = 10033
	RM_GROUPHEALTHSPELLCHANGED = 10034
	RM_GROUPLEVELCHANGED       = 10035
	RM_DAYCHANGING             = 10036

	RM_MOVEMESSAGE       = 10037 //滚动公告   清清 2007.11.13
	RM_DELETEMOVEMESSAGE = 10038
	RM_SYSMESSAGE        = 10039
	RM_GROUPMESSAGE      = 10040
	RM_GUILDMESSAGE      = 10041
	RM_NATIONMESSAGE     = 10042
	RM_CAMPMESSAGE       = 10043
	RM_ITEMSHOW          = 10044
	RM_ITEMHIDE          = 10045
	RM_DOOROPEN          = 10046
	RM_DOORCLOSE         = 10047
	RM_SENDUSEITEMS      = 10048 //发送使用的物品
	RM_WEIGHTCHANGED     = 10049

	RM_FEATURECHANGED    = 10050
	RM_CLEAROBJECTS      = 10051
	RM_CHANGEMAP         = 10052
	RM_BUTCH             = 10053 //挖
	RM_MAGICFIRE         = 10054
	RM_SENDMYMAGIC       = 10055 //发送使用的魔法
	RM_MAGIC_LVEXP       = 10056
	RM_SKELETON          = 10057
	RM_DURACHANGE        = 10058 //持久改变
	RM_MERCHANTSAY       = 10059
	RM_GOLDCHANGED       = 10060
	RM_CHANGELIGHT       = 10061
	RM_CHARSTATUSCHANGED = 10062
	RM_DELAYMAGIC        = 10063

	RM_BUFFERADDED  = 10064
	RM_BUFFEREMOVED = 10065
	SM_BUFFERADDED  = 4036
	SM_BUFFEREMOVED = 4037

	RM_DIGUP       = 10066
	RM_DIGDOWN     = 10067
	RM_FLYAXE      = 10068
	RM_LIGHTING    = 10069
	RM_SUBABILITY  = 10070
	RM_TRANSPARENT = 10071

	RM_SPACEMOVE_SHOW  = 10072
	RM_RECONNECTION    = 10073
	RM_SPACEMOVE_SHOW2 = 10074 //?
	RM_HIDEEVENT       = 10075 //隐藏烟花
	RM_SHOWEVENT       = 10076 //显示烟花
	RM_HIDEGATE        = 10077
	RM_SHOWGATE        = 10078
	RM_ZEN_BEE         = 10079

	RM_OPENHEALTH   = 10080
	RM_CLOSEHEALTH  = 10081
	RM_DOOPENHEALTH = 10082
	RM_CHANGEFACE   = 10083

	RM_ITEMUPDATE = 10084
	RM_MONSTERSAY = 10085
	RM_MAKESLAVE  = 10086

	SM_PLAYSOUND  = 4038
	SM_PLAYVIDEO  = 4039
	SM_REQUESTURL = 4040
	RM_ZHUIXINHIT = 10087 //追心刺  人刚刚开始的动作
	RM_ZHUIXIN_OK = 10088 //追心刺  冲撞过去的动作
	RM_MONMOVE    = 10089
	SS_200        = 200
	SS_201        = 201
	SS_202        = 202
	SS_WHISPER    = 203
	SS_204        = 204
	SS_205        = 205
	SS_206        = 206
	SS_207        = 207
	SS_208        = 208
	SS_209        = 219
	SS_210        = 210
	SS_211        = 211
	SS_212        = 212
	SS_213        = 213
	SS_214        = 214

	RM_10205           = 10090
	RM_DAMAGE          = 10091
	RM_ALIVE           = 10092 //复活
	RM_CHANGEGUILDNAME = 10093
	RM_10414           = 10094
	RM_POISON          = 10095
	LA_UNDEAD          = 1 //不死系

	RM_DELAYPUSHED = 10096

	CM_GETBACKPASSWORD = 2010 //密码找回
	CM_SPELL           = 3017 //施魔法
	CM_QUERYUSERNAME   = 80   //进入游戏,服务器返回角色名到客户端

	CM_DROPITEM           = 1000 //从包裹里扔出物品到地图,此时人物如果在安全区可能会提示安全区不允许扔东西
	CM_PICKUP             = 1001 //捡东西
	CM_TAKEONITEM         = 1003 //装配装备到身上的装备位置
	CM_TAKEOFFITEM        = 1004 //从身上某个装备位置取下某个装备
	CM_EAT                = 1006 //吃药
	CM_BUTCH              = 1007 //挖
	CM_MAGICKEYCHANGE     = 1008 //魔法快捷键改变
	CM_HEROMAGICKEYCHANGE = 1046 //英雄魔法开关设置 20080606
	CM_1005               = 1005

	//与商店NPC交易相关
	CM_CLICKNPC                = 1010 //用户点击了某个NPC进行交互
	CM_MERCHANTDLGSELECT       = 1011 //商品选择,大类
	CM_MERCHANTQUERYSELLPRICE  = 1012 //返回价格,标准价格,我们知道商店用户卖入的有些东西掉持久或有特殊
	CM_USERSELLITEM            = 1013 //用户卖东西
	CM_USERBUYITEM             = 1014 //用户买入东西
	CM_USERGETDETAILITEM       = 1015 //取得商品清单,比如点击"蛇眼戒指"大类,会出现一列蛇眼戒指供你选择
	CM_DROPGOLD                = 1016 //用户放下金钱到地上
	CM_LOGINNOTICEOK           = 1018 //健康游戏忠告点了确实,进入游戏
	CM_GROUPMODE               = 1019 //关组还是开组
	CM_CREATEGROUP             = 1020 //新建组队
	CM_ADDGROUPMEMBER          = 1021 //组内添人
	CM_DELGROUPMEMBER          = 1022 //组内删人
	CM_GROUPLEAVE              = 4000 //离开队伍、解散队伍
	CM_USERREPAIRITEM          = 1023 //用户修理东西
	CM_MERCHANTQUERYREPAIRCOST = 1024 //客户端向NPC取得修理费用
	CM_DEALTRY                 = 1025 //开始交易,交易开始
	CM_AGREEDEAL               = 4001 //同意交易
	CM_REJECTDEAL              = 4002 //拒绝交易
	CM_DEALADDITEM             = 1026 //加东东到交易物品栏上
	CM_DEALDELITEM             = 1027 //从交易物品栏上撤回东东???好像不允许哦
	CM_DEALCANCEL              = 1028 //取消交易
	CM_DEALCHGGOLD             = 1029 //本来交易栏上金钱为0,,如有金钱交易,交易双方都会有这个消息
	CM_DEALEND                 = 1030 //交易成功,完成交易
	CM_USERSTORAGEITEM         = 1031 //用户寄存东西
	CM_USERTAKEBACKSTORAGEITEM = 1032 //用户向保管员取回东西
	CM_WANTMINIMAP             = 1033 //用户点击了"小地图"按钮
	CM_USERMAKEDRUGITEM        = 1034 //用户制造毒药(其它物品)
	CM_OPENGUILDDLG            = 1035 //用户点击了"行会"按钮
	CM_GUILDHOME               = 1036 //点击"行会主页"
	CM_GUILDMEMBERLIST         = 1037 //点击"成员列表"
	CM_GUILDADDMEMBER          = 1038 //增加成员
	CM_GUILDDELMEMBER          = 1039 //踢人出行会
	CM_GUILDUPDATENOTICE       = 1040 //修改行会公告
	CM_GUILDUPDATERANKINFO     = 1041 //更新联盟信息(取消或建立联盟)
	CM_ADJUST_BONUS            = 1043 //用户得到奖励??私服中比较明显,小号升级时会得出金钱声望等,不是很确定,//求经过测试的高手的验证
	CM_SPEEDHACKUSER           = 4003 //用户加速作弊检测

	CM_PASSWORD    = 1105
	CM_CHGPASSWORD = 1221 //?
	CM_SETPASSWORD = 1222 //?

	CM_HORSERUN = 3009

	CM_THROW = 3005 //射箭、抛符

	CM_ADDITEMTOJEWELRYBOX = 3001 //客户端穿戴首饰盒
	CM_TAKEONZODIACSIGN    = 3002 //客户端穿戴十二生肖物品

	CM_JEWELRYBOXITEMTOBAG = 3003 //客户端首饰盒物品脱下
	CM_ZODIACSIGNTOBAG     = 3004 //客户端十二生肖物品到背包

	//动作命令1
	CM_TURN     = 3010 //转身(方向改变)
	CM_WALK     = 3011 //走
	CM_SITDOWN  = 3012 //挖(蹲下)
	CM_RUN      = 3013 //跑
	CM_HIT      = 3014 //普通物理近身攻击
	CM_HEAVYHIT = 3015 //跳起来打的动作
	CM_HOTCLICK = 3700
	CM_HELP     = 3701

	RM_OPENPULSE_OK          = 10097
	RM_OPENPULSE_FAIL        = 10098
	RM_RUSHPULSE_OK          = 10099
	RM_RUSHPULSE_FAIL        = 10100
	RM_PULSECHANGED          = 10101
	RM_BATTERORDER           = 10102
	RM_BATTERUSEFINALLY      = 10103
	RM_HEROBATTERORDER       = 10104
	RM_HEROPULSECHANGED      = 10105
	RM_STORMSRATE            = 10106
	RM_STORMSRATECHANGED     = 10107
	RM_HEROSTORMSRATECHANGED = 10108
	RM_OPENPULSENEEDLEV      = 10109
	//双英雄 相关
	//  RM_GETDOUBLEHEROINFO     = 36013
	RM_HEROATTECTMODE    = 10110
	RM_GETASSESSHEROINFO = 10111
	RM_QUERYASSESSHERO   = 10112
	RM_SHOWASSESSDLG     = 10113
	RM_ISDEHERO          = 10114
	RM_OPENTRAININGDLG   = 10115

	CM_SAY         = 3030 //角色发言
	CM_40HIT       = 3026
	CM_41HIT       = 3027
	CM_42HIT       = 3029
	CM_43HIT       = 3028
	CM_USEBATTER   = 3080  //使用连击
	RM_SLAVECREATE = 10116 //召唤宝宝

	RM_MENU_OK              = 10117 //菜单
	RM_MERCHANTDLGCLOSE     = 10118
	RM_SENDDELITEMLIST      = 10119 //发送删除项目的名单
	RM_SENDUSERSREPAIR      = 10120 //发送用户修理
	RM_SENDGOODSLIST        = 10121 //发送商品名单
	RM_SENDUSERSELL         = 10122 //发送用户出售
	RM_SENDUSERREPAIR       = 10123 //发送用户修理
	RM_USERMAKEDRUGITEMLIST = 10124 //用户做药品项目的名单
	RM_USERSTORAGEITEM      = 10125 //用户仓库项目
	RM_USERGETBACKITEM      = 10126 //用户获得回的仓库项目

	RM_SPACEMOVE_FIRE2 = 10127 //空间移动
	RM_SPACEMOVE_FIRE  = 10128 //空间移动

	RM_BUYITEM_SUCCESS     = 10129 //购买项目成功
	RM_BUYITEM_FAIL        = 10130 //购买项目失败
	RM_SENDDETAILGOODSLIST = 10131 //发送详细的商品名单
	RM_SENDBUYPRICE        = 10132 //发送购买价格
	RM_USERSELLITEM_OK     = 10133 //用户出售成功
	RM_USERSELLITEM_FAIL   = 10134 //用户出售失败
	RM_MAKEDRUG_SUCCESS    = 10135 //做药成功
	RM_MAKEDRUG_FAIL       = 10136 //做药失败
	RM_SENDREPAIRCOST      = 10137 //发送修理成本
	RM_USERREPAIRITEM_OK   = 10138 //用户修理项目成功
	RM_USERREPAIRITEM_FAIL = 10139 //用户修理项目失败

	MAXBAGITEM      = 54 //人物背包最大数量
	MAXSTOREITEM    = 46
	RM_10155        = 10140
	RM_ADJUST_BONUS = 10141

	RM_BUILDGUILD_OK   = 10142
	RM_BUILDGUILD_FAIL = 10143
	RM_DONATE_OK       = 10144

	RM_MYSTATUS = 10145

	CM_QUERYMENUSTATE  = 4004
	SM_QUERYMENUSTATE  = 4041
	CM_EXECMENUITEM    = 4005
	SM_GUILD_ERROR     = 4042
	SM_GUILD_INVITE    = 4043
	CM_GUILD_AGREE     = 4006
	CM_GUILD_REJECT    = 4007
	SM_GUILD_REQ       = 4044
	CM_GUILD_REQAGREE  = 4008
	CM_GUILD_REQREJECT = 4009
	SM_GROUP_INVITE    = 4045
	CM_GROUP_AGREE     = 4010
	CM_GROUP_REJECT    = 4011
	SM_GROUP_REQ       = 4046
	CM_GROUP_REQAGREE  = 4012
	CM_GROUP_REQREJECT = 4013
	SM_GROUP_ERROR     = 4047
	CM_QUERYUSERSTATE  = 82 //查询用户状态(用户登录进去,实际上是客户端向服务器索取查询最近一次,退出服务器前的状态的过程,
	//服务器自动把用户最近一次下线以让游戏继续的一些信息返回到客户端)

	CM_QUERYBAGITEMS = 81 //查询包裹物品
	CM_QUERYUSEITEMS = 83

	CM_QUERYUSERSET = 4014

	CM_OPENDOOR       = 1002 //开门,人物走到地图的某个过门点时
	CM_SOFTCLOSE      = 1009 //游戏小退
	CM_SOFTEXIT       = 4015 //游戏大退
	CM_1042           = 1042
	CM_GUILDALLY      = 1044
	CM_GUILDBREAKALLY = 1045

	RM_HORSERUN = 10146
	RM_HEAVYHIT = 10147
	RM_BIGHIT   = 10148
	RM_MOVEFAIL = 10149
	RM_RUSHKUNG = 10150

	RM_MAGICFIREFAIL  = 10151
	RM_LAMPCHANGEDURA = 10152
	RM_GROUPCANCEL    = 10153

	RM_DONATE_FAIL = 10154

	RM_BREAKWEAPON = 10155

	RM_PASSWORD = 10156

	RM_PASSWORDSTATUS = 10157

	SM_40 = 35
	SM_41 = 36
	SM_42 = 37
	SM_43 = 38
	SM_44 = 39 //龙影剑法

	SM_716 = 716

	SM_PASSWORD       = 3030
	SM_PASSWORDSTATUS = 4048

	SM_GAMEGOLDNAME = 55 //向客户端发送游戏币,游戏点,金刚石,灵符数量

	SM_SERVERCONFIG = 4049
	SM_GETREGINFO   = 4050

	ET_DIGOUTZOMBI = 1
	ET_PILESTONES  = 3
	ET_HOLYCURTAIN = 4
	ET_FIRE        = 5
	ET_SCULPEICE   = 6
	//{6种烟花}
	ET_FIREFLOWER_1    = 7 //一心一意
	ET_FIREFLOWER_2    = 8 //心心相印
	ET_FIREFLOWER_3    = 9
	ET_FIREFLOWER_4    = 10
	ET_FIREFLOWER_5    = 11
	ET_FIREFLOWER_6    = 12
	ET_FIREFLOWER_7    = 13
	ET_FIREFLOWER_8    = 14 //没有图片
	ET_FOUNTAIN        = 15 //喷泉效果 20080624
	ET_DIEEVENT        = 16 //人型庄主死亡动画效果 20080918
	ET_FIREDRAGON      = 17 //守护兽小火圈效果 20090123
	ET_HEROLOGOUT      = 18
	ET_SOULETRAP       = 19
	ET_SOULETRAPLOCKED = 20
	ET_MAGICEVENT      = 21 //自定义技能渲染外观
	ET_SKILLEVENT      = 22 //自定义技能 随云

	CM_PROTOCOL       = 2000
	CM_IDPASSWORD     = 2001 //客户端向服务器发送ID和密码
	CM_ADDNEWUSER     = 2002 //新建用户,就是注册新账号,登录时选择了"新用户"并操作成功
	CM_CHANGEPASSWORD = 2003 //修改密码
	CM_UPDATEUSER     = 2004 //更新注册资料??
	CM_RANDOMCODE     = 2006 //取验证码 20080612
	SM_RANDOMCODE     = 2007

	CLIENT_VERSION_NUMBER = 920151212 //9+客户端版本号
	CM_3037               = 3039      //2007.10.15改了值  以前是  3037

	SM_NEEDPASSWORD = 4051
	CM_POWERBLOCK   = 0

	//商铺相关
	CM_OPENSHOP                = 4016 //打开商铺
	SM_SENGSHOPITEMS           = 4052 // SERIES 7 每页的数量    wParam 总页数
	CM_BUYSHOPITEM             = 4017
	SM_BUYSHOPITEM_SUCCESS     = 4053
	SM_BUYSHOPITEM_FAIL        = 4054
	SM_SENGSHOPSPECIALLYITEMS  = 4055 //奇珍类型
	SM_BUYSHOPITEMGIVE_FAIL    = 4056
	SM_BUYSHOPITEMGIVE_SUCCESS = 4057

	RM_OPENSHOPSpecially       = 10158
	RM_OPENSHOP                = 10159
	RM_BUYSHOPITEM_FAIL        = 10160 //商铺购买物品失败
	RM_BUYSHOPITEMGIVE_FAIL    = 10161
	RM_BUYSHOPITEMGIVE_SUCCESS = 10162
	//==============================新增物品寄售系统(拍卖)==========================
	RM_SENDSELLOFFGOODSLIST     = 10163 //拍卖
	SM_SENDSELLOFFGOODSLIST     = 4058  //拍卖
	RM_SENDUSERSELLOFFITEM      = 10164 //寄售物品
	SM_SENDUSERSELLOFFITEM      = 4059  //寄售物品
	RM_SENDSELLOFFITEMLIST      = 10165 //查询得到的寄售物品
	CM_SENDSELLOFFITEMLIST      = 4018  //查询得到的寄售物品
	RM_SENDBUYSELLOFFITEM_OK    = 10166 //购买寄售物品成功
	SM_SENDBUYSELLOFFITEM_OK    = 4060  //购买寄售物品成功
	RM_SENDBUYSELLOFFITEM_FAIL  = 10167 //购买寄售物品失败
	SM_SENDBUYSELLOFFITEM_FAIL  = 4061  //购买寄售物品失败
	RM_SENDBUYSELLOFFITEM       = 10168 //购买选择寄售物品
	CM_SENDBUYSELLOFFITEM       = 4019  //购买选择寄售物品
	RM_SENDQUERYSELLOFFITEM     = 10169 //查询选择寄售物品
	CM_SENDQUERYSELLOFFITEM     = 4020  //查询选择寄售物品
	RM_SENDSELLOFFITEM          = 10170 //接受寄售物品
	CM_SENDSELLOFFITEM          = 4021  //接受寄售物品
	RM_SENDUSERSELLOFFITEM_FAIL = 10171 //R = -3  寄售物品失败
	RM_SENDUSERSELLOFFITEM_OK   = 10172 //寄售物品成功
	SM_SENDUSERSELLOFFITEM_FAIL = 4062  //R = -3  寄售物品失败
	SM_SENDUSERSELLOFFITEM_OK   = 4063  //寄售物品成功
	//==============================元宝寄售系统(20080316)==========================
	RM_SENDDEALOFFFORM      = 10173 //打开出售物品窗口
	SM_SENDDEALOFFFORM      = 4064  //打开出售物品窗口
	CM_SELLOFFADDITEM       = 4022  //客户端往出售物品窗口里加物品
	SM_SELLOFFADDITEM_OK    = 4065  //客户端往出售物品窗口里加物品 成功
	RM_SELLOFFADDITEM_OK    = 10174
	SM_SellOffADDITEM_FAIL  = 4066 //客户端往出售物品窗口里加物品 失败
	RM_SellOffADDITEM_FAIL  = 10175
	CM_SELLOFFDELITEM       = 4023 //客户端删除出售物品窗里的物品
	SM_SELLOFFDELITEM_OK    = 4067 //客户端删除出售物品窗里的物品 成功
	RM_SELLOFFDELITEM_OK    = 10176
	SM_SELLOFFDELITEM_FAIL  = 4068 //客户端删除出售物品窗里的物品 失败
	RM_SELLOFFDELITEM_FAIL  = 10177
	CM_SELLOFFCANCEL        = 4024  //客户端取消元宝寄售
	RM_SELLOFFCANCEL        = 10178 // 元宝寄售取消出售
	SM_SellOffCANCEL        = 4069  //元宝寄售取消出售
	CM_SELLOFFEND           = 4025  //客户端元宝寄售结束
	SM_SELLOFFEND_OK        = 4070  //客户端元宝寄售结束 成功
	RM_SELLOFFEND_OK        = 10179
	SM_SELLOFFEND_FAIL      = 4071 //客户端元宝寄售结束 失败
	RM_SELLOFFEND_FAIL      = 10180
	RM_QUERYYBSELL          = 10181 //查询正在出售的物品
	SM_QUERYYBSELL          = 4072  //查询正在出售的物品
	RM_QUERYYBDEAL          = 10182 //查询可以的购买物品
	SM_QUERYYBDEAL          = 4073  //查询可以的购买物品
	CM_CANCELSELLOFFITEMING = 4026  //取消正在寄售的物品 20080318(出售人)
	//SM_CANCELSELLOFFITEMING_OK =23018//取消正在寄售的物品 成功
	CM_SELLOFFBUYCANCEL = 4027 //取消寄售 物品购买 20080318(购买人)
	CM_SELLOFFBUY       = 4028 //确定购买寄售物品 20080318
	SM_SELLOFFBUY_OK    = 4074 //购买成功
	RM_SELLOFFBUY_OK    = 10183
	//SM_SELLOFFBUY_FAIL =23029//购买失败
	//RM_SELLOFFBUY_FAIL =23030
	//==============================================================================
	//英雄
	////////////////////////////////////////////////////////////////////////////////
	CM_RECALLHERO = 4029 //召唤英雄
	SM_RECALLHERO = 4075
	SM_CREATEHERO = 4076 //创建英雄

	SM_HERODEATH                   = 4077 //创建死亡
	CM_HEROCHGSTATUS               = 4030 //改变英雄状态
	CM_HEROATTACKTARGET            = 4031 //英雄锁定目标
	CM_HEROPROTECT                 = 4032 //守护目标
	CM_HEROTAKEONITEM              = 4033 //打开物品栏
	CM_HEROTAKEOFFITEM             = 4034 //关闭物品栏
	CM_TAKEOFFITEMHEROBAG          = 4035 //装备脱下到英雄包裹
	CM_TAKEOFFITEMTOMASTERBAG      = 4036 //装备脱下到主人包裹
	CM_SENDITEMTOMASTERBAG         = 4037 //主人包裹到英雄包裹
	CM_SENDITEMTOHEROBAG           = 4038 //英雄包裹到主人包裹
	SM_HEROTAKEON_OK               = 4078
	SM_HEROTAKEON_FAIL             = 4079
	SM_HEROTAKEOFF_OK              = 4080
	SM_HEROTAKEOFF_FAIL            = 4081
	SM_TAKEOFFTOHEROBAG_OK         = 4082
	SM_TAKEOFFTOHEROBAG_FAIL       = 4083
	SM_TAKEOFFTOMASTERBAG_OK       = 4084
	SM_TAKEOFFTOMASTERBAG_FAIL     = 4085
	CM_HEROTAKEONITEMFORMMASTERBAG = 4039 //从主人包裹穿装备到英雄包裹
	CM_TAKEONITEMFORMHEROBAG       = 4040 //从英雄包裹穿装备到主人包裹
	SM_SENDITEMTOMASTERBAG_OK      = 4086 //主人包裹到英雄包裹成功
	SM_SENDITEMTOMASTERBAG_FAIL    = 4087 //主人包裹到英雄包裹失败
	SM_SENDITEMTOHEROBAG_OK        = 4088 //英雄包裹到主人包裹
	SM_SENDITEMTOHEROBAG_FAIL      = 4089 //英雄包裹到主人包裹
	CM_QUERYHEROBAGCOUNT           = 4041 //查看英雄包裹容量
	SM_QUERYHEROBAGCOUNT           = 4090 //查看英雄包裹容量
	CM_QUERYHEROBAGITEMS           = 4042 //查看英雄包裹
	SM_SENDHEROUSEITEMS            = 4091 //发送英雄身上装备
	SM_HEROBAGITEMS                = 4092 //接收英雄物品
	SM_HEROADDITEM                 = 4093 //英雄包裹添加物品
	SM_HERODELITEM                 = 4094 //英雄包裹删除物品
	SM_HEROUPDATEITEM              = 4095 //英雄包裹更新物品
	SM_HEROADDMAGIC                = 4096 //添加英雄魔法
	SM_HEROSENDMYMAGIC             = 4097 //发送英雄的魔法
	SM_HERODELMAGIC                = 4098 //删除英雄魔法
	SM_HEROABILITY                 = 4099 //英雄属性1
	SM_HEROSUBABILITY              = 4100 //英雄属性2
	SM_HEROWEIGHTCHANGED           = 4101
	CM_HEROEAT                     = 4043 //吃东西
	SM_HEROEAT_OK                  = 4102 //吃东西成功
	SM_HEROEAT_FAIL                = 4103 //吃东西失败
	SM_HEROMAGIC_LVEXP             = 4104 //魔法等级
	SM_HERODURACHANGE              = 4105 //英雄持久改变
	SM_HEROWINEXP                  = 4106 //英雄增加经验
	SM_HEROLEVELUP                 = 4107 //英雄升级
	SM_HEROCHANGEITEM              = 4108 //好象没用上？
	SM_HERODELITEMS                = 4109 //删除英雄物品
	CM_HERODROPITEM                = 4044 //英雄往地上扔物品
	SM_HERODROPITEM_SUCCESS        = 4110 //英雄扔物品成功
	SM_HERODROPITEM_FAIL           = 4111 //英雄扔物品失败
	CM_HEROGOTETHERUSESPELL        = 4045 //使用合击
	SM_GOTETHERUSESPELL            = 4112 //使用合击
	SM_FIRDRAGONPOINT              = 4113 //英雄怒气值
	CM_REPAIRFIRDRAGON             = 4046 //修补火龙之心
	SM_REPAIRFIRDRAGON_OK          = 4114 //修补火龙之心成功
	SM_REPAIRFIRDRAGON_FAIL        = 4115 //修补火龙之心失败
	//---------------------------------------------------
	//祝福罐.魔令包功能 20080102
	CM_REPAIRDRAGON      = 4047 //祝福罐.魔令包功能
	SM_REPAIRDRAGON_OK   = 4116 //修补祝福罐.魔令包成功
	SM_REPAIRDRAGON_FAIL = 4117 //修补祝福罐.魔令包失败
	//----------------------------------------------------

	//-------连击 经脉----- /
	CM_OPENPULSE            = 4048
	CM_RUSHPULSE            = 4049
	CM_QUERYOPENPULSE       = 4050
	CM_SETBATTERORDER       = 4051
	CM_SETHEROBATTERORDER   = 4052
	CM_QUERYHEROOPENPULSE   = 4053
	CM_RUSHHEROPULSE        = 4054
	CM_CHANGEHEROATTECTMODE = 4055 //改变副将英雄攻击模式
	CM_QUERYASSESSHERO      = 4056
	CM_ASSESSMENTHERO       = 4057
	CM_TRAININGHERO         = 4058

	RM_RECALLHERO        = 10184 //召唤英雄
	RM_HEROWEIGHTCHANGED = 10185
	RM_SENDHEROUSEITEMS  = 10186
	RM_SENDHEROMYMAGIC   = 10187
	RM_HEROMAGIC_LVEXP   = 10188
	RM_QUERYHEROBAGCOUNT = 10189
	RM_HEROABILITY       = 10190
	RM_HERODURACHANGE    = 10191
	RM_HERODEATH         = 10192
	RM_HEROLEVELUP       = 10193
	RM_HEROWINEXP        = 10194
	//RM_HEROLOGOUT = 20010
	RM_CREATEHERO     = 10195
	RM_MAKEGHOSTHERO  = 10196
	RM_HEROSUBABILITY = 10197

	RM_GOTETHERUSESPELL = 10198 //使用合击
	RM_FIRDRAGONPOINT   = 10199 //发送英雄怒气值
	RM_CHANGETURN       = 10200
	//-----------------------------------月灵重击
	RM_FAIRYATTACKRATE = 10201
	SM_FAIRYATTACKRATE = 4118
	//-----------------------------------
	RM_DESTROYHERO = 10202 //英雄销毁
	SM_DESTROYHERO = 4119  //英雄销毁

	ET_PROTECTION_STRUCK = 20022 //护体受攻击  20080108
	ET_PROTECTION_PIP    = 20023 //护体被破

	SM_MYSHOW = 4120  //显示自身动画
	RM_MYSHOW = 10203 //

	RM_OPENBOXS       = 10204 //打开宝箱 20080115
	SM_OPENBOXS       = 4121  //打开宝箱 20080115
	RM_OPENShuffle    = 10205 //打开洗牌
	SM_OPENShuffle    = 4122  //打开洗牌
	CM_GETShuffleItem = 4059  //

	CM_OPENBOXS                 = 4060  //打开宝箱 20080115 清清加
	CM_MOVEBOXS                 = 4061  //转动宝箱 20080117
	RM_MOVEBOXS                 = 10206 //转动宝箱 20080117
	SM_MOVEBOXS                 = 4123  //转动宝箱 20080117
	CM_GETBOXS                  = 4062  //客户端取得宝箱物品 20080118
	SM_GETBOXS                  = 4124
	RM_GETBOXS                  = 10207
	SM_OPENBOOKS                = 4125 //打开卧龙NPC 20080119
	RM_OPENBOOKS                = 10208
	RM_DRAGONPOINT              = 10209 //发送黄条气值 20080201
	SM_DRAGONPOINT              = 4126
	ET_OBJECTLEVELUP            = 20038 //人物升级动画显示 20080222
	RM_CHANGEATTATCKMODE        = 10210 //改变攻击模式 20080228
	SM_CHANGEATTATCKMODE        = 4127  //改变攻击模式 20080228
	CM_EXCHANGEGAMEGIRD         = 4063  //商铺兑换灵符  20080302
	SM_EXCHANGEGAMEGIRD_FAIL    = 4128  //商铺购买物品失败
	SM_EXCHANGEGAMEGIRD_SUCCESS = 4129
	RM_EXCHANGEGAMEGIRD_FAIL    = 10211
	RM_EXCHANGEGAMEGIRD_SUCCESS = 10212
	RM_OPENDRAGONBOXS           = 10213 //卧龙开宝箱 20080306
	SM_OPENDRAGONBOXS           = 4130  //卧龙开宝箱 20080306
	// SM_OPENBOXS_OK = 20047 //打开宝箱成功 20080306
	RM_OPENBOXS_FAIL = 10214 //打开宝箱失败 20080306
	SM_OPENBOXS_FAIL = 4131  //打开宝箱失败 20080306

	RM_EXPTIMEITEMS = 10215 //聚灵珠 发送时间改变消息 20080306
	SM_EXPTIMEITEMS = 4132  //聚灵珠 发送时间改变消息 20080306

	ET_OBJECTBUTCHMON = 20053 //人物挖尸体得到物品显示 20080325
	ET_DRINKDECDRAGON = 20054 //喝酒抵御合击，显示自身效果 20090105

	//SM_CLOSEDRAGONPOINT = 20055  //关闭龙影黄条 20080329
	//---------------------------粹练系统------------------------------------------
	RM_QUERYREFINEITEM = 10216 //打开粹练框口
	SM_QUERYREFINEITEM = 4133  //打开粹练框口
	CM_REFINEITEM      = 4064  //客户端发送粹练物品 20080507

	SM_UPDATERYREFINEITEM  = 4134 //更新粹练物品 20080507
	CM_REPAIRFINEITEM      = 4065 //修补火云石 20080507 20080507
	SM_REPAIRFINEITEM_OK   = 4135 //修补火云石成功  20080507
	SM_REPAIRFINEITEM_FAIL = 4136 //修补火云石失败  20080507
	//-----------------------------------------------------------------------------

	RM_GETHEROINFO            = 10217
	SM_GETHEROINFO            = 4137  //获得英雄数据
	CM_SELGETHERO             = 4066  //取出英雄
	RM_SENDUSERPLAYDRINK      = 10218 //出现请酒对话框 20080515
	SM_SENDUSERPLAYDRINK      = 4138  //出现请酒对话框 20080515
	CM_USERPLAYDRINKITEM      = 4067  //请酒框放上物品发送到M2
	SM_USERPLAYDRINK_OK       = 4139  //请酒成功  20080515
	SM_USERPLAYDRINK_FAIL     = 4140  //请酒失败 20080515
	RM_PLAYDRINKSAY           = 10219 //
	SM_PLAYDRINKSAY           = 4141
	CM_PlAYDRINKDLGSELECT     = 4068  //商品选择,大类
	RM_OPENPLAYDRINK          = 10220 //打开窗口
	SM_OPENPLAYDRINK          = 4142  //打开窗口
	CM_PlAYDRINKGAME          = 4069  //发送猜拳码数 20080517
	RM_PlayDrinkToDrink       = 10221 //发送到客户端谁赢谁输
	SM_PlayDrinkToDrink       = 4143  //
	CM_DrinkUpdateValue       = 4070  //发送喝酒
	RM_DrinkUpdateValue       = 10222 //返回喝酒
	SM_DrinkUpdateValue       = 4144  //返回喝酒
	RM_CLOSEDRINK             = 10223 //关闭斗酒，请酒窗口
	SM_CLOSEDRINK             = 4145  //关闭斗酒，请酒窗口
	CM_USERPLAYDRINK          = 4071  //客户端发送请酒物品
	SM_USERPLAYDRINKITEM_OK   = 4146  //请酒物品成功
	SM_USERPLAYDRINKITEM_FAIL = 4147  //请酒物品失败
	RM_OPENURL                = 10224 //连接指定网站
	SM_OPENURL                = 4148

	RM_PIXINGHIT = 10225 //劈星斩效果 20080611
	SM_PIXINGHIT = 4149

	RM_LEITINGHIT = 10226 //雷霆一击效果 20080611
	SM_LEITINGHIT = 4150

	CM_CHECKNUM       = 4072 //检测验证码 20080612
	SM_CHECKNUM_OK    = 4151
	CM_CHANGECHECKNUM = 4073

	RM_AUTOGOTOXY = 10227 //自动寻路
	SM_AUTOGOTOXY = 4152
	//-----------------------酿酒系统---------------------------------------------
	RM_OPENMAKEWINE    = 10228 //打开酿酒窗口
	SM_OPENMAKEWINE    = 4153  //打开酿酒窗口
	CM_BEGINMAKEWINE   = 4074  //开始酿酒(即把材料全放上窗口)
	RM_MAKEWINE_OK     = 10229 //酿酒成功
	SM_MAKEWINE_OK     = 4154  //酿酒成功
	RM_MAKEWINE_FAIL   = 10230 //酿酒失败
	SM_MAKEWINE_FAIL   = 4155  //酿酒失败
	RM_NPCWALK         = 10231 //酿酒NPC走动
	SM_NPCWALK         = 4156  //酿酒NPC走动
	RM_MAGIC68SKILLEXP = 10232 //酒气护体技能经验
	SM_MAGIC68SKILLEXP = 4157  //酒气护体技能经验
	//------------------------挑战系统--------------------------------------------
	SM_CHALLENGE_FAIL = 4158 //挑战失败
	SM_CHALLENGEMENU  = 4159 //打开挑战抵押物品窗口
	CM_CHALLENGETRY   = 4075 //玩家点挑战

	CM_CHALLENGEADDITEM       = 4076 //玩家把物品放到挑战框中
	SM_CHALLENGEADDITEM_OK    = 4160 //玩家增加抵押物品成功
	SM_CHALLENGEADDITEM_FAIL  = 4161 //玩家增加抵押物品失败
	SM_CHALLENGEREMOTEADDITEM = 4162 //发送增加抵押的物品后,给客户端显示

	CM_CHALLENGEDELITEM       = 4077 //玩家从挑战框中取回物品
	SM_CHALLENGEDELITEM_OK    = 4163 //玩家删除抵押物品成功
	SM_CHALLENGEDELITEM_FAIL  = 4164 //玩家删除抵押物品失败
	SM_CHALLENGEREMOTEDELITEM = 4165 //发送删除抵押的物品后,给客户端显示

	CM_CHALLENGECANCEL = 4078 //玩家取消挑战
	SM_CHALLENGECANCEL = 4166 //玩家取消挑战

	CM_CHALLENGECHGGOLD     = 4079 //客户端把金币放到挑战框中
	SM_CHALLENCHGGOLD_FAIL  = 4167 //客户端把金币放到挑战框中失败
	SM_CHALLENCHGGOLD_OK    = 4168 //客户端把金币放到挑战框中成功
	SM_CHALLENREMOTECHGGOLD = 4169 //客户端把金币放到挑战框中,给客户端显示

	CM_CHALLENGECHGDIAMOND     = 4080 //客户端把金刚石放到挑战框中
	SM_CHALLENCHGDIAMOND_FAIL  = 4170 //客户端把金刚石放到挑战框中失败
	SM_CHALLENCHGDIAMOND_OK    = 4171 //客户端把金刚石放到挑战框中成功
	SM_CHALLENREMOTECHGDIAMOND = 4172 //客户端把金刚石放到挑战框中,给客户端显示

	CM_CHALLENGEEND   = 4081 //挑战抵押物品结束
	SM_CLOSECHALLENGE = 4173 //关闭挑战抵押物品窗口
	//----------------------------------------------------------------------------
	RM_PLAYMAKEWINEABILITY = 10233 //酒2相关属性 20080804
	SM_PLAYMAKEWINEABILITY = 4174  //酒2相关属性 20080804
	RM_HEROMAKEWINEABILITY = 10234 //酒2相关属性 20080804
	SM_HEROMAKEWINEABILITY = 4175  //酒2相关属性 20080804

	RM_CANEXPLORATION = 10235 //可探索 20080810
	SM_CANEXPLORATION = 4176  //可探索 20080810
	//----------------------------------------------------------------------------
	SM_SENDLOGINKEY  = 4177 //网关给客户端或登陆器发送登陆器封包码 20080901
	SM_GATEPASS_FAIL = 4178 //和网关的密码错误

	RM_HEROMAGIC68SKILLEXP = 10236 //英雄酒气护体技能经验  20080925
	SM_HEROMAGIC68SKILLEXP = 4179  //英雄酒气护体技能经验  20080925

	RM_USERBIGSTORAGEITEM = 10237 //用户无限仓库项目
	RM_USERBIGGETBACKITEM = 10238 //用户获得回的无限仓库项目

	RM_HEROAUTOOPENDEFENCE = 10239 //英雄内挂自动持续开盾 20080930
	SM_HEROAUTOOPENDEFENCE = 4180  //英雄内挂自动持续开盾 20080930
	CM_HEROAUTOOPENDEFENCE = 4082  //英雄内挂自动持续开盾 20080930

	RM_MAGIC69SKILLEXP     = 10240 //内功心法经验
	SM_MAGIC69SKILLEXP     = 4181  //内功心法经验
	RM_HEROMAGIC69SKILLEXP = 10241 //英雄内功心法经验  20080930
	SM_HEROMAGIC69SKILLEXP = 4182  //英雄内功心法经验  20080930

	RM_MAGIC69SKILLNH = 10242 //内力值(黄条) 20081002
	SM_MAGIC69SKILLNH = 4183  //内力值(黄条) 20081002
	RM_WINNHEXP       = 10243 //取得内功经验 20081007
	SM_WINNHEXP       = 4184  //取得内功经验 20081007
	RM_HEROWINNHEXP   = 10244 //英雄取得内功经验 20081007
	SM_HEROWINNHEXP   = 4185  //英雄取得内功经验 20081007

	SM_SHOWSIGHICON  = 4186 //显示感叹号图标 20090126
	CM_CLICKSIGHICON = 4083 //点击感叹号图标 20090126
	SM_UPDATETIME    = 4187 //统一与客户端的倒计时 20090129

	RM_OPENEXPCRYSTAL     = 10245 //显示天地结晶图标 20090201
	SM_OPENEXPCRYSTAL     = 4188  //显示天地结晶图标 20090201
	SM_SENDCRYSTALNGEXP   = 4189  //发送天地结晶的内功经验 20090201
	SM_SENDCRYSTALEXP     = 4190  //发送天地结晶的内功经验 20090201
	SM_SENDCRYSTALLEVEL   = 4191  //发送天地结晶的等级 20090202
	CM_CLICKCRYSTALEXPTOP = 4084  //点击天地结晶获得经验 20090202
	SM_ZHUIXIN_OK         = 4192  //追心刺

	RM_MEMBERINFO         = 10246 //会员信息
	SM_MEMBERINFO         = 4193
	CM_OPENMEMBER         = 4085
	CM_SENDITEMClickFunc  = 4086
	RM_AfterITEMClickFunc = 10247
	RM_MERCHANTSAYCUSTOM  = 10248
	SM_MERCHANTSAYCUSTOM  = 4194
	RM_Effect             = 10249
	SM_Effect             = 4195
	CM_OPENPAYHOME        = 4087
	CM_SENDITEMUnite      = 4088 //合并物品
	CM_SENDITEMSPLIT      = 4089 //拆分物品

	CM_MARKET_COMMAND      = 4090 //客户端操作市场
	CM_MARKET_PUTON        = 1    //上架
	CM_MARKET_PUTOFF       = 2    //下架
	CM_MARKET_SETNAME      = 3    //改摊位名称
	CM_MARKET_UPDATE       = 4    //改售价
	CM_MARKET_BUY          = 5    //购买
	CM_MARKET_GETLIST      = 6    //刷新列表
	CM_MARKET_GETITEMS     = 7    //取摊位物品信息
	CM_MARKET_EXTRACT      = 8    //提取资金
	SM_MARKET_NAMELIST     = 4196 //摊位信息列表
	SM_MARKET_ITEMS        = 4197 //我的摊位物品信息
	SM_MARKET_WHOITEMS     = 4198 //他的摊位物品信息
	SM_MARKET_PUTON_OK     = 4199 //上架成功
	SM_MARKET_PUTON_FAILD  = 4200 //上架失败
	SM_MARKET_PUTOFF_OK    = 4201 //下架成功
	SM_MARKET_PUTOFF_FAILD = 4202 //下架失败
	SM_MARKET_BUY_OK       = 4203
	SM_MARKET_BUY_FAILD    = 4204
	SM_MARKET_UPDATE_OK    = 4205
	SM_MARKET_UPDATE_FAILD = 4206
	SM_MARKET_EXTRACT      = 4207

	//静态配置数据
	SM_CLIENTDATA_COMMAND           = 4208
	CD_CLIENTDATA_ITEMS             = 1
	CD_CLIENTDATA_ITEMSDSC          = 2
	CD_CLIENTDATA_SUITES            = 3
	CD_CLIENTDATA_MAP               = 4
	CD_CLIENTDATA_UI                = 5
	CD_CLIENTDATA_TYPENAMES         = 6
	CD_CLIENTDATA_MAGICS            = 7
	CD_CLIENTDATA_MAGICDESC         = 8
	CD_CLIENTDATA_VERCHECK          = 9
	CD_CLIENTDATA_ITEMWAY           = 10
	CD_CLIENTDATA_ACTORACTION       = 11
	CD_CLIENTDATA_SKILLCONFIG       = 12
	CD_CLIENTDATA_SKILLEFFECTCONFIG = 13

	CM_CLIENTDATA_VERSION = 4091
	SM_CLIENTEXTENDBUTTON = 4209  //客户端扩展按钮
	CM_CLIENTEXTENDBUTTON = 4092  //客户端扩展按钮执行功能
	RM_CLIENTACTION       = 10250 //让客户端执行某个动作
	SM_CLIENTACTION       = 4210  //让客户端执行某个动作
	CA_OPENBAG            = 0     //打开背包
	CA_OPENMAG            = 1     //打开技能
	CA_OPENMARKET         = 2     //打开市场
	CA_OPENSHOP           = 3     //打开商城
	CA_OPENMAILBOX        = 4     //打开信箱
	CA_RELOADBAG          = 5     //刷新包裹
	CA_OPENMAILVIEW       = 6     //打开M地图
	CA_OPENREFINEBOX      = 7     //打开淬炼窗口
	CA_OPENSTALL          = 8     //打开摆摊

	SM_SHOWPROGRESS  = 4211
	SM_CLOSEPROGRESS = 4212
	SM_RELATION      = 4213  //关系系统
	RM_RELATION      = 10251 //关系系统
	RT_ONLINE        = 1     //玩家上线
	RT_OFFLINE       = 2     //玩家下线
	RT_FirendsNames  = 3     //好友列表
	RT_EnemiesNames  = 4     //仇人列表
	RT_ADDFRIENDSUC  = 5     //添加好友成功
	RT_ADDENEMIYSUC  = 6     //添加黑名单成功
	RT_DELFIRENTSUC  = 7     //删除好友成功
	RT_DELENEMIYSUC  = 8     //删除黑名单成功
	RT_ADDFRIENDWAIT = 9     //添加好友等待回应
	RT_ADDFRIENDREQ  = 10    //收到好友添加请求
	RT_ADDFRIENDREJ  = 11    //否决了好友添加
	RT_PLAYNOTFOUND  = 12    //玩家不在线

	CM_RELATION               = 4093
	_CM_RELATION_ADDFIREND    = 0
	_CM_RELATION_ADDENEMIY    = 1
	_CM_RELATION_DELFRIEND    = 2
	_CM_RELATION_DELENEMIY    = 3
	_CM_RELATION_REJECTFIREND = 4
	_CM_RELATION_AGREEFIREND  = 5

	SM_CHANGESTATE = 4214
	CM_CHANGESTATE = 4094

	SM_CHANGESPEED = 4215

	SM_EMAIL                    = 4216
	CM_EMAIL                    = 4095
	RM_EMAIL                    = 10252
	_MAIL_UNREADCOUNT           = 0
	_MAIL_OPEN                  = 1
	_MAIL_LIST                  = 2
	_MAIL_DATA                  = 3
	_MAIL_NEW                   = 4
	_MAIL_DEL                   = 5
	_MAIL_DELALL                = 6
	_MAIL_DELSUC                = 7
	_MAIL_DELFAIL               = 8
	_MAIL_EXTRACT               = 9
	_MAIL_EXTRACTSUC            = 10
	_MAIL_EXTRACTFAIL           = 11
	_MAIL_EXTRACTFAIL_EXTRACTED = 1
	_MAIL_EXTRACTFAIL_BAGFULL   = 2
	_MAIL_EXTRACTFAIL_GOLDFULL  = 3
	_MAIL_EXTRACTFAIL_NOBUYGOLD = 4

	_MAIL_NEW_SUC            = 12
	_MAIL_NEW_FAIL           = 13
	_MAIL_NEW_FAIL_NOITEM    = 1
	_MAIL_NEW_FAIL_PRICE     = 2
	_MAIL_NEW_FAIL_ITEMNOEX  = 3
	_MAIL_NEW_FAIL_ITEMBIND  = 4
	_MAIL_NEW_FAIL_QF        = 5
	_MAIL_NEW_FAIL_GOLD      = 6
	_MAIL_NEW_FAIL_GAMEGOLD  = 7
	_MAIL_NEW_FAIL_GAMEPOINT = 8
	_MAIL_GOLDADD            = 14

	CM_SETMESSAGESTATE = 4096

	CM_COLLECT           = 4097
	RM_COLLECT           = 10253
	SM_COLLECT           = 4217
	_COLLECT_SUCCESS     = 1
	_COLLECT_SUSPEND     = 2
	_COLLECT_FINISH      = 3
	_COLLECT_FAILD_DEATH = 4
	_COLLECT_FAILD_SLOW  = 5

	CM_CTRLLBUTTONITEM = 4098

	RM_SPACEMOVE_SHOW3 = 10254
	SM_SPACEMOVE_SHOW3 = 4218
	RM_CLOSEWINDOW     = 10255
	SM_CLOSEWINDOW     = 4219

	CM_SNEAK          = 4099
	RM_SNEAK          = 10256
	SM_SNEAK          = 4220
	SM_SNEAKSTATE     = 4221
	SM_SKILLSTATE     = 4222
	SM_OPERATESTATE   = 4223
	RM_SKILLSTATE     = 10257
	_OP_FAIL          = 0
	_OP_SUC           = 1
	_OP_DIG           = 2
	SM_DISAPPEARIDS   = 4224
	RM_SENDMYTITLES   = 10258
	SM_SENDMYTITLES   = 4225
	RM_ADDTITLE       = 10259
	SM_ADDTITLE       = 4226
	RM_REMOVETITLE    = 10260
	SM_REMOVETITLE    = 4227
	RM_SETACTIVETITLE = 10261
	SM_SETACTIVETITLE = 4228
	CM_SETACTIVETITLE = 4100
	RM_FLASHWINDDOW   = 10262
	SM_FLASHWINDDOW   = 4229
	CM_OPENREFINEBOX  = 4101
	RM_QUESTION       = 10263
	SM_QUESTION       = 4230
	CM_QUERYSORT      = 4102
	RM_QUERYSORT      = 10264
	SM_QUERYSORT      = 4231

	RM_MISSIONLIST          = 10265
	SM_MISSIONLIST          = 4232
	RM_ADDMISSION           = 10266
	SM_ADDMISSION           = 4233
	RM_DELETEMISSION        = 10267
	SM_DELETEMISSION        = 4234
	RM_UPDATEMISSION        = 10268
	SM_UPDATEMISSION        = 4235
	RM_MISSIONLINKS         = 10269
	SM_MISSIONLINKS         = 4236
	RM_ADDMISSIONLINK       = 10270
	SM_ADDMISSIONLINK       = 4237
	RM_DELMISSIONLINK       = 10271
	SM_DELMISSIONLINK       = 4238
	CM_MISSIONCOMMANDSELECT = 4103

	SM_LOCKMOVEITEM      = 4239
	RM_SENDCUSTOMMESSAGE = 10272
	SM_SENDCUSTOMMESSAGE = 4240
	CM_CUSTOMMESSAGE     = 4104
	CM_RUNGATEHEARTBEAT  = 4105
	CM_PLAYDICE          = 4106
	RM_PLAYDICE          = 10273
	SM_PLAYDICE          = 4241

	RM_SendStall = 10274
	SM_SendStall = 4242

	CM_STALL_COMMAND         = 4107
	RM_STALL_COMMAND         = 10275
	SM_STALL_COMMAND         = 4243
	_STALL_Query             = 1
	_STALL_PutOn             = 2
	_STALL_PutOff            = 3
	_STALL_Update            = 4
	_STALL_PutOnBuy          = 5
	_STALL_PutOffBuy         = 6
	_STALL_UpdateBuy         = 7
	_STALL_Start             = 8
	_STALL_Stop              = 9
	_STALL_Buy               = 10
	_STALL_SaleToBuy         = 11
	_STALL_Message           = 12
	_STALL_GetBack           = 13
	SM_STALL_ITEMS           = 4244 //我的摊位物品信息
	SM_STALL_PUTON_OK        = 4245 //上架成功
	SM_STALL_PUTON_FAILD     = 4246 //上架失败
	SM_STALL_PUTOFF_OK       = 4247 //下架成功
	SM_STALL_PUTOFF_FAILD    = 4248 //下架失败
	SM_STALL_BUY_OK          = 4249
	SM_STALL_BUY_FAILD       = 4250
	SM_STALL_UPDATE_OK       = 4251
	SM_STALL_UPDATE_FAILD    = 4252
	SM_STALL_PUTONBUY_OK     = 4253
	SM_STALL_PUTONBUY_FAILD  = 4254
	SM_STALL_PUTOFFBUY_OK    = 4255
	SM_STALL_PUTOFFBUY_FAILD = 4256
	SM_STALL_STATEChanged    = 4257
	SM_STALL_EXTRACT         = 4258
	SM_STALL_AddMessageOK    = 4259
	SM_STALL_START_FAILD     = 4260
	SM_STALL_SALETOBUY_OK    = 4261
	SM_STALL_SALETOBUY_FAILD = 4262
	SM_STALL_UPDATED         = 4263
	SM_STALL_UPDATEBUY_OK    = 4264
	SM_STALL_UPDATEBUY_FAILD = 4265
	SM_STALL_GETBACK_INFOR   = 4266

	CM_SHORTCUT               = 4108
	CM_SIDEBUTTONCLICK        = 4109
	RM_SIDEBUTTONCommand      = 10276
	SM_SIDEBUTTONCommand      = 4267
	CM_GUILDEXTENDBUTTONCLICK = 4110
	RM_SETCHATTEXT            = 10277
	SM_SETCHATTEXT            = 4268
	//{ TODO -o随云 -c消息定义 : RM消息开始 }
	RM_CHECKBUFFSTATE   = 10278 //上线检查是否需要同步客户端BUFF状态
	RM_FUNCTIONSTATE    = 10279 //上线推送功能状态
	RM_CHARDESC         = 10280 //更新角色的角色描述
	RM_REALUSERNAME     = 10281 //真实的玩家角色名称
	RM_VIEWNEWACTOR     = 10282 //视野内出现新的玩家  随云 为了优化流量 和性能 重写 角色出现消息
	RM_CLOSEDICE        = 10283 //关闭骰子窗口
	RM_SERVERCONFIG     = 10284 //刷新服务器配置
	RM_BloodRushHit     = 10285 //血魄一击奔跑后攻击
	RM_BloodRush        = 10286 //血魄一击突进
	RM_STATECHANGEDELY  = 10287 //延迟广播 StateChange
	RM_TITLEEFFECTARRAY = 10288 //头顶特效数组
	RM_SKILLEVENTDAMAGE = 10289 //自定义技能的地图伤害事件
	RM_DELAYDAMAGE      = 10290 //延迟伤害

	//{ TODO -o随云 -c消息定义 : CM消息开始 }
	CM_Reserver              = 4111 //待用
	CM_ScriptButton          = 4112 //脚本按钮
	CM_CLIENTCHECKFLAG       = 4113 //客户端时间检查
	CM_CLICKGAMESHOP         = 4114 //点击商铺按钮
	CM_UNLOCKCLIENT_PASSWORD = 4115 //客户端在锁定的情况下提交密码

	//{ TODO -o随云 -c消息定义 : SM消息开始 }
	SM_FunctionState        = 4269 //同步功能状态
	SM_CHARDESC             = 4270
	SM_VIEWNEWACTOR         = 4271 //视野内出现角色
	SM_CLOSEDICE            = 4272 //关闭骰子窗口
	SM_BLOODRUSHHIT         = 4273 //刺客血魄一击 跑完 后攻击
	SM_BloodRush            = 4274 //血魄一击突进
	SM_StartBeCool          = 4275 //开始冷酷 客户端清空所有攻击栈
	SM_SetControlVisible    = 4276 //设置组件可视和不可视
	SM_CLIENTCHECKFLAG      = 4277 //客户端时间检查
	SM_CLIENTCHECKFLAG_FAIL = 4278 //客户端时间检查失败客户端做屏幕小退处理
	SM_LockClientState      = 4279 //锁定客户端
	SM_SETCLIENTUPPROPERTY  = 4280 //设置客户端UI属性。
	SM_LOADCLIENTUI         = 4281 //读取客户端UI文件
	SM_TITLEEFFECTARRAY     = 4282 //读取头顶特效数组
	SM_UpdateMagic          = 4283
	SM_OtherActorSkillSpell = 4284 //此消息 服务端不下发 用于客户端消息中转 随云。

	//ENDMESSAGE
	_SC_AttackMode     = 1 //@AttackMode
	_SC_Rest           = 2 //@Rest
	_SC_Alliance       = 3 //@Alliance
	_SC_CancelAlliance = 4 //@CancelAlliance 行会名称
	_SC_SwitchHorse    = 5 //骑马
	_SC_DEALTRY        = 6 //交易
	_SC_OPENGUILDDLG   = 7 //打开行会
	_SC_OPENSTALL      = 8 //打开自己的摊位
	_SC_UPDATESTALL    = 9

	CUSTOM_SKILL_EFFECT = 65535
	// 技能ID常量
	SKILL_ATTACK          = 0
	SKILL_FIREBALL        = 1  // 火球术
	SKILL_HEALLING        = 2  // 治愈术
	SKILL_ONESWORD        = 3  // 基本剑术
	SKILL_ILKWANG         = 4  // 精神力战法
	SKILL_FIREBALL2       = 5  //大火球
	SKILL_AMYOUNSUL       = 6  //施毒术
	SKILL_YEDO            = 7  // 攻杀
	SKILL_FIREWIND        = 8  //抗拒火环
	SKILL_FIRE            = 9  // 地狱火
	SKILL_SHOOTLIGHTEN    = 10 // 疾光电影
	SKILL_LIGHTENING      = 11 // 雷电术
	SKILL_ERGUM           = 12 // 刺杀剑术
	SKILL_FIRECHARM       = 13 // 灵魂火符
	SKILL_HANGMAJINBUB    = 14 //幽灵盾
	SKILL_DEJIWONHO       = 15 //神圣战甲术
	SKILL_HOLYSHIELD      = 16 // 困魔咒
	SKILL_SKELLETON       = 17 // 召唤骷髅
	SKILL_CLOAK           = 18 // 隐身术
	SKILL_BIGCLOAK        = 19 // 集体隐身术
	SKILL_TAMMING         = 20 //诱惑之光
	SKILL_SPACEMOVE       = 21 //瞬息移动
	SKILL_EARTHFIRE       = 22 // 火墙
	SKILL_FIREBOOM        = 23 // 爆裂火焰
	SKILL_LIGHTFLOWER     = 24 //地狱雷光
	SKILL_BANWOL          = 25 // 半月弯刀
	SKILL_FIRESWORD       = 26 // 烈火剑法
	SKILL_MOOTEBO         = 27 // 野蛮冲撞
	SKILL_SHOWHP          = 28 //心灵启示
	SKILL_BIGHEALLING     = 29 // 群体治疗术
	SKILL_SINSU           = 30 // 召唤神兽
	SKILL_SHIELD          = 31 // 魔法盾
	SKILL_KILLUNDEAD      = 32 //圣言术
	SKILL_SNOWWIND        = 33 // 冰咆哮
	SKILL_UNAMYOUNSUL     = 34 //解毒术
	SKILL_WINDTEBO        = 35 // 狮吼功
	SKILL_MABE            = 36 // 火焰冰
	SKILL_GROUPLIGHTENING = 37 // 群体雷电术
	SKILL_GROUPAMYOUNSUL  = 38 // 群体施毒术
	SKILL_GROUPDEDING     = 39 // 地钉
	SKILL_40              = 40 // 抱月刀法
	SKILL_41              = 41 // 狮子吼
	SKILL_42              = 42 // 开天斩
	SKILL_43              = 43 // 龙影剑法
	SKILL_44              = 44 // 寒冰掌
	SKILL_45              = 45 // 灭天火
	SKILL_46              = 46 // 分身术
	SKILL_47              = 47 // 火龙焰
	SKILL_48              = 48 // 气功波
	SKILL_49              = 49 // 净化术
	SKILL_50              = 50 // 无极真气
	SKILL_51              = 51 // 飓风破
	SKILL_52              = 52 // 诅咒术
	SKILL_53              = 53 // 血咒
	SKILL_54              = 54 // 骷髅咒
	SKILL_55              = 55 // 擒龙手
	SKILL_56              = 56 // 移行换位
	SKILL_57              = 57 // 复活术
	SKILL_58              = 58 // 流星火雨 20080510
	SKILL_59              = 59 // 噬血术 20080511
	SKILL_60              = 60 // 破魂斩
	SKILL_61              = 61 // 劈星斩
	SKILL_62              = 62 // 雷霆一击
	SKILL_63              = 63 // 噬魂沼泽
	SKILL_64              = 64 // 末日审判
	SKILL_65              = 65 // 火龙气焰
	SKILL_66              = 66 // 4级魔法盾 20080624
	SKILL_67              = 67 // 先天元力 20080625
	SKILL_68              = 68 // 酒气护体 20080625

	SKILL_69     = 69 // 倚天劈地  三职业
	SKILL_SACRED = 71 // 召唤圣兽
	SKILL_72     = 72 // 召唤月灵
	SKILL_73     = 73 // 道力盾  20080301
	SKILL_74     = 74 // 逐日剑法 20080511
	SKILL_75     = 75 // 护体神盾 20080107

	SKILL_76 = 76 // 三绝杀    战士
	SKILL_77 = 77 // 双龙破    法师
	SKILL_78 = 78 // 虎啸绝    道士
	SKILL_79 = 79 // 追心刺    战士
	SKILL_80 = 80 // 凤舞祭    法师
	SKILL_81 = 81 // 八卦掌    道士
	SKILL_82 = 82 // 断岳斩    战士
	SKILL_83 = 83 // 惊雷爆    法师
	SKILL_84 = 84 // 三焰咒    道士
	SKILL_85 = 85 // 横扫千军  战士
	SKILL_86 = 86 // 冰天雪地  法师
	SKILL_87 = 87 // 万剑归宗  道士
	SKILL_96 = 96 // 血魄一击(战)
	// SKILL_69 = 69 //
	// SKILL_69 = 69 //

	SKILL_88 = 88 // 四级基本剑术
	SKILL_89 = 89 // 四级刺杀
	SKILL_90 = 90 // 圆月弯刀
	SKILL_91 = 91 // 四级雷电
	SKILL_92 = 92 // 四级流星火雨
	SKILL_93 = 93 // 四级施毒术
	SKILL_94 = 94 // 四级噬血术

	//弓箭手
	SKILL_150 = 150 //基准箭术
	SKILL_151 = 151 //蓄势待发
	SKILL_152 = 152 //寒冰箭
	SKILL_153 = 153 //天罡震气
	SKILL_154 = 154 //排山倒海
	SKILL_155 = 155 //天神下凡
	SKILL_156 = 156 //恶魔降临
	SKILL_157 = 157 //万箭齐发
	SKILL_158 = 158 //雷光之眼
	//刺客
	SKILL_159 = 159 //暴击术
	SKILL_160 = 160 //潜行
	SKILL_161 = 161 //致残毒药
	SKILL_162 = 162 //精准术
	SKILL_163 = 163 //炎龙波
	SKILL_164 = 164 //旋风腿
	SKILL_165 = 165 //鬼灵步
	SKILL_166 = 166 //冷酷
	SKILL_167 = 167 //火镰狂舞
	SKILL_168 = 168 //血魄一击
	SKILL_169 = 169 //灵魂陷阱
	SKILL_170 = 170 //霜月
	//武僧
	SKILL_204 = 204 //罗汉棍法
	SKILL_205 = 205 //易筋经
	SKILL_206 = 206 //推山掌
	SKILL_207 = 207 //天雷阵
	SKILL_208 = 208 //降龙伏虎
	SKILL_209 = 209 //达摩棍法
	SKILL_210 = 210 //降魔棍法

	SKILL_200 = 200 // 怒之攻杀
	SKILL_201 = 201 // 静之攻杀
	Skill_202 = 202 // 怒之半月
	Skill_203 = 203 // 静之半月
	//  Skill_204 = 204 // 怒之烈火
	//  Skill_205 = 205 // 静之烈火
	//  Skill_206 = 206 // 怒之逐日
	//  Skill_207 = 207 // 静之逐日
	//  Skill_208 = 208 // 怒之火球
	//  Skill_209 = 209 // 静之火球
	//  Skill_210 = 210 // 怒之大火球
	Skill_211 = 211 // 静之大火球
	Skill_212 = 212 // 怒之火墙
	Skill_213 = 213 // 静之火墙
	Skill_214 = 214 // 怒之地狱火
	Skill_215 = 215 // 静之地狱火
	Skill_216 = 216 // 怒之疾光电影
	Skill_217 = 217 // 静之疾光电影
	Skill_218 = 218 // 怒之爆裂火焰
	Skill_219 = 219 // 静之爆裂火焰
	Skill_220 = 220 // 怒之冰咆哮
	Skill_221 = 221 // 静之冰咆哮
	Skill_222 = 222 // 怒之雷电
	Skill_223 = 223 // 静之雷电
	Skill_224 = 224 // 怒之地狱雷光
	Skill_225 = 225 // 静之地狱雷光
	Skill_226 = 226 // 怒之寒冰掌
	Skill_227 = 227 // 静之寒冰掌
	Skill_228 = 228 // 怒之灭天火
	Skill_229 = 229 // 静之灭天火
	Skill_230 = 230 // 怒之火符
	Skill_231 = 231 // 静之火符
	Skill_232 = 232 // 怒之噬血
	Skill_233 = 233 // 静之噬血
	Skill_234 = 234 // 怒之流星火雨
	Skill_235 = 235 // 静之流星火雨
	Skill_236 = 236 // 怒之内功剑法
	Skill_237 = 237 // 静之内功剑法

	////////////////////////////////////////////////////////////////////////////////
	UNITX = 48
	UNITY = 32
	HALFX = 24
	HALFY = 16
	//MAXBAGITEM = 46 //用户背包最大数量
	//MAXMAGIC = 30{20} //原来54 200081227 注释
	MAXSTORAGEITEM = 50
	LOGICALMAPUNIT = 40

	MsgHeaderSize  = 21
	ClientItemSize = 452
	StdItemSize    = 142
	UserItemSize   = 440

	//技能状态
	_MAGIC_STATE_FORBIDSPELL = 1 //禁止释放
	_MAGIC_STATE_LockTarget  = 2 //魔法锁定

	//战力计算属性保值
	_TA_ACMIN         = 0
	_TA_ACMAX         = 1
	_TA_MACMIN        = 2
	_TA_MACMAX        = 3
	_TA_DCMIN         = 4
	_TA_DCMAX         = 5
	_TA_MCMIN         = 6
	_TA_MCMAX         = 7
	_TA_SCMIN         = 8
	_TA_SCMAX         = 9
	_TA_TCMIN         = 10
	_TA_TCMAX         = 11
	_TA_PCMIN         = 12
	_TA_PCMAX         = 13
	_TA_WCMIN         = 30
	_TA_WCMAX         = 31
	_TA_Absorbing     = 14
	_TA_Rebound       = 15
	_TA_AttackAdd     = 16
	_TA_PunchHit      = 17
	_TA_CriticalHit   = 18
	_TA_HealthRecover = 19
	_TA_SpellRecover  = 20
	_TA_MaxHP         = 23
	_TA_MaxMP         = 24
	_TA_HitPoint      = 25
	_TA_SpeedPoint    = 26
	_TA_AntiMagic     = 27
	_TA_AntiPoison    = 28
	_TA_PoisonRecover = 29

	//自定义消息
	CM_GAMESESSION   = 8887 //服务器session验证
	CM_GAMELOGIN     = 8888
	SM_REQTOKEN      = 8889 //服务器token获取
	CM_RESTOKEN      = 8890 //客户端token应答
	CM_SELRESSESSION = 8891 //游戏网关发送session效验
)
