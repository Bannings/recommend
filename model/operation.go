package model

type MgResourceVO struct {
	ResType int    `json:"resType"`
	URL     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type MgOperationVO struct {
	OpId         string `json:"opId"`
	OpName       string `json:"opName"`
	OpDesc       string `json:"opDesc"`
	OpAcitonType int    `json:"opAcitonType"`
	OpActionInfo string `json:"opActionInfo"`
	EsIds        string `json:"EsIds"`
}

type Operationinfo struct {
	OposId        string               `json:"oposId"`
	OposName      string               `json:"oposName"`
	OposOrder     int                  `json:"oposOrder"`
	ResId         string               `json:"resId"`
	ResourceVO    MgResourceVO         `json:"mgResourceVO"`
	OpId          string               `json:"opId"`
	OperationVO   MgOperationVO        `json:"mgOperationVO"`
	CloseFeatures int                  `json:"closeFeatures"`
	CloseShow     int                  `json:"closeShow"`
	Frequency     int                  `json:"frequency"`
	ABTest        int                  `json:"abTest"`
	Passthrough   *OperationinfoABExps `json:"passthrough,omitempty"`
}

type OperationinfoABExps struct {
	ExpId      string `json:"exp_id"`
	BucketId   string `json:"bucket_id"`
	BucketName string `json:"bucket_name"`
	Uid        int64  `json:"uid"`
	Udid       string `json:"udid"`
	ServeTime  int64  `json:"serve_time"`
}

type OperationData struct {
	Message string          `json:"message"`
	Data    []Operationinfo `json:"data"`
}
