package request

type Operator struct {
	PsnAccount string `json:"psnAccount"`
	PsnId      string `json:"psnId"`
}

type Organization struct {
	OrgId   string `json:"orgId"`
	OrgName string `json:"orgName"`
}

type SignUrlRedirectConfig struct {
	RedirectUrl       string `json:"redirectUrl"`
	RedirectDelayTime int    `json:"redirectDelayTime"`
}

type SignUrl struct {
	NeedLogin      bool                   `json:"needLogin"`
	UrlType        int32                  `json:"urlType"`
	Operator       Operator               `json:"operator"`
	Organization   *Organization          `json:"organization"`
	RedirectConfig *SignUrlRedirectConfig `json:"redirectConfig"`
	ClientType     string                 `json:"clientType"`
	AppScheme      string                 `json:"appScheme"`
}

type BatchSignUrl struct {
	PperatorId  string   `json:"operatorId"`
	SignFlowIds []string `json:"signFlowIds"`
	ForcedRead  bool     `json:"forcedRead"`
	RedirectUrl string   `json:"redirectUrl"`
}
