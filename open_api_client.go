package esign

import (
	"encoding/json"
	"fmt"

	"github.com/Twistt1109/esign-v3api-go-sdk/model"
	"github.com/Twistt1109/esign-v3api-go-sdk/model/request"
	"github.com/Twistt1109/esign-v3api-go-sdk/model/response"
)

type OpenApiClient struct {
	appId     string
	appSecret string
	serverUrl string
}

const (
	grantType = "client_credentials" //客户端凭证

	tokenUrl   = "/v1/oauth2/access_token"
	refreshUrl = "/v1/oauth2/refresh_token"

	fileUploadUrl = "/v3/files/file-upload-url" //获取文件上传地址
	fileStatusUrl = "/v3/files/%s"              //获取文件上传状态

	keywordPositionsUrl = "/v3/files/%s/keyword-positions" //获取文件关键字位置

	createByFileUrl = "/v3/sign-flow/create-by-file" //创建签署流程
	signUrl         = "/v3/sign-flow/%s/sign-url"    //获取签署地址

	orgInfoUrl = "/v3/organizations/identity-info" //获取企业身份信息

	psnAuthUrl      = "/v3/psn-auth-url"          //获取个人身份认证地址
	identityInfoUrl = "/v3/persons/identity-info" // 查询个人认证信息

	ownSealList        = "/v3/seals/org-own-seal-list?orgId=%s&pageNum=%d&pageSize=%d"        // 获取企业自己印章列表
	authorizedSealList = "/v3/seals/org-authorized-seal-list?orgId=%s&pageNum=%d&pageSize=%d" // 获取其他企业授权的印章列表

	fileDownloadUrl = "/v3/sign-flow/%s/file-download-url" // 获取已签署文件下载地址

	batchSignUrl = "/v3/sign-flow/batch-sign-url" // 获取批量签页面链接（多流程）
)

func NewClient(AppID, AppSecret, ServerUrl string) *OpenApiClient {
	return &OpenApiClient{
		appId:     AppID,
		appSecret: AppSecret,
		serverUrl: ServerUrl,
	}
}

func (o *OpenApiClient) GetAccessToken() model.AccessTokenRes {
	var accessTokenRes model.AccessTokenRes

	requestUrl := o.serverUrl + tokenUrl
	url := fmt.Sprintf("%s?appId=%s&secret=%s&grantType=%s", requestUrl, o.appId, o.appSecret, grantType)
	rspBody, _ := o.SendCommHttp(url, "", "GET", "")

	err := json.Unmarshal(rspBody, &accessTokenRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return accessTokenRes
}

func (o *OpenApiClient) GetFileUploadUrl(request request.GetUploadFileUrl, token string) response.GetUploadFileUrl {
	var uploadFileUrlRes response.GetUploadFileUrl

	reqStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println("json序列化失败")
	}

	requestUrl := o.serverUrl + fileUploadUrl
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)

	err = json.Unmarshal(rspBody, &uploadFileUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return uploadFileUrlRes
}

func (o *OpenApiClient) GetFileUploadStatus(fileId, token string) response.GetFileUploadStatus {
	var fileUploadStatus response.GetFileUploadStatus

	requestUrl := o.serverUrl + fmt.Sprintf(fileStatusUrl, fileId)
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &fileUploadStatus)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return fileUploadStatus
}

func (o *OpenApiClient) GetKeywordPositions(keywords *request.Keyword, fileId, token string) response.KeywordPositionRes {
	var keywordPositionsRes response.KeywordPositionRes

	reqStr, err := json.Marshal(keywords)
	if err != nil {
		fmt.Println("json序列化失败")
	}

	requestUrl := o.serverUrl + fmt.Sprintf(keywordPositionsUrl, fileId)
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)
	err = json.Unmarshal(rspBody, &keywordPositionsRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return keywordPositionsRes
}

func (o *OpenApiClient) CreateSignTask(req *request.CreateSignTask, token string) response.CreateSignTask {
	var createSignTaskRes response.CreateSignTask

	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}

	requestUrl := o.serverUrl + createByFileUrl
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)

	err = json.Unmarshal(rspBody, &createSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return createSignTaskRes

}

func (o *OpenApiClient) GetOrgInfo(orgId, token string) response.OrgInfoRes {
	var orgInfoRes response.OrgInfoRes

	requestUrl := o.serverUrl + orgInfoUrl + "?orgId=" + orgId
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &orgInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return orgInfoRes
}

// GetPsnAuthUrl 获取个人身份认证地址
func (o *OpenApiClient) GetPsnAuthUrl(req *request.PsnAuth, token string) response.PsnAuthRes {
	var psnAuthRes response.PsnAuthRes

	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}

	requestUrl := o.serverUrl + psnAuthUrl
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)

	err = json.Unmarshal(rspBody, &psnAuthRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return psnAuthRes
}

// GetIdentityInfo 查询个人认证信息
func (o *OpenApiClient) GetIdentityInfo(psnAccount, token string) response.IdentityInfo {
	var identityInfoRes response.IdentityInfo

	requestUrl := o.serverUrl + identityInfoUrl + "?psnAccount=" + psnAccount
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &identityInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return identityInfoRes
}

func (o *OpenApiClient) GetSignUrl(req *request.SignUrl, signFlowId, token string) response.SignUrl {
	var signUrlRes response.SignUrl

	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}

	requestUrl := o.serverUrl + fmt.Sprintf(signUrl, signFlowId)
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)

	err = json.Unmarshal(rspBody, &signUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return signUrlRes
}

// 获取ownSealList
func (o *OpenApiClient) GetOwnSealList(orgId, token string) response.SealRes {
	var ownSealListRes response.SealRes

	requestUrl := o.serverUrl + fmt.Sprintf(ownSealList, orgId, 1, 20) + "&sealBizTypes=PUBLIC,CONTRACT"
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &ownSealListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return ownSealListRes
}

// GetAuthorizedSealList 获取其他企业授权的印章列表
func (o *OpenApiClient) GetAuthorizedSealList(orgId, token string) response.SealRes {
	var authorizedSealListRes response.SealRes

	requestUrl := o.serverUrl + fmt.Sprintf(authorizedSealList, orgId, 1, 20) + "&sealBizTypes=PUBLIC,CONTRACT"
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &authorizedSealListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return authorizedSealListRes
}

func (o *OpenApiClient) GetFileDownloadUrl(signFlowId, token string) response.FileDownload {
	var fileDownloadRes response.FileDownload

	requestUrl := o.serverUrl + fmt.Sprintf(fileDownloadUrl, signFlowId)
	rspBody, _ := o.SendCommHttp(requestUrl, "", "GET", token)

	err := json.Unmarshal(rspBody, &fileDownloadRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}

	return fileDownloadRes

}

func (o *OpenApiClient) BatchSignUrl(req *request.BatchSignUrl, token string) response.BatchSignUrl {
	var batchSignUrlRes response.BatchSignUrl
	reqStr, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	requestUrl := o.serverUrl + batchSignUrl
	rspBody, _ := o.SendCommHttp(requestUrl, string(reqStr), "POST", token)
	err = json.Unmarshal(rspBody, &batchSignUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	return batchSignUrlRes
}
