package model

type ComicListConfig struct {
	ComicListNum      int `yaml:"comic_list_num"`       // 书单数
	MinRecommendIds   int `yaml:"min_recommend_ids"`    // 最小推荐长度
	MinComicListCount int `yaml:"min_comic_list_count"` // 书单最小漫画数量
}

// OutComicList 输出的书单，title和section_id已确定
type OutComicList struct {
	BookTitle    string   `yaml:"book_title"`
	CartoonIDs   []string `yaml:"cartoon_ids"`
	CartoonTypes []string `yaml:"cartoon_types"`
	SectionID    string   `yaml:"section_id"`
	Subtitle     string
}

// DynamicComicList 动态书单备选集（同种类型书单包含多个备选书单，title和section_id不同
type DynamicComicList struct {
	CartoonIDs   []string `yaml:"cartoon_ids"`
	CartoonTypes []string `yaml:"cartoon_types"`
	Titles       []string `yaml:"titles"`
	SectionIDs   []string `yaml:"section_ids"`
	Subtitle     string   `yaml:"subtitle"`
}

type DynamicComicLists []DynamicComicList
type GenderDynamicComicLists struct {
	Male   DynamicComicLists `yaml:"male"`
	Female DynamicComicLists `yaml:"female"`
}

func (d GenderDynamicComicLists) GetComicListOutMap() map[string]OutComicList {
	dynamicComicListSectionIDTitleMap := make(map[string]OutComicList, len(d.Male)+len(d.Female))
	for _, dynamicComicList := range d.Male {
		for i := range dynamicComicList.SectionIDs {
			dynamicComicListSectionIDTitleMap[dynamicComicList.SectionIDs[i]] = OutComicList{BookTitle: dynamicComicList.Titles[i],
				CartoonTypes: dynamicComicList.CartoonTypes, Subtitle: dynamicComicList.Subtitle}
		}
	}
	for _, dynamicComicList := range d.Female {
		for i := range dynamicComicList.SectionIDs {
			dynamicComicListSectionIDTitleMap[dynamicComicList.SectionIDs[i]] = OutComicList{BookTitle: dynamicComicList.Titles[i],
				CartoonTypes: dynamicComicList.CartoonTypes, Subtitle: dynamicComicList.Subtitle}
		}
	}
	return dynamicComicListSectionIDTitleMap
}
