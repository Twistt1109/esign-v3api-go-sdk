package demo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Twistt1109/esign-v3api-go-sdk"
	"github.com/Twistt1109/esign-v3api-go-sdk/model/request"
	tools "github.com/Twistt1109/esign-v3api-go-sdk/utils"
)

const (
	KEYWORD_JIA    = "加"
	KEYWORD_DATE   = "日期"
	KEYWORD_SIGN   = "签名"
	E_CODE_SUCCESS = 200

	ESIGN_HOST         = "https://smlopenapi.esign.cn"              // e签宝测试地址
	ESIGN_APP_ID       = "7439061373"                               // 应用ID
	ESIGN_APP_SECRET   = "afea0d8fbeb2aade0eaac266f57b6491"         // 应用秘钥
	ESIGN_ORGID        = "6fbe694b3d784a4d934f8aabbf316277"         // 企业自身ID
	ESIGN_NOTIFY_URI   = "http://api.xxxx.com/contarct/id/notify"   // 签署成功回调通知地址
	ESIGN_REDIRECT_URI = "http://h5.xxxx.com/contract/sign-success" // 签署成功跳转地址
	E_SIGN_FLOW_ID     = "esign_sign_flow_id_%s"
)

type SvipSign interface {
	Download(signFlowId string) (string, error)
	GetVipSignUrl(phone, fileName, filePath, companyName, ucid string) (string, error)
}

// E签宝
type eSign struct {
	client *esign.OpenApiClient
	token  string
}

func NewESign(client *esign.OpenApiClient, token string) SvipSign {
	return &eSign{client: client, token: token}
}

// 返回认证链接, 或者已经认证的跳转签署流程
func (s *eSign) Download(signFlowId string) (string, error) {
	token, err := s.getToken()
	if err != nil {
		return "", err
	}

	fileDownloadRes := s.client.GetFileDownloadUrl(signFlowId, token)
	printInfo("获取获取文件下载地址", fileDownloadRes)

	if fileDownloadRes.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取获取文件下载地址失败: %s", fileDownloadRes.Message))
	}

	// 创建一个HTTP GET请求
	resp, err := http.Get(fileDownloadRes.Data.Files[0].DownloadUrl)
	if err != nil {
		fmt.Printf("发送请求时出错: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态码是否为200（成功）
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
		return "", nil
	}

	tempPath := fmt.Sprintf("upload/user/%s", time.Now().Format("2006/01"))

	err = os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fileDownloadRes.Data.Files[0].FileName
	outPutPath := fmt.Sprintf("%s/%s", tempPath, fileName)
	// 创建一个本地文件用于保存下载的内容
	file, err := os.Create(outPutPath)
	if err != nil {
		fmt.Printf("创建文件时出错: %v\n", err)
		return "", err
	}
	defer file.Close()

	// 将响应体内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("写入文件时出错: %v\n", err)
		return "", err
	}

	fmt.Println("文件下载成功！")
	// redisCli.Del(fmt.Sprintf(E_SIGN_FLOW_ID, fileName))

	return outPutPath, nil
}

func (s *eSign) GetVipSignUrl(phone, fileName, filePath, companyName, ucid string) (string, error) {

	_, err := s.getToken()
	if err != nil {
		return "", err
	}

	// 个人id
	persionPsnId, err := s.getPersionPsnId(phone)
	if err != nil {
		return "", err
	}

	fileName = fileName + ".pdf"
	filePath = "./" + filePath

	// 该协议如果已经发起过签署, 则直接获取签署地址
	// key := fmt.Sprintf(E_SIGN_FLOW_ID, fileName)
	// signFlowId, err := redisCli.Get(key).Result()
	// if err == nil && signFlowId != "" {
	// 	url, err := s.getSignUrl(persionPsnId, signFlowId)
	// 	if err == nil {
	// 		fmt.Println("--------------------缓存过的签署协议")
	// 		return url, nil
	// 	}
	// }

	// 文件id
	fileID, err := s.getFileId(fileName, filePath)
	if err != nil {
		return "", err
	}

	// 获取关键字位置
	signFieldPosition, err := s.getKeywordPositions(fileID)
	if err != nil {
		return "", err
	}

	// 获取印章
	sealId, err := s.getSealId(companyName)
	if err != nil {
		return "", err
	}

	// 发起签署
	createSignTask := s.client.CreateSignTask(&request.CreateSignTask{
		Docs: []request.Document{
			{
				FileId:   fileID,
				FileName: fileName,
			},
		},
		SignFlowConfig: request.SignFlowConfig{
			SignFlowTitle: fileName,
			AutoStart:     true,
			AutoFinish:    true,
			NotifyUrl:     ESIGN_NOTIFY_URI,
			SignConfig: request.SignConfig{
				AvailableSignClientTypes: "1", // 签署终端类型 - 自己的app
			},
			NoticeConfig: request.NoticeConfig{
				NoticeTypes: "1", // 通知类型 - 短信通知
			},
			AuthConfig: request.AuthConfig{
				WillingnessAuthModes: []string{"CODE_SMS"}, // 签署意愿认证方式 - 短信
			},
		},
		Signers: []request.Signer{
			{
				SignConfig: request.SignerSignConfig{
					SignOrder: 2, // 签署顺序
				},
				NoticeConfig: request.NoticeConfig{
					NoticeTypes: "1", // 通知类型 - 短信通知
				},
				SignerType: 1, // 0 - 个人，1 - 企业/机构，2 - 法定代表人，3 - 经办人
				SignFields: []request.SignField{
					{
						FileId:        fileID,
						SignFieldType: 0, // 签署控件类型 - 企业印章
						NormalSignFieldConfig: &request.NormalSignFieldConfig{
							AutoSign:       true,   // 是否自动签署
							AssignedSealId: sealId, // 指定印章id
							// OrgSealBizTypes:   "PUBLIC,CONTRACT", // 印章类型 - 企业公章 合同章
							SignFieldStyle:    tools.NewIntPtr(1), // 1 - 单页签章，2 - 骑缝签章
							SignFieldPosition: signFieldPosition[KEYWORD_JIA][0],
						},
						SignDateConfig: request.SignDateConfig{
							ShowSignDate:      1, // 是否显示签署日期
							SignDatePositionX: signFieldPosition[KEYWORD_DATE][0].PositionX - float64(50),
							SignDatePositionY: signFieldPosition[KEYWORD_DATE][0].PositionY,
						},
					},
				},
			},
			{
				SignConfig: request.SignerSignConfig{
					SignOrder: 1, // 签署顺序
				},
				NoticeConfig: request.NoticeConfig{
					NoticeTypes: "1", // 通知类型 - 短信通知
				},
				SignerType: 0,
				PsnSignerInfo: &request.PsnSignerInfo{
					PsnId: persionPsnId, // 个人id
				},
				SignFields: []request.SignField{
					{
						FileId:        fileID,
						SignFieldType: 0, // 签署控件类型 - 企业印章 / 手写签名
						NormalSignFieldConfig: &request.NormalSignFieldConfig{
							PsnSealStyles:     "0", // 手写签名
							SignFieldPosition: signFieldPosition[KEYWORD_SIGN][0],
							SignFieldStyle:    tools.NewIntPtr(1), // 1 - 单页签章，2 - 骑缝签章
						},
						SignDateConfig: request.SignDateConfig{
							ShowSignDate:      1, // 是否显示签署日期
							SignDatePositionX: signFieldPosition[KEYWORD_DATE][1].PositionX - float64(50),
							SignDatePositionY: signFieldPosition[KEYWORD_DATE][1].PositionY,
						},
					},
				},
			},
		},
	}, s.token)
	if createSignTask.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("签署发起失败: %s", createSignTask.Message))
	}
	printInfo("签署发起", createSignTask)

	signFlowId := createSignTask.Data.SignFlowId
	// redisCli.Set(key, signFlowId, 0)

	return s.getSignUrl(persionPsnId, signFlowId)
}

func (s *eSign) getSignUrl(persionPsnId, signFlowId string) (string, error) {
	signUrl := s.client.GetSignUrl(&request.SignUrl{
		Operator: request.Operator{PsnId: persionPsnId},
		RedirectConfig: &request.SignUrlRedirectConfig{
			RedirectUrl:       ESIGN_REDIRECT_URI, // 签署成功后跳转地址
			RedirectDelayTime: 0,
		},
	}, signFlowId, s.token)

	if signUrl.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取签署地址失败: %s", signUrl.Message))
	}

	return signUrl.Data.Url, nil
}

// esignKeywordsConversion 将关键字转换为定位控件信息 keywords-关键字定位信息
func (s *eSign) getKeywordPositions(fileID string) (map[string][]request.SignFieldPosition, error) {
	keywordMap := make(map[string][]request.SignFieldPosition)

	// 获取关键字位置
	keywordPositionRes := s.client.GetKeywordPositions(&request.Keyword{
		Keywords: []string{KEYWORD_JIA, KEYWORD_SIGN, KEYWORD_DATE},
	}, fileID, s.token)
	if keywordPositionRes.Code != E_CODE_SUCCESS {
		return keywordMap, errors.New(fmt.Sprintf("获取获取关键字位置失败: %s", keywordPositionRes.Message))
	}
	// printInfo("获取关键字位置", keywordPositionRes)

	keywords := keywordPositionRes.Data.KeywordPositions

	// 定义一个map, key是关键字, value是一个切片, 用来存储对应的坐标信息
	for _, v := range keywords {

		for i, position := range v.Positions {

			if (v.Keyword == KEYWORD_JIA && i != len(v.Positions)-1) || v.SearchResult == false {
				continue // 只需要最后一个甲方关键字
			}

			fmt.Println("----------关键字坐标: ", v.Keyword, position)

			for _, coordinate := range position.Coordinates {
				keywordMap[v.Keyword] = append(keywordMap[v.Keyword], request.SignFieldPosition{
					PositionPage: strconv.FormatInt(int64(position.PageNum), 10),
					PositionX:    coordinate.PositionX + float64(130),
					PositionY:    coordinate.PositionY,
				})
			}
		}
	}

	// jsonData, _ := json.MarshalIndent(keywordMap, "", "  ")
	// fmt.Println("----------文档中添加控件坐标:\n", string(jsonData))

	return keywordMap, nil
}

// 获取token 有效期2小时
func (s *eSign) getToken() (string, error) {
	accessTokenRes := s.client.GetAccessToken()
	if accessTokenRes.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取token失败: %s", accessTokenRes.Message))
	}
	token := accessTokenRes.Data.Token
	s.token = token

	fmt.Println("----------token", token)
	return token, nil

	// 暂不使用缓存
	// key := "esign_token_" + app.Config.EsignAppid

	// token, err := redisCli.Get(key).Result()
	// if err == redis.Nil {
	// 	accessTokenRes := s.client.GetAccessToken()
	// 	if accessTokenRes.Code != E_CODE_SUCCESS {
	// 		return "", errors.New(fmt.Sprintf("获取token失败: %s", accessTokenRes.Message))
	// 	}
	// 	token = accessTokenRes.Data.Token
	// 	redisCli.Set(key, token, 110*time.Minute)
	// }

	// s.token = token

	// fmt.Println("----------token", token)
	// return token, nil
}

// getPersionPsnId 通过手机号获取 个人标识
func (s *eSign) getPersionPsnId(phone string) (string, error) {
	getIdentityInfo := s.client.GetIdentityInfo(phone, s.token)
	printInfo("1---查询个人认证信息", getIdentityInfo)
	if getIdentityInfo.Code == E_CODE_SUCCESS {
		return getIdentityInfo.Data.PsnId, nil
	}

	// 没有个人认证则, 获取个人认证&授权页面链接
	s.client.GetPsnAuthUrl(&request.PsnAuth{
		PsnAuthConfig: request.PsnAuthConfig{
			PsnAccount: phone,
		},
	}, s.token)

	getIdentityInfo = s.client.GetIdentityInfo(phone, s.token)
	printInfo("2---查询个人认证信息", getIdentityInfo)
	if getIdentityInfo.Code == E_CODE_SUCCESS {
		return getIdentityInfo.Data.PsnId, nil
	}

	return "", errors.New(fmt.Sprintf("查询个人认证信息失败: %s", getIdentityInfo.Message))
}

// getFileId 获取 文件id
func (s *eSign) getFileId(fileName, filePath string) (string, error) {
	contentMd5, size := tools.CountFileMd5(filePath)
	contentType := "application/pdf"

	// 1. 获取文件上传地址
	getUploadFileUrl := s.client.GetFileUploadUrl(request.GetUploadFileUrl{
		ContentMd5:  contentMd5,
		ContentType: contentType,
		FileName:    fileName,
		FileSize:    size,
	}, s.token)
	if getUploadFileUrl.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取获取文件上传地址失败: %s", getUploadFileUrl.Message))
	}

	fileID := getUploadFileUrl.Data.FileId // 文件id

	printInfo("获取文件上传地址", getUploadFileUrl)

	// 2. 上传合同文件
	result := esign.UpLoadFile(getUploadFileUrl.Data.FileUploadUrl, filePath, contentMd5, contentType)
	printInfo("上传文件", result)

	// 获取文件上传状态
	getFileUploadStatus := s.client.GetFileUploadStatus(fileID, s.token)
	if getFileUploadStatus.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取获取文件上传状态失败: %s", getFileUploadStatus.Message))
	}
	printInfo("获取文件上传状态", getFileUploadStatus)

	fmt.Println("----------fileID", fileID)
	return fileID, nil
}

func (s *eSign) getSealId(companyName string) (string, error) {
	sealId := ""

	ownSealList := s.client.GetOwnSealList(ESIGN_ORGID, s.token)
	printInfo("获取本企业印章列表", ownSealList)
	if ownSealList.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取印章列表失败: %s", ownSealList.Message))
	}
	for _, v := range ownSealList.Data.Seals {
		if v.SealName == companyName {
			sealId = v.SealId
			break
		}
	}

	if sealId != "" {
		return sealId, nil
	}

	authorizedSealList := s.client.GetAuthorizedSealList(ESIGN_ORGID, s.token)
	printInfo("获取其他企业印章列表", authorizedSealList)
	if authorizedSealList.Code != E_CODE_SUCCESS {
		return "", errors.New(fmt.Sprintf("获取印章列表失败: %s", authorizedSealList.Message))
	}

	for _, v := range authorizedSealList.Data.Seals {
		if v.SealName == companyName {
			sealId = v.SealId
			break
		}
	}

	if sealId != "" {
		return sealId, nil
	}

	return "", errors.New("没有该公司印章")
}

func printInfo(msg string, v any) {
	jsonData, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("----------%s:%s\n", msg, string(jsonData))
}
