package response

type GetUploadFileUrl struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	FileUploadUrl string `json:"fileUploadUrl"`
	FileId        string `json:"fileId"`
}
