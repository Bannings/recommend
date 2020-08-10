package model

type ABType int

const (
	BannerA ABType = iota
	BannerB
	BannerC
)

type ABResponse struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Payload         string `json:"payload"`
	Assign          string `json:"assignment"`
	ExperimentLabel string `json:"experimentLabel"`
	Status          string `json:"status"`
}

type Passthrough struct {
	OnlineService interface{} `json:"online_service,omitempty"`
	ABTestExps    []ABExps    `json:"abtest_exps,omitempty"`
}

type ABExps struct {
	ExpId        string `json:"exp_id"`
	BucketId     string `json:"bucket_id"`
	BucketName   string `json:"bucket_name"`
	Uid          string `json:"uid"`
	Udid         string `json:"udid"`
	AlgId        string `json:"alg_id"`
	RecServiceId string `json:"rec_service_id"`
	ServeTime    int64  `json:"serve_time"`
}

type OnlineService struct {
	RecServiceId string `json:"rec_service_id"`
}

type RecommendResult struct {
	RecommendMap map[string][]int32
	ABTestExps   []ABExps
}
