package response

type OrgInfoRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    OrgInfoData `json:"data"`
}

type OrgInfoData struct {
	AuthorizeUserInfo bool    `json:"authorizeUserInfo"`
	RealnameStatus    int     `json:"realnameStatus"`
	OrgId             string  `json:"orgId"`
	OrgName           string  `json:"orgName"`
	OrgAuthMode       string  `json:"orgAuthMode"`
	OrgInfo           OrgInfo `json:"orgInfo"`
}

type OrgInfo struct {
	OrgType                  *string `json:"orgType"`
	OrgIDCardNum             string  `json:"orgIDCardNum"`
	OrgIDCardType            string  `json:"orgIDCardType"`
	LegalRepName             string  `json:"legalRepName"`
	LegalRepIDCardNum        *string `json:"legalRepIDCardNum"`
	LegalRepIDCardType       *string `json:"legalRepIDCardType"`
	OrgBankAccountNum        *string `json:"orgBankAccountNum"`
	CorporateAccount         *string `json:"corporateAccount"`
	CnapsCode                *string `json:"cnapsCode"`
	LicenseDownloadUrl       *string `json:"licenseDownloadUrl"`
	AuthorizationDownloadUrl *string `json:"authorizationDownloadUrl"`
	AdminName                string  `json:"adminName"`
	AdminAccount             string  `json:"adminAccount"`
}
