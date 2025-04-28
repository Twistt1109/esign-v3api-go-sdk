package request

type PsnAuthConfig struct {
	PsnAccount        string            `json:"psnAccount"`
	PsnInfo           PsnInfo           `json:"psnInfo"`
	PsnAuthPageConfig PsnAuthPageConfig `json:"psnAuthPageConfig"`
}

type PsnInfo struct {
	PsnName       string `json:"psnName"`
	PsnIDCardNum  string `json:"psnIDCardNum"`
	PsnIDCardType string `json:"psnIDCardType"`
}

type PsnAuthPageConfig struct {
	PsnDefaultAuthMode    string   `json:"psnDefaultAuthMode"`
	PsnAvailableAuthModes []string `json:"psnAvailableAuthModes"`
}

type AuthorizeConfig struct {
	AuthorizedScopes []string `json:"authorizedScopes"`
}

type AuthRedirectConfig struct {
	RedirectUrl string `json:"redirectUrl"`
}

type PsnAuth struct {
	PsnAuthConfig   PsnAuthConfig      `json:"psnAuthConfig"`
	AuthorizeConfig AuthorizeConfig    `json:"authorizeConfig"`
	NotifyUrl       string             `json:"notifyUrl"`
	ClientType      string             `json:"clientType"`
	RedirectConfig  AuthRedirectConfig `json:"redirectConfig"`
}
