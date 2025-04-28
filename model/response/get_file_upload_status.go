package response

type GetFileUploadStatus struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetFileUploadStatusData `json:"data"`
}

type GetFileUploadStatusData struct {
	FileId             string `json:"fileId"`
	FileName           string `json:"fileName"`
	FileSize           *int   `json:"fileSize"`
	FileStatus         int    `json:"fileStatus"`
	FileDownloadUrl    string `json:"fileDownloadUrl"`
	FileTotalPageCount int    `json:"fileTotalPageCount"`
	PageWidth          *int   `json:"pageWidth"`
	PageHeight         *int   `json:"pageHeight"`
}
