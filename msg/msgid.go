package msg

const (
	//SM_??    辑滚 -> 努扼捞攫飘肺
	//  1 ~ 2000
	SM_TEST = 1;
	//儒抚力绢 疙飞
	SM_STOPACTIONS = 2; //葛电 某腐磐/付过狼 悼累阑 肛冕促.
	//促弗 甘俊 甸绢埃 版快,

	//青悼俊 包访 疙飞
	SM_ACTION_MIN = 5;
	SM_THROW      = 5;
	SM_RUSH       = 6; //菊栏肺 傈柳
	SM_RUSHKUNG   = 7; //菊栏肺 傈柳角菩
	SM_FIREHIT    = 8; //堪拳搬
	SM_BACKSTEP   = 9; //缔吧澜龙,
	SM_TURN       = 10;
	SM_WALK       = 11;
	SM_SITDOWN    = 12;
	SM_RUN        = 13;
	SM_HIT        = 14;
	SM_HEAVYHIT   = 15;
	SM_BIGHIT     = 16;
	SM_SPELL      = 17;
	SM_POWERHIT   = 18;
	SM_LONGHIT    = 19; //歹 技霸 锭覆
	SM_DIGUP      = 20; //顶颇绊 唱坷促.
	SM_DIGDOWN    = 21; //顶颇绊 甸绢啊 见促.
	SM_FLYAXE     = 22;
	SM_LIGHTING   = 23; //付过 荤侩
	SM_WIDEHIT    = 24;
	SM_ACTION_MAX = 25;
	// 2003/03/15 脚痹公傍
	SM_CROSSHIT = 35; //堡浅曼, 林函8鸥老 傍拜

	SM_ACTION2_MIN = 1000;
	//SM_READYFIREHIT         = 1000;  //努扼捞攫飘俊辑父 静烙, 堪拳搬 霖厚

	SM_ACTION2_MAX = 1099;
	SM_DIE       = 26; //荤扼 咙
	SM_ALIVE     = 27;
	SM_MOVEFAIL  = 28;
	SM_HIDE      = 29;
	SM_DISAPPEAR = 30;
	SM_STRUCK    = 31;
	SM_DEATH     = 32;
	SM_SKELETON  = 33;
	SM_NOWDEATH  = 34;
	SM_HEAR           = 40;
	SM_FEATURECHANGED = 41;
	SM_USERNAME       = 42;
	SM_WINEXP         = 44;
	SM_LEVELUP        = 45;
	SM_DAYCHANGING    = 46;
	SM_LOGON              = 50;
	SM_NEWMAP             = 51;
	SM_ABILITY            = 52;
	SM_HEALTHSPELLCHANGED = 53;
	SM_MAPDESCRIPTION     = 54;
	SM_SYSMESSAGE   = 100;
	SM_GROUPMESSAGE = 101;
	SM_CRY          = 102;
	SM_WHISPER      = 103;
	SM_GUILDMESSAGE = 104;

	//ITEM ?
	SM_ADDITEM    = 200; //酒捞袍阑 货肺 掘澜  Series(荐樊)
	SM_BAGITEMS   = 201; //啊规狼 葛电 酒捞袍
	SM_DELITEM    = 202; //粹酒辑 绝绢瘤绰 殿狼 捞蜡肺 绝绢咙
	SM_UPDATEITEM = 203; //酒捞袍狼 荤剧捞 函窃
	//Magic
	SM_ADDMAGIC    = 210;
	SM_SENDMYMAGIC = 211; //
	SM_DELMAGIC    = 212;
	SM_VERSION_AVAILABLE   = 500;
	SM_VERSION_FAIL        = 501;
	SM_PASSWD_SUCCESS      = 502;
	SM_PASSWD_FAIL         = 503;
	SM_NEWID_SUCCESS       = 504; //货酒捞叼 肋 父甸绢 脸澜
	SM_NEWID_FAIL          = 505; //货酒捞叼 父甸扁 角菩
	SM_CHGPASSWD_SUCCESS   = 506;
	SM_CHGPASSWD_FAIL      = 507;
	SM_QUERYCHR            = 520; //某腐府胶飘
	SM_NEWCHR_SUCCESS      = 521;
	SM_NEWCHR_FAIL         = 522;
	SM_DELCHR_SUCCESS      = 523;
	SM_DELCHR_FAIL         = 524;
	SM_STARTPLAY           = 525;
	SM_STARTFAIL           = 526;
	SM_QUERYCHR_FAIL       = 527;
	SM_OUTOFCONNECTION     = 528; //楷搬 秦力凳
	SM_PASSOK_SELECTSERVER = 529;
	SM_SELECTSERVER_OK     = 530;
	SM_NEEDUPDATE_ACCOUNT  = 531; //拌沥狼 沥焊甫 促矫 涝仿窍扁 官恩 芒..
	SM_UPDATEID_SUCCESS    = 532;
	SM_UPDATEID_FAIL       = 533;
	SM_PASSOK_WRONGSSN     = 534;
	SM_NOT_IN_SERVICE      = 535;
	SM_DROPITEM_SUCCESS    = 600; //酒捞袍 滚府扁 己傍
	SM_DROPITEM_FAIL       = 601; //
	SM_ITEMSHOW            = 610;
	SM_ITEMHIDE            = 611;
	SM_OPENDOOR_OK         = 612;
	SM_OPENDOOR_LOCK       = 613;
	SM_CLOSEDOOR           = 614;
	SM_TAKEON_OK           = 615; //馒侩 己傍, + New Feature
	SM_TAKEON_FAIL         = 616; //馒侩 角菩
	SM_EXCHGTAKEON_OK      = 617; //馒侩酒捞袍 背券 己傍
	SM_EXCHGTAKEON_FAIL    = 618; //馒侩酒捞袍 背券 角菩
	SM_TAKEOFF_OK          = 619; //哈扁 己傍, + New Feature
	SM_TAKEOFF_FAIL        = 620; //
	SM_SENDUSEITEMS        = 621; //馒侩 酒捞袍 葛滴 焊晨
	SM_WEIGHTCHANGED       = 622;
	SM_CLEAROBJECTS        = 633;
	SM_CHANGEMAP           = 634;
	SM_EAT_OK              = 635;
	SM_EAT_FAIL            = 636;
	SM_BUTCH               = 637;
	SM_MAGICFIRE           = 638; //付过 惯荤凳  CM_SPELL -> SM_SPELL + SM_MAGICFIRE
	SM_MAGICFIRE_FAIL      = 639;
	SM_MAGIC_LVEXP         = 640;
	SM_SOUND               = 641;
	SM_DURACHANGE          = 642;
	SM_MERCHANTSAY         = 643;
	SM_MERCHANTDLGCLOSE    = 644;
	SM_SENDGOODSLIST       = 645;
	SM_SENDUSERSELL        = 646;
	SM_SENDBUYPRICE        = 647;
	SM_USERSELLITEM_OK     = 648;
	SM_USERSELLITEM_FAIL   = 649;
	SM_BUYITEM_SUCCESS     = 650;
	SM_BUYITEM_FAIL        = 651;
	SM_SENDDETAILGOODSLIST = 652;
	SM_GOLDCHANGED         = 653;
	SM_CHANGELIGHT         = 654;
	SM_LAMPCHANGEDURA      = 655;
	SM_CHANGENAMECOLOR     = 656;
	SM_CHARSTATUSCHANGED   = 657;
	SM_SENDNOTICE          = 658;
	SM_GROUPMODECHANGED    = 659;
	SM_CREATEGROUP_OK      = 660;
	SM_CREATEGROUP_FAIL    = 661;
	SM_GROUPADDMEM_OK      = 662;
	SM_GROUPDELMEM_OK      = 663;
	SM_GROUPADDMEM_FAIL    = 664;
	SM_GROUPDELMEM_FAIL    = 665;
	SM_GROUPCANCEL         = 666;
	SM_GROUPMEMBERS        = 667;
	SM_SENDUSERREPAIR      = 668;
	SM_USERREPAIRITEM_OK   = 669;
	SM_USERREPAIRITEM_FAIL = 670;
	SM_SENDREPAIRCOST      = 671;
	SM_DEALMENU            = 673;
	SM_DEALTRY_FAIL        = 674;
	SM_DEALADDITEM_OK      = 675;
	SM_DEALADDITEM_FAIL    = 676;
	SM_DEALDELITEM_OK      = 677;
	SM_DEALDELITEM_FAIL    = 678;
	//SM_DEALREMOTEADDITEM_OK = 679;
	//SM_DEALREMOTEDELITEM_OK = 680;
	SM_DEALCANCEL                  = 681; //档吝俊 芭贰 秒家凳
	SM_DEALREMOTEADDITEM           = 682; //惑措规捞 背券 酒捞袍阑 眠啊
	SM_DEALREMOTEDELITEM           = 683; //惑措规捞 背券 酒捞袍阑 画
	SM_DEALCHGGOLD_OK              = 684;
	SM_DEALCHGGOLD_FAIL            = 685;
	SM_DEALREMOTECHGGOLD           = 686;
	SM_DEALSUCCESS                 = 687;
	SM_SENDUSERSTORAGEITEM         = 700;
	SM_STORAGE_OK                  = 701;
	SM_STORAGE_FULL                = 702; //歹 焊包 给 窃.
	SM_STORAGE_FAIL                = 703; //焊包 俊矾
	SM_SAVEITEMLIST                = 704;
	SM_TAKEBACKSTORAGEITEM_OK      = 705;
	SM_TAKEBACKSTORAGEITEM_FAIL    = 706;
	SM_TAKEBACKSTORAGEITEM_FULLBAG = 707;
	SM_AREASTATE                   = 708; //救傈,措访,老馆..
	SM_DELITEMS                    = 709;
	SM_READMINIMAP_OK              = 710;
	SM_READMINIMAP_FAIL            = 711;
	SM_SENDUSERMAKEDRUGITEMLIST    = 712;
	SM_MAKEDRUG_SUCCESS            = 713;
	SM_MAKEDRUG_FAIL               = 714;
	SM_ALLOWPOWERHIT               = 715;
	SM_NORMALEFFECT                = 716; //扁夯 瓤苞

	SM_CHANGEGUILDNAME      = 750; //辨靛狼 捞抚 趣狼 辨靛郴狼 流氓捞抚捞 函版
	SM_SENDUSERSTATE        = 751; //
	SM_SUBABILITY           = 752;
	SM_OPENGUILDDLG         = 753;
	SM_OPENGUILDDLG_FAIL    = 754;
	SM_SENDGUILDHOME        = 755;
	SM_SENDGUILDMEMBERLIST  = 756;
	SM_GUILDADDMEMBER_OK    = 757;
	SM_GUILDADDMEMBER_FAIL  = 758;
	SM_GUILDDELMEMBER_OK    = 759;
	SM_GUILDDELMEMBER_FAIL  = 760;
	SM_GUILDRANKUPDATE_FAIL = 761;
	SM_BUILDGUILD_OK        = 762;
	SM_BUILDGUILD_FAIL      = 763;
	SM_DONATE_FAIL          = 764;
	SM_DONATE_OK            = 765;
	SM_MYSTATUS             = 766;
	SM_MENU_OK              = 767; //description栏肺 皋技瘤 傈崔
	SM_GUILDMAKEALLY_OK     = 768;
	SM_GUILDMAKEALLY_FAIL   = 769;
	SM_GUILDBREAKALLY_OK    = 770;
	SM_GUILDBREAKALLY_FAIL  = 771;
	SM_DLGMSG               = 772;
	SM_SPACEMOVE_HIDE  = 800; //鉴埃捞悼 荤扼咙
	SM_SPACEMOVE_SHOW  = 801; //唱鸥巢
	SM_RECONNECT       = 802;
	SM_GHOST           = 803; //拳搁俊 唱鸥抄 儡惑烙
	SM_SHOWEVENT       = 804;
	SM_HIDEEVENT       = 805;
	SM_SPACEMOVE_HIDE2 = 806; //鉴埃捞悼 荤扼咙
	SM_SPACEMOVE_SHOW2 = 807; //唱鸥巢

	SM_TIMECHECK_MSG = 810; //努扼捞攫飘俊辑 矫埃
	SM_ADJUST_BONUS  = 811; //焊呈胶 器牢飘甫 炼沥窍扼.
	// Frined System -------------
	SM_FRIEND_DELETE = 812; //模备 昏力
	SM_FRIEND_INFO   = 813; //模备 眠啊 棺 沥焊函版
	SM_FRIEND_RESULT = 814; //模备包访 搬苞蔼 傈价
	// Tag System ----------------
	SM_TAG_ALARM         = 815; //率瘤吭澜 舅覆
	SM_TAG_LIST          = 816; //率瘤府胶飘
	SM_TAG_INFO          = 817; //率瘤沥焊 函版
	SM_TAG_REJECT_LIST   = 818; //芭何磊 府胶飘
	SM_TAG_REJECT_ADD    = 819; //芭何磊 眠啊
	SM_TAG_REJECT_DELETE = 820; //芭何磊 昏力
	SM_TAG_RESULT        = 821; //率瘤包访 搬苞蔼 傈价
	// User System ---------------
	SM_USER_INFO = 822; //蜡历狼 立加惑怕棺 甘沥焊傈价

	//1000 ~ 1099  咀记栏肺 抗距

	SM_OPENHEALTH        = 1100; //眉仿捞 惑措规俊 焊烙
	SM_CLOSEHEALTH       = 1101; //眉仿捞 惑措规俊霸 焊捞瘤 臼澜
	SM_BREAKWEAPON       = 1102;
	SM_INSTANCEHEALGUAGE = 1103;
	SM_CHANGEFACE        = 1104; //函脚...
	SM_NEXTTIME_PASSWORD = 1105; //促澜锅俊绰 厚剐锅龋 涝仿 葛靛捞促.
	SM_CHECK_CLIENTVALID = 1106; //努扼捞攫飘狼 荐沥 咯何 犬牢

	SM_PLAYDICE = 1200;
	// 2003/02/11 弊缝盔 困摹 沥焊
	SM_GROUPPOS = 1312;

	//CM_??   努扼捞攫飘 -> 辑滚肺
	//  2000 ~ 4000
	CM_PROTOCOL       = 2000;
	CM_IDPASSWORD     = 2001;
	CM_ADDNEWUSER     = 2002;
	CM_CHANGEPASSWORD = 2003;
	CM_UPDATEUSER     = 2004;

	CM_QUERYCHR = 100;
	CM_NEWCHR = 101;
	CM_DELCHR               = 102;
	CM_SELCHR = 103;
	CM_SELECTSERVER = 104; //辑滚甫 急琶 (+ 辑滚捞抚)

	//3000 - 3099 努扼捞攫飘 捞悼 皋技瘤档 抗距
	//辑滚俊辑 捞悼 皋技瘤档 0..99 荤捞 捞绢具 茄促.
	CM_THROW = 3005;
	CM_TURN = 3010; //CM_TURN - 3000 = SM_TURN 痹蘑阑 馆靛矫 瘤难具 窃
	CM_WALK                 = 3011;
	CM_SITDOWN = 3012;
	CM_RUN = 3013;
	CM_HIT = 3014;
	CM_HEAVYHIT = 3015;
	CM_BIGHIT               = 3016;
	CM_SPELL = 3017;
	CM_POWERHIT = 3018; //歹 技霸 锭覆
	CM_LONGHIT = 3019;  //歹 技霸 锭覆
	CM_WIDEHIT = 3024;
	CM_FIREHIT              = 3025;
	CM_SAY = 3030;
	// 2003/03/15 脚痹公傍
	CM_CROSSHIT = 3035;

	CM_QUERYUSERNAME = 80; //QUERY 矫府令 疙飞绢
	CM_QUERYBAGITEMS = 81;
	CM_QUERYUSERSTATE       = 82; //鸥牢狼 惑怕 焊扁

	CM_DROPITEM = 1000;
	CM_PICKUP = 1001;
	CM_OPENDOOR = 1002;
	CM_TAKEONITEM = 1003; //汗厘阑 馒侩
	CM_TAKEOFFITEM          = 1004; //汗厘阑 哈绰促
	CM_EXCHGTAKEONITEM = 1005;      //馒侩茄 酒捞袍阑 谅快甫 官槽促.(馆瘤,迫骂)
	CM_EAT = 1006;   //冈促, 付矫促
	CM_BUTCH = 1007; //档氟窍促
	CM_MAGICKEYCHANGE = 1008;
	CM_SOFTCLOSE            = 1009;
	CM_CLICKNPC = 1010;
	CM_MERCHANTDLGSELECT = 1011;
	CM_MERCHANTQUERYSELLPRICE = 1012;
	CM_USERSELLITEM = 1013; //酒捞袍 迫扁
	CM_USERBUYITEM          = 1014;
	CM_USERGETDETAILITEM = 1015;
	CM_DROPGOLD = 1016;
	CM_TEST = 1017; //抛胶飘
	CM_LOGINNOTICEOK = 1018;
	CM_GROUPMODE            = 1019;
	CM_CREATEGROUP = 1020;
	CM_ADDGROUPMEMBER = 1021;
	CM_DELGROUPMEMBER = 1022;
	CM_USERREPAIRITEM = 1023;
	CM_MERCHANTQUERYREPAIRCOST = 1024;
	CM_DEALTRY = 1025;
	CM_DEALADDITEM = 1026;
	CM_DEALDELITEM = 1027;
	CM_DEALCANCEL = 1028;
	CM_DEALCHGGOLD          = 1029; //背券窍绰 捣捞 函版凳
	CM_DEALEND = 1030;
	CM_USERSTORAGEITEM = 1031;
	CM_USERTAKEBACKSTORAGEITEM = 1032;
	CM_WANTMINIMAP = 1033;
	CM_USERMAKEDRUGITEM     = 1034;
	CM_OPENGUILDDLG = 1035;
	CM_GUILDHOME = 1036;
	CM_GUILDMEMBERLIST = 1037;
	CM_GUILDADDMEMBER = 1038;
	CM_GUILDDELMEMBER       = 1039;
	CM_GUILDUPDATENOTICE = 1040;
	CM_GUILDUPDATERANKINFO = 1041;
	CM_SPEEDHACKUSER = 1042;
	CM_ADJUST_BONUS = 1043;
	CM_GUILDMAKEALLY        = 1044;
	CM_GUILDBREAKALLY = 1045;
	// Frined System---------------
	CM_FRIEND_ADD = 1046;    // 模备眠啊
	CM_FRIEND_DELETE = 1047; // 模备昏力
	CM_FRIEND_EDIT = 1048;   // 模备汲疙 函版
	CM_FRIEND_LIST          = 1049; // 模备 府胶飘 夸没
	// Tag System -----------------
	CM_TAG_ADD = 1050; // 率瘤 眠啊
	CM_TAG_DELETE = 1051;  // 率瘤 昏力
	CM_TAG_SETINFO = 1052; // 率瘤 惑怕 函版
	CM_TAG_LIST = 1053;    // 率瘤 府胶飘 夸没
	CM_TAG_NOTREADCOUNT     = 1054; // 佬瘤臼篮 率瘤 俺荐 夸没
	CM_TAG_REJECT_LIST = 1055;      // 芭何磊 府胶飘
	CM_TAG_REJECT_ADD = 1056;    // 芭何磊 眠啊
	CM_TAG_REJECT_DELETE = 1057; // 芭何磊 昏力

	CM_CLIENT_CHECKTIME = 1100;
	CM_GAMELOGIN=8888;
)
