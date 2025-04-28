package request

type GetUploadFileUrl struct {
	ContentMd5   string `json:"contentMd5"`
	ContentType  string `json:"contentType"`
	ConvertToPDF bool   `json:"convertToPDF"`
	FileName     string `json:"fileName"`
	FileSize     int64  `json:"fileSize"`
}
