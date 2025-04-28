package response

type Seal struct {
	SealId                 string `json:"sealId"`
	SealName               string `json:"sealName"`
	SealCreateTime         int64  `json:"sealCreateTime"`
	DefaultSealFlag        bool   `json:"defaultSealFlag"`
	SealWidth              int    `json:"sealWidth"`
	SealHeight             int    `json:"sealHeight"`
	SealBizType            string `json:"sealBizType"`
	SealBizTypeDescription string `json:"sealBizTypeDescription"`
	SealStyle              int    `json:"sealStyle"`
	SealStatus             int    `json:"sealStatus"`
	StatusDescription      string `json:"statusDescription"`
	RejectReason           string `json:"rejectReason"`
	SealImageDownloadUrl   string `json:"sealImageDownloadUrl"`
}

type SealData struct {
	Total int    `json:"total"`
	Seals []Seal `json:"seals"`
}

type SealRes struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Data    SealData `json:"data"`
}
