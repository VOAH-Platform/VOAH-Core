package rootuser

import (
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/logger"
)

func InitRootUser() {
	db := database.DB

	// find root user is exists
	rootUser := new(models.User)
	if err := db.Where(&models.User{Email: configs.Env.RootUser.Email}).First(&rootUser).Error; err != nil {

		log := logger.Logger

		// create root team
		rootTeamID := uuid.New()
		newRootTeam := &models.Team{
			ID:          rootTeamID,
			Public:      false,
			Displayname: "root",
			Description: "root team",
		}
		if err := db.Create(newRootTeam).Error; err != nil {
			log.Error("Failed to create root team")
			log.Fatal(err)
		}
		// create root user
		newRootUser := &models.User{
			ID:       uuid.New(),
			Visible:  false,
			Email:    configs.Env.RootUser.Email,
			PWHash:   configs.Env.RootUser.PWHash,
			Username: "root",
			TeamID:   rootTeamID,
		}
		if err := db.Create(newRootUser).Error; err != nil {
			log.Error("Failed to create root user")
			log.Fatal(err)
		}
		// create root role
		roleID := uuid.New()
		newRootRole := &models.Role{
			ID:          roleID,
			Displayname: "root",
			Description: "root role",
			Users:       []models.User{*newRootUser},
		}
		if err := db.Create(newRootRole).Error; err != nil {
			log.Error("Failed to create root role")
			log.Fatal(err)
		}
		// create root permission
		newRootPermission := &models.Permission{
			ID:     uuid.New(),
			Type:   models.RootObject,
			Target: models.CompanyID,
			Scope:  models.AdminPermissionScope,
			RoleID: roleID,
		}
		if err := db.Create(newRootPermission).Error; err != nil {
			log.Error("Failed to create root permission")
			log.Fatal(err)
		}

		log.Info("Root user created")
	}
}
