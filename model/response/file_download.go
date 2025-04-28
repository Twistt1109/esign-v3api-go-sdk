package response

type File struct {
	FileId      string `json:"fileId"`
	FileName    string `json:"fileName"`
	DownloadUrl string `json:"downloadUrl"`
}

type FileData struct {
	Files                  []File      `json:"files"`
	Attachments            []File      `json:"attachments"`
	CertificateDownloadUrl interface{} `json:"certificateDownloadUrl"`
}

type FileDownload struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    FileData `json:"data"`
}
