package model

type BookInfo struct {
	CountDay     string   `json:"count_day"`
	BookId       string   `json:"book_id"`
	BookName     string   `json:"book_name"`
	Content      string   `json:"content"`
	Reason       string   `json:"reason"`
	SortTypelist string   `json:"sort_typelist"`
	BookTypes    []string `json:"book_types"`
	ImgUrl       []string `json:"img_urls"`
	Url          string   `json:"url"`
}

type CartoonTag struct {
	Cid      int    `json:"cid" gorm:"column:cid"`
	Tag      string `json:"tag" gorm:"column:tag"`
	Category string `json:"category" gorm:"column:category"`
}

type TagInfo struct {
	Tag      string
	Category string
}

func (c *CartoonTag) TableName() string {
	return "rule_nor_tag_for_cartoon"
}

type CartoonDescription struct {
	CartoonId string `json:"cartoon_id" gorm:"column:cartoon_id"`
	Title     string `json:"title" gorm:"column:title"`
}

func (c *CartoonDescription) TableName() string {
	return "cartoon_new_title"
}

type ComicFree struct {
	DateFlag string `json:"date_flag" gorm:"column:date_flag"`
	ComicIds string `json:"ids" gorm:"column:ids"`
}

func (c *ComicFree) TableName() string {
	return "cartoon_limit_pay"
}

type CartoonClassify struct {
	MainType string `json:"main_type" gorm:"column:main_type"`
}

func (c *CartoonClassify) TableName() string {
	return "cartoon_type"
}
