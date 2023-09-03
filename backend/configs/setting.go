package configs

var Setting *MainSetting

type companySetting struct {
	Name        string `json:"name"`        // Company name
	Description string `json:"description"` // Company description
	Domain      string `json:"domain"`      // Company domain
}

type authSetting struct {
	AllowRegister            bool   `json:"allowRegister"` // Allow user to register
	AllowOnlyDomain          bool   `json:"allowOnlyDomain"`
	EmailVerificattionExpire int    `json:"emailVerificationExpire"`  // Email verification code expire time in minutes
	VerificationEmailSubject string `json:"verificationEmailSubject"` // verification email subject
	VerificationEmailBody    string `json:"verificationEmailBody"`    //  verification email body
}

type MainSetting struct {
	Company companySetting `json:"company"` // Company setting
	Auth    authSetting    `json:"auth"`    // Auth setting
}
