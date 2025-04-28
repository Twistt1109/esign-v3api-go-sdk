package response

type IdentityInfo struct {
	Code      int              `json:"code"`
	ErrorCode *int             `json:"error_code"`
	Message   string           `json:"message"`
	Data      IdentityInfoData `json:"data"`
}

type IdentityInfoData struct {
	AuthorizeUserInfo bool       `json:"authorizeUserInfo"`
	RealnameStatus    int        `json:"realnameStatus"`
	PsnId             string     `json:"psnId"`
	PsnAccount        PsnAccount `json:"psnAccount"`
	PsnInfo           PsnInfo    `json:"psnInfo"`
}

type PsnAccount struct {
	AccountMobile string  `json:"accountMobile"`
	AccountEmail  *string `json:"accountEmail"`
}

type PsnInfo struct {
	PsnName        string  `json:"psnName"`
	PsnNationality *string `json:"psnNationality"`
	PsnIDCardNum   string  `json:"psnIDCardNum"`
	PsnIDCardType  string  `json:"psnIDCardType"`
	BankCardNum    *string `json:"bankCardNum"`
	PsnMobile      string  `json:"psnMobile"`
}
