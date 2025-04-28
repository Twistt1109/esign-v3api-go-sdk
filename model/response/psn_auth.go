package response

type PsnAuthRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    PsnAuthData `json:"data"`
}

type PsnAuthData struct {
	AuthFlowId   string `json:"authFlowId"`
	AuthUrl      string `json:"authUrl"`
	AuthShortUrl string `json:"authShortUrl"`
}
