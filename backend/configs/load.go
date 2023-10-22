package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/utils/logger"
)

func getEnvStr(key string, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		fmt.Printf("Environment variable %s is not set, Keep Going with Default Value '%s' \n", key, defaultValue)
		return defaultValue
	}
	return
}

func getEnvInt(key string, defaultValue int) (intValue int) {
	value := getEnvStr(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return
}

func LoadEnv() {
	var jwtSecretStr = getEnvStr("AUTH_JWT_SECRET", "d495ce948d89f228cf4e")

	Env = &MainEnv{
		Database: databaseEnv{
			Host:     getEnvStr("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnvStr("DB_USER", "postgres"),
			Password: getEnvStr("DB_PASSWORD", "postgres"),
			DBName:   getEnvStr("DB_NAME", "voah-core-db"),
		},
		Redis: redisEnv{
			Host:             getEnvStr("REDIS_HOST", "localhost"),
			Port:             getEnvInt("REDIS_PORT", 6379),
			Password:         getEnvStr("REDIS_PASSWORD", "redis"),
			SessionDB:        getEnvInt("REDIS_SESSION_DB", 0),
			LastActivityDB:   getEnvInt("REDIS_LAST_ACTIVITY_DB", 1),
			LastRefreshDB:    getEnvInt("REDIS_LAST_REFRESH_DB", 2),
			PasswordResetDB:  getEnvInt("REDIS_PASSWORD_RESET_DB", 3),
			RegisterVerifyDB: getEnvInt("REDIS_REGISTER_VERIFY_DB", 4),
		},
		Server: serverEnv{
			Host:       getEnvStr("SERVER_HOST", "0.0.0.0"),
			HostURL:    getEnvStr("SERVER_HOST_URL", "http://localhost:3000"),
			Port:       getEnvInt("SERVER_PORT", 3000),
			CSRFOrigin: getEnvStr("SERVER_CSRF_ORIGIN", "*"),
			DataDir:    getEnvStr("SERVER_DATA_DIR", "./data"),
		},
		Auth: authEnv{
			JWTExpire: getEnvInt("AUTH_JWT_EXPIRE", 3600),
			JWTSecret: []byte(jwtSecretStr),
		},
		SMTP: smtpEnv{
			Host:               getEnvStr("SMTP_HOST", "smtp.gmail.com"),
			Port:               getEnvInt("SMTP_PORT", 587),
			Username:           getEnvStr("SMTP_USERNAME", ""),
			Password:           getEnvStr("SMTP_PASSWORD", ""),
			SSL:                getEnvStr("SMTP_SSL", "false") == "true",
			STARTTLS:           getEnvStr("SMTP_STARTTLS", "false") == "true",
			InsecureSkipVerify: getEnvStr("SMTP_INSECURE_SKIP_VERIFY", "false") == "true",
			SystemAddress:      getEnvStr("SMTP_SYSTEM_ADDRESS", ""),
		},
		RootUser: rootUserEnv{
			Email:  getEnvStr("VOAH_ROOT_EMAIL", "root@example.com"),
			PWHash: getEnvStr("VOAH_ROOT_PW_HASH", "$2a$12$Ca5jWnI0VgviBleFVr8PLOyKI5S8QX2mybsTNPi2dDULrWgVVG9uW"),
		},
	}
}

func LoadSetting() {
	jsonFile, err := os.ReadFile(fmt.Sprintf("%s/setting.json", Env.Server.DataDir))
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(jsonFile, &Setting) != nil {
		panic(err)
	}
}

func LoadAPIKey(wait *sync.WaitGroup) {
	log := logger.Logger

	//if not found api.key, create new one
	if _, err := os.Stat(fmt.Sprintf("%s/api.key", Env.Server.DataDir)); os.IsNotExist(err) {
		apiKey := uuid.New().String()
		err = os.WriteFile(fmt.Sprintf("%s/api.key", Env.Server.DataDir), []byte(apiKey), 0744)
		if err != nil {
			log.Fatal(err)
		}
	}

	apiKeyByte, err := os.ReadFile(fmt.Sprintf("%s/api.key", Env.Server.DataDir))
	if err != nil {
		log.Fatal(err)
	}
	apiKey := string(apiKeyByte)

	APIKey = apiKey
	log.Info("Loaded API Key")
	defer wait.Done()
}
