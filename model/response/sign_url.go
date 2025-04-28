package response

type SignUrl struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    SignUrlData `json:"data"`
}

type SignUrlData struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

type BatchSignUrlData struct {
	BatchSerialId                 string `json:"batchSerialId"`
	BatchSignUrl                  string `json:"batchSignUrl"`
	BatchSignShortUrl             string `json:"batchSignShortUrl"`
	BatchSignUrlWithoutLogin      string `json:"batchSignUrlWithoutLogin"`
	BatchSignShortUrlWithoutLogin string `json:"batchSignShortUrlWithoutLogin"`
}

type BatchSignUrl struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    BatchSignUrlData `json:"data"`
}
