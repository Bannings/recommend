package model

import (
	. "recommend/golbal"
)

var (
	BannerSection         = &SectionDataConfig{DisplayType: 1, SectionType: "banner", Config: SectionConfig{ShowHeader: 0, ShowSubtitle: 0, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	HighlySection         = &SectionDataConfig{DisplayType: 2, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 1, ShowSwitch: 1, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	QualitySection        = &SectionDataConfig{DisplayType: 5, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	UniqueSection         = &SectionDataConfig{DisplayType: 2, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 1, ShowSwitch: 1, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	NewlySection          = &SectionDataConfig{DisplayType: 7, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	VipSection            = &SectionDataConfig{DisplayType: 7, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	FreeSection           = &SectionDataConfig{DisplayType: 7, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	PersonalSection       = &SectionDataConfig{DisplayType: 7, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	OldTypeSection        = &SectionDataConfig{DisplayType: 7, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 0, ShowSwitch: 0, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	TypeSection           = &SectionDataConfig{DisplayType: 8, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 0, ShowSubtitle: 1, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	RelateSection         = &SectionDataConfig{DisplayType: 13, SectionType: "实时推荐", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	BehaviorSection       = &SectionDataConfig{DisplayType: 14, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	BehaviorRelateSection = &SectionDataConfig{DisplayType: 14, SectionType: "实时推荐", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	ComicListSection2     = &SectionDataConfig{DisplayType: 3, SectionType: "漫画", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	ComicListSection1     = &SectionDataConfig{DisplayType: 2, SectionType: "漫画", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 1, ShowSwitch: 1, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	FixedType1            = &SectionDataConfig{DisplayType: 18, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 1, ShowMore: 1, ShowSwitch: 1, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	FixedType2            = &SectionDataConfig{DisplayType: 19, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	FixedType3            = &SectionDataConfig{DisplayType: 20, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 1, ShowSubtitle: 0, ShowMore: 1, ShowSwitch: 1, ShowLabel: 1, ShowCountNum: 0, ShowSectionSubtitle: 1}}
	ADSection             = &SectionDataConfig{DisplayType: 21, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 0, ShowSubtitle: 0, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	OperatonBannerSection = &SectionDataConfig{DisplayType: 23, SectionType: "banner", Config: SectionConfig{ShowHeader: 0, ShowSubtitle: 0, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	OperatonADSection     = &SectionDataConfig{DisplayType: 24, SectionType: "固定栏", Config: SectionConfig{ShowHeader: 0, ShowSubtitle: 0, ShowMore: 0, ShowSwitch: 0, ShowLabel: 0, ShowCountNum: 0, ShowSectionSubtitle: 0}}
	DisplayLenMap         = map[int]int{1: 4, 2: 4, 3: 4, 4: 3, 5: 6, 6: 4, 7: 6, 11: 4, 12: 4, 18: 4, 19: 4, 20: 8} //每个样式对应显示漫画数
)

type ComicInfo struct {
	CountDay         string `json:"count_day" gorm:"column:update_time"`
	ComicID          string `json:"comic_id" gorm:"column:cartoon_id"`
	ComicName        string `json:"comic_name" gorm:"column:cartoon_name"`
	ComicUrlID       string `json:"comic_urlid" gorm:"column:PinYin_name"`
	AuthorName       string `json:"author_name" gorm:"column:cartoon_author_list_name"`
	TypeList         string `json:"sort_typelist" gorm:"column:cartoon_type_list_name"`
	LastChapterID    string `json:"lastchapter_id" gorm:"column:latest_cartoon_topic_id"`
	LastChapterUrlID string `json:"lastchapter_urlid" gorm:"column:latest_cartoon_topic_newid"`
	LastChapterTitle string `json:"lastchapter_title" gorm:"column:latest_cartoon_topic_name"`
	ComicFeature     string `json:"comic_feature" gorm:"column:comic_feature"`
	ComicScore       string `json:"comic_score" gorm:"column:pingfen"`
	CountNum         string `json:"count_num" gorm:"column:total_view_num"`
	CollectNum       string `json:"shoucang" gorm:"column:shoucang"`
	Reason           string `json:"reason,omitempty"`
	CurrchapterId    string `json:"currchapter_id,omitempty"`
	CurrchapterTitle string `json:"currchapter_title,omitempty"`
	ShowName         string `json:"show_name,omitempty"`
	ImgUrl           string `json:"img_url"`
	Url              string `json:"url"`
	OptionCover      string `json:"comic_cover"`
	ComicCategory    string `json:"comic_category"`
	ComicRank        string `json:"comic_rank"`
	ComicTips        string `json:"comic_tips"`
}

func (c *ComicInfo) TableName() string {
	return "comic_main"
}

type SectionData struct {
	SectionName string      `json:"section_name"`
	SectionId   string      `json:"section_id"`
	ComicInfo   []ComicInfo `json:"comic_info"`
}

type SectionConfig struct {
	ShowHeader          int `json:"show_header"`           // 是否显示标题
	ShowSubtitle        int `json:"show_subtitle"`         // 是否显示漫画副标题
	ShowMore            int `json:"show_more"`             // 是否显示更多
	ShowSwitch          int `json:"show_switch"`           // 是否显示换一换
	ShowLabel           int `json:"show_label"`            // 是否显示标签（漫画类型）
	ShowCountNum        int `json:"show_count_num"`        // 是否显示人气值
	ShowSectionSubtitle int `json:"show_section_subtitle"` // 是否显示栏目副标题
}

type SectionDataConfig struct {
	SectionName     string          `json:"section_name"`     // 栏目标题
	SectionSubtitle string          `json:"section_subtitle"` // 栏目副标题
	SectionId       string          `json:"section_id"`       // 栏目ID（固定不变）
	SectionType     string          `json:"section_type"`     // 栏目类型
	SectionIcon     string          `json:"section_icon"`     // 栏目标题icon
	Passthrough     Passthrough     `json:"passthrough"`
	SectionOrder    int             `json:"section_order"` // 栏目所在位置
	DisplayType     int             `json:"display_type"`  // 栏目展示类型
	ComicInfo       []ComicInfo     `json:"comic_info,omitempty"`
	BookInfo        []BookInfo      `json:"book_info,omitempty"`
	OperationInfo   []Operationinfo `json:"operation_info,omitempty"`
	Config          SectionConfig   `json:"config"`
}

func GetSection(index string) *SectionDataConfig {
	switch index {
	case INDEX_BANNER:
		return BannerSection
	case INDEX_HIGHLY:
		return HighlySection
	case INDEX_QUALITY:
		return QualitySection
	case INDEX_UNIQUE:
		return UniqueSection
	case INDEX_NEW:
		return NewlySection
	case INDEX_VIP:
		return FixedType2
	case INDEX_FREE:
		return FreeSection
	case INDEX_PERSONAL:
		return PersonalSection
	case INDEX_TYPE:
		return TypeSection
	case OLD_INDEX_TYPE:
		return OldTypeSection
	case INDEX_RELATE:
		return RelateSection
	case INDEX_BEHAVIORRELATE:
		return BehaviorRelateSection
	case INDEX_BEHAVIOR:
		return BehaviorSection
	case INDEX_COMICLIST1:
		return ComicListSection1
	case INDEX_COMICLIST2:
		return ComicListSection2
	case INDEX_BOOKLIST1:
		return FixedType3
	case INDEX_BOOKLIST2:
		return FixedType1
	case INDEX_BOOKLIST3:
		return FixedType2
	case INDEX_OPERATION_AD:
		return OperatonADSection
	case INDEX_OPERATION_BANNER:
		return OperatonBannerSection
	case INDEX_HIGHLY_NEW:
		return FixedType1
	case INDEX_FREE_NEW:
		return FreeSection
	case INDEX_NOVEL:
		return FixedType2
	default:
		return PersonalSection
	}
}
