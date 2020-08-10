package model

//type BookList struct {
//	CountDay    string   `json:"count_day" `
//	BookID      string   `json:"book_id" `
//	BookName    string   `json:"book_name" `
//	ComicUrlID  string   `json:"comic_urlid" `
//	AuthorName  string   `json:"author_name" `
//	Content     string   `json:"content"`
//	TypeList    string   `json:"sort_typelist"`
//	BookFeature string   `json:"book_feature"`
//	ImgUrl      []string `json:"img_url"`
//	Url         string   `json:"url"`
//	Reason      string   `json:"reason"`
//}

type BillBoardComic struct {
	CountDay              string `json:"count_day" db:"count_day"`
	ComicId               string `json:"comic_id" db:"comic_id"`
	ComicName             string `json:"comic_name" db:"comic_name"`
	SortTypelist          string `json:"sort_typelist" db:"sort_typelist"`
	ComicScore            string `json:"comic_score" db:"comic_score"`
	CountNum              string `json:"count_num" db:"count_num"`
	AuthorName            string `json:"author_name" db:"author_name"`
	LastchapterId         string `json:"lastchapter_id" db:"lastchapter_id"`
	LastchapterUrlid      string `json:"lastchapter_urlid" db:"lastchapter_urlid"`
	ComicUrlid            string `json:"comic_urlid" db:"comic_urlid"`
	ComicFeature          string `json:"comic_feature" db:"comic_feature"`
	LastchapterTitle      string `json:"lastchapter_title" db:"lastchapter_title"`
	RiseRank              string `json:"rise_rank" db:"rise_rank"`
	RiseWithinLeaderboard string `json:"rise_within_leaderboard"`
}

type RankInfo struct {
	RankType string           `json:"rank_type"`
	RankName string           `json:"rank_name"`
	Info     []BillBoardComic `json:"rank_info"`
}

type BillBoardReq struct {
	Uid           int64 `json:"uid" form:"uid"`
	BillBoardNum  int   `json:"board_num" form:"board_num"`
	BillBoardSize int   `json:"board_size" form:"board_size"`
}
