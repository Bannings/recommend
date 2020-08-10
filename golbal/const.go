package golbal

//
const (
	INDEX_BANNER               = "0" //老版banner
	INDEX_HIGHLY               = "1" //老版人气
	INDEX_QUALITY              = "2" //老版
	INDEX_UNIQUE               = "3" //老版独家
	INDEX_NEW                  = "4" //老版新作
	INDEX_VIP                  = "5" //vip
	INDEX_PERSONAL             = "6" //猜你喜欢
	OLD_CLASSIFY_LOVE          = "7"
	OLD_CLASSIFY_FANTASY       = "8"
	OLD_CLASSIFY_NEW           = "9"
	OLD_CLASSIFY_TIMETRAVEL    = "10"
	OLD_CLASSIFY_SCHOOL        = "11"
	OLD_CLASSIFY_VIP           = "12"
	OLD_CLASSIFY_AMUSE         = "13"
	OLD_CLASSIFY_CITY          = "14"
	OLD_CLASSIFY_SCIENCE       = "15"
	OLD_CLASSIFY_PRESIDENT     = "16"
	OLD_CLASSIFY_BLOOD         = "17"
	INDEX_TYPE                 = "18" //老版编辑部
	CLASSIFY_LOVE              = "19"
	CLASSIFY_FANTASY           = "20"
	CLASSIFY_NEW               = "21"
	CLASSIFY_TIMETRAVEL        = "22"
	CLASSIFY_SCHOOL            = "23"
	CLASSIFY_VIP               = "24"
	CLASSIFY_AMUSE             = "25"
	CLASSIFY_CITY              = "26"
	CLASSIFY_SCIENCE           = "27"
	CLASSIFY_PRESIDENT         = "28"
	CLASSIFY_BLOOD             = "29"
	INDEX_FREE                 = "30" //免费漫画
	INDEX_EXPOSURE             = "31"
	INDEX_BOOKLIST             = "32" //老版动态书单
	INDEX_BILLBOARD            = "33" //mini榜单
	INDEX_RELATE               = "34" //相关推荐
	INDEX_BEHAVIOR             = "35" //老版编辑部
	INDEX_COMICLIST1           = "36"
	INDEX_COMICLIST2           = "37"
	INDEX_SPECIAL              = "38"
	INDEX_BEHAVIORRELATE       = "39"
	INDEX_BOOKLIST1            = "40"
	INDEX_BOOKLIST2            = "41"
	INDEX_BOOKLIST3            = "42"
	INDEX_BOOKLIST4            = "43"
	INDEX_OPERATION_BANNER     = "44" //运营banner
	INDEX_OPERATION_AD         = "45" //运营banner广告
	INDEX_HIGHLY_NEW           = "46" //新版人气
	INDEX_FREE_NEW             = "47" //新版免费
	INDEX_NOVEL                = "48" //小说漫改
	INDEX_AD                   = "49" //客户端广告
	INDEX_SUPERNATANT          = "70"
	INDEX_DETAIL_COLLABORATION = "71"
	INDEX_DETAIL_ITEM          = "72"
	OLD_INDEX_TYPE             = "99"
	INDEX_DYNAMIC_COMIC_LIST   = "100"
	DYNAMIC_COMIC_LIST         = 100
	RECOMMEND_COMIC_LIST       = 500
	NEW_USER_COMIC_LIST        = 1000000
)

const (
	MALE_ID                     = 0
	FEMALE_ID                   = 1
	MIX_GENDER_ID               = 2
	MALE_STR                    = "male"
	FEMALE_STR                  = "female"
	BANNER_PLATFORM_ALL         = 0
	BANNER_PLATFORM_ANDROID     = 1
	BANNER_PLATFORM_IOS         = 2
	BANNER_PLATFORM_WXAPP       = 3
	BANNER_PLATFORM_ANDROID_STR = "android"
	BANNER_PLATFORM_IOS_STR     = "ios"
	VERSION                     = "0.2.5"

	COMIC_SHOW = "show"

	ComicListNum    = 4 //书单数
	MaxRecommendLen = 200
)

var (
	ColumnNameMap = map[string]string{
		INDEX_BANNER:           "banner",
		INDEX_HIGHLY:           "入坑只需1秒",
		INDEX_QUALITY:          "不好看你打我吧",
		INDEX_UNIQUE:           "独家·只此一家别无分店",
		INDEX_NEW:              "%s月番抢先看",
		INDEX_FREE:             "免费漫画看到爽",
		INDEX_VIP:              "VIP专属",
		INDEX_PERSONAL:         "为你推荐",
		INDEX_BOOKLIST:         "推荐书单",
		INDEX_BILLBOARD:        "mini榜单",
		INDEX_SPECIAL:          "五一爆更漫画",
		INDEX_BEHAVIOR:         "来自编辑部の推荐",
		INDEX_BEHAVIORRELATE:   "来自编辑部の推荐",
		INDEX_TYPE:             "来自编辑部の推荐",
		OLD_INDEX_TYPE:         "来自编辑部の推荐",
		INDEX_OPERATION_BANNER: "运营banner",
		INDEX_OPERATION_AD:     "运营广告",
		INDEX_HIGHLY_NEW:       "入坑只需1秒",
		INDEX_FREE_NEW:         "免费漫画看到爽",
		INDEX_NOVEL:            "小说漫改",
	}

	WxappColumnMap = map[string]string{
		INDEX_HIGHLY:   "热门作品大放送",
		INDEX_QUALITY:  "飒漫精品",
		INDEX_UNIQUE:   "独家作品",
		INDEX_NEW:      "上新了，殿下",
		INDEX_VIP:      "VIP热销",
		INDEX_PERSONAL: "猜你喜欢看",
		INDEX_TYPE:     "来自编辑部の推荐",
	}
	OldColumns = []string{
		INDEX_BANNER,
		INDEX_PERSONAL,
		INDEX_HIGHLY,
		INDEX_NEW,
		INDEX_FREE,
		INDEX_VIP,
	}
	NewColumns = []string{
		INDEX_BANNER,
		INDEX_HIGHLY,   //超人气！
		INDEX_PERSONAL, //猜你喜欢看
		INDEX_NEW,      //上新了
		INDEX_FREE,     //免费漫画看到爽
		INDEX_VIP,      //VIP热销

	}
	NewColumnOrderMap = map[string]int{
		INDEX_BANNER:    0,
		INDEX_HIGHLY:    1, //超人气！
		INDEX_PERSONAL:  2, //猜你喜欢看
		INDEX_NEW:       3, //上新了
		INDEX_BOOKLIST:  4, //书单
		INDEX_BILLBOARD: 5, //排行榜
		INDEX_FREE:      6, //免费漫画看到爽
		INDEX_VIP:       7, //VIP热销
	}
	OldColumnOrderMap = map[string]int{
		INDEX_BANNER:    0,
		INDEX_PERSONAL:  1, //猜你喜欢看
		INDEX_HIGHLY:    2, //超人气！
		INDEX_NEW:       3, //上新了
		INDEX_BOOKLIST:  4, //书单
		INDEX_BILLBOARD: 5, //排行榜
		INDEX_FREE:      6, //免费漫画看到爽
		INDEX_VIP:       7, //VIP热销
	}
	V2Columns = []string{
		INDEX_BANNER,
		INDEX_PERSONAL,
		INDEX_HIGHLY,
		INDEX_NEW,
		INDEX_FREE,
		INDEX_VIP,
	}
	V3Columns = []string{
		INDEX_OPERATION_BANNER,
		INDEX_PERSONAL,
		INDEX_OPERATION_AD,
		INDEX_HIGHLY_NEW,
		INDEX_OPERATION_AD,
		INDEX_NEW,
		INDEX_OPERATION_AD,
		INDEX_NOVEL,
		INDEX_OPERATION_AD,
		INDEX_FREE_NEW,
		INDEX_OPERATION_AD,
		INDEX_VIP,
		INDEX_OPERATION_AD,
	}

	RecommendRecallColumns = []string{
		INDEX_PERSONAL,
		INDEX_HIGHLY_NEW,
		INDEX_NEW,
		INDEX_NOVEL,
		INDEX_FREE_NEW,
		INDEX_VIP,
	}
	V2ColumnOrderMap = map[string]int{
		INDEX_BANNER:   0,
		INDEX_PERSONAL: 1, //猜你喜欢看
		INDEX_HIGHLY:   2,
		INDEX_NEW:      3, //X月新番
		INDEX_FREE:     4,
		INDEX_VIP:      5,
	}
	V3ColumnOrderMap = map[string]int{
		INDEX_OPERATION_BANNER: 0,
		INDEX_PERSONAL:         1, //猜你喜欢看
		INDEX_HIGHLY_NEW:       2,
		INDEX_OPERATION_AD:     3,
		INDEX_NEW:              4, //X月新番
		INDEX_NOVEL:            5,
		INDEX_FREE_NEW:         6,
		INDEX_VIP:              7,
	}
	WxappTolColumns = []string{
		INDEX_BANNER,
		INDEX_HIGHLY,
		INDEX_NEW,
		INDEX_PERSONAL,
	}

	NewUserInterestCtypes = map[string]string{
		"male":   "玄幻|异能|悬疑|搞笑|校园|都市|热血|后宫|穿越|古风",
		"female": "校园|古风|霸总|玄幻|都市|悬疑|穿越|搞笑|日常|真人",
	}
)
