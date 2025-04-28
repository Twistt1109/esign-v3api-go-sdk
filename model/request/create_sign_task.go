package request

type CreateSignTask struct {
	Docs              []Document         `json:"docs"`
	Attachments       []Document         `json:"attachments"`
	SignFlowConfig    SignFlowConfig     `json:"signFlowConfig"`
	SignFlowInitiator *SignFlowInitiator `json:"signFlowInitiator"`
	Signers           []Signer           `json:"signers"`
	Copiers           []Copier           `json:"copiers"`
}

type Document struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
}

type SignFlowConfig struct {
	SignFlowTitle      string          `json:"signFlowTitle"`
	SignFlowExpireTime int64           `json:"signFlowExpireTime,omitempty"`
	AutoStart          bool            `json:"autoStart"`
	AutoFinish         bool            `json:"autoFinish"`
	IdentityVerify     bool            `json:"identityVerify"`
	SignConfig         SignConfig      `json:"signConfig"`
	NotifyUrl          string          `json:"notifyUrl"`
	NoticeConfig       NoticeConfig    `json:"noticeConfig"`
	RedirectConfig     *RedirectConfig `json:"redirectConfig"`
	AuthConfig         AuthConfig      `json:"authConfig"`
}

type SignConfig struct {
	AvailableSignClientTypes string `json:"availableSignClientTypes"`
	ShowBatchDropSealButton  bool   `json:"showBatchDropSealButton"`
}

type RedirectConfig struct {
	RedirectDelayTime string `json:"redirectDelayTime"`
	RedirectUrl       string `json:"redirectUrl"`
}

type AuthConfig struct {
	PsnAvailableAuthModes []string `json:"psnAvailableAuthModes"`
	WillingnessAuthModes  []string `json:"willingnessAuthModes"`
	OrgAvailableAuthModes []string `json:"orgAvailableAuthModes"`
}

type SignFlowInitiator struct {
	OrgInitiator OrgInitiator `json:"orgInitiator"`
}

type OrgInitiator struct {
	OrgId      string     `json:"orgId"`
	Transactor Transactor `json:"transactor"`
}

type Transactor struct {
	PsnId string `json:"psnId"`
}

type Signer struct {
	SignConfig    SignerSignConfig `json:"signConfig"`
	NoticeConfig  NoticeConfig     `json:"noticeConfig"`
	SignerType    int              `json:"signerType"`
	PsnSignerInfo *PsnSignerInfo   `json:"psnSignerInfo,omitempty"`
	OrgSignerInfo *OrgSignerInfo   `json:"orgSignerInfo,omitempty"`
	SignFields    []SignField      `json:"signFields"`
}

type SignerSignConfig struct {
	ForcedReadingTime    string `json:"forcedReadingTime"`
	AgreeSkipWillingness *bool  `json:"agreeSkipWillingness"`
	SignOrder            int    `json:"signOrder"`
	SignTaskType         int    `json:"signTaskType"`
}

type NoticeConfig struct {
	NoticeTypes string `json:"noticeTypes"`
}

type PsnSignerInfo struct {
	PsnAccount string     `json:"psnAccount"`
	PsnId      string     `json:"psnId"`
	PsnInfo    PersonInfo `json:"psnInfo"`
}

type OrgSignerInfo struct {
	OrgId          string         `json:"orgId"`
	OrgName        string         `json:"orgName"`
	OrgInfo        OrgInfo        `json:"orgInfo"`
	TransactorInfo TransactorInfo `json:"transactorInfo"`
}

type PersonInfo struct {
	PsnName       string `json:"psnName"`
	PsnIDCardNum  string `json:"psnIDCardNum"`
	PsnIDCardType string `json:"psnIDCardType"`
}

type OrgInfo struct {
	OrgIDCardNum  string `json:"orgIDCardNum"`
	OrgIDCardType string `json:"orgIDCardType"`
}

type TransactorInfo struct {
	PsnAccount string     `json:"psnAccount"`
	PsnInfo    PersonInfo `json:"psnInfo"`
}

type SignField struct {
	FileId                string                 `json:"fileId"`
	CustomBizNum          string                 `json:"customBizNum"`
	SignFieldType         int                    `json:"signFieldType"`
	NormalSignFieldConfig *NormalSignFieldConfig `json:"normalSignFieldConfig"`
	RemarkSignFieldConfig *RemarkSignFieldConfig `json:"remarkSignFieldConfig"`
	SignDateConfig        SignDateConfig         `json:"signDateConfig"`
}

type NormalSignFieldConfig struct {
	AutoSign          bool              `json:"autoSign"`
	FreeMode          bool              `json:"freeMode"`
	MovableSignField  bool              `json:"movableSignField"`
	PsnSealStyles     string            `json:"psnSealStyles"`
	AssignedSealId    string            `json:"assignedSealId"`
	OrgSealBizTypes   string            `json:"orgSealBizTypes"`
	SignFieldSize     string            `json:"signFieldSize"`
	SignFieldStyle    *int              `json:"signFieldStyle,omitempty"`
	SignFieldPosition SignFieldPosition `json:"signFieldPosition"`
}

type RemarkSignFieldConfig struct {
	InputType       int `json:"inputType"`
	SignFieldHeight int `json:"signFieldHeight"`
	SignFieldWidth  int `json:"signFieldWidth"`
	// RemarkContent   string `json:"remarkContent"`

	SignFieldPosition SignFieldPosition `json:"signFieldPosition"`
}

type SignFieldPosition struct {
	PositionPage string  `json:"positionPage"`
	PositionX    float64 `json:"positionX"`
	PositionY    float64 `json:"positionY"`
}

type SignDateConfig struct {
	DateFormat        string  `json:"dateFormat"`
	ShowSignDate      int     `json:"showSignDate"`
	SignDatePositionX float64 `json:"signDatePositionX"`
	SignDatePositionY float64 `json:"signDatePositionY"`
}

type Copier struct {
	CopierOrgInfo *CopierOrgInfo `json:"copierOrgInfo,omitempty"`
	CopierPsnInfo *CopierPsnInfo `json:"copierPsnInfo,omitempty"`
}

type CopierOrgInfo struct {
	OrgName string `json:"orgName"`
}

type CopierPsnInfo struct {
	PsnAccount string `json:"psnAccount"`
}
