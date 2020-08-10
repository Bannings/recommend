package model

var (
	Banner = new(BannerData)
)

type BannerInfo struct {
	ComicId   int    `json:"comic_id" gorm:"column:comic_id"`
	ComicName string `json:"comic_name" gorm:"column:comic_name"`
	Url       string `json:"url" gorm:"column:url"`
	ImgUrl    string `json:"img_url" gorm:"column:img_url"`
	Platform  int    `json:"platform" gorm:"column:platform"`
}

type BannerData struct {
	BannerDatas [][]BannerInfo
}

func (banner *BannerInfo) TableName() string {
	return "banner_operation"
}
