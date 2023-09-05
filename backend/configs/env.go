package configs

var Env *MainEnv

type databaseEnv struct {
	Host     string // Database host
	Port     int    // Database port
	User     string // Database user
	Password string // Database password
	DBName   string // Database name
}

type redisEnv struct {
	Host             string // Redis host
	Port             int    // Redis port
	Password         string // Redis password
	SessionDB        int    // Redis DB for store Session(Refresh token)
	LastActivityDB   int    // Redis DB for store LastActivity time
	LastRefreshDB    int    // Redis DB for store LastRefresh time
	PasswordResetDB  int    // Redis DB for store PasswordReset code
	RegisterVerifyDB int    // Redis DB for store RegisterVerify code
}

type serverEnv struct {
	Host       string // ex) 0.0.0.0
	HostURL    string // ex) http://localhost:3000
	Port       int
	CSRFOrigin string
	DataDir    string
}

type authEnv struct {
	JWTExpire int
	JWTSecret []byte
}

type smtpEnv struct {
	Host               string
	Port               int
	Username           string
	Password           string
	SSL                bool
	STARTTLS           bool
	InsecureSkipVerify bool
	SystemAddress      string
}

type rootUserEnv struct {
	Email  string
	PWHash string // bcrypt hash
}

type MainEnv struct {
	Database databaseEnv
	Redis    redisEnv
	Server   serverEnv
	Auth     authEnv
	SMTP     smtpEnv
	RootUser rootUserEnv
}
