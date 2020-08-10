package model

type RelateRecommendReq struct {
	Uid          int64    `json:"uid" form:"uid"`
	Cid          int      `json:"cid" form:"cid" binding:"required"`
	ComicName    string   `json:"comic_name" form:"comic_name"`
	SectionOrder int      `json:"section_order" form:"section_order"`
	Gender       int      `json:"gender" form:"gender"`
	ReadingCount int      `json:"reading_count" form:"reading_count"`
	SessionId    string   `json:"sessionid" form:"sessionid"`
	Expose       []string `json:"expose" form:"expose"`
}
