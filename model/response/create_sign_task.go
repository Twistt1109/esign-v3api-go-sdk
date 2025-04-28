package response

type CreateSignTask struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateSignTaskData `json:"data"`
}

type CreateSignTaskData struct {
	SignFlowId string `json:"signFlowId"`
}
