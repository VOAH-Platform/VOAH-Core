package module

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/logger"
)

func InitModules() {
	var err error

	log := logger.Logger
	db := database.DB
	allModules := []models.Module{}
	if err = db.Find(&allModules).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
	}

	hashedAPIKeyByte, err := bcrypt.GenerateFromPassword([]byte(configs.APIKey), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	hashedAPIKey := string(hashedAPIKeyByte)

	for _, module := range allModules {
		if !module.Enabled {
			log.Info("Module " + module.Name + " is skipped, because it is disabled")
		} else {
			client := resty.New()
			resp, err := client.R().
				SetHeader("HASHED-API-KEY", hashedAPIKey).
				Get(module.HostInternalURL + "/api/info/init")
			if err != nil {
				module.Enabled = false
				db.Save(&module)
				log.Error("Failed to get module info from " + module.Name)
				log.Error(err.Error())
			} else if resp.StatusCode() != 200 || string(resp.Body()) != configs.APIKey {
				module.Enabled = false
				db.Save(&module)
				log.Error("[CRITICAL] Failed to verify module " + module.Name)
				log.Error("Status code: " + fmt.Sprint(resp.StatusCode()))
				log.Error("Response body: " + string(resp.Body()))
			} else {
				log.Info("Module " + module.Name + " is enabled")
			}
		}
	}
}
