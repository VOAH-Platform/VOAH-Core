package checkperm

import (
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
)

func PermissionCheck(userPerms []models.Permission, requirePerms []models.Permission) bool {
	for _, requrequirePerm := range requirePerms {
		for _, userPerm := range userPerms {
			if userPerm.Type == configs.RootObject || userPerm.Type == configs.SystemObject {
				return true
			} else if requrequirePerm.Type == userPerm.Type && requrequirePerm.Target == userPerm.Target {
				if configs.AdminPermissionScope == userPerm.Scope {
					break
				} else if requrequirePerm.Scope == userPerm.Scope {
					return true
				}
			}
		}
	}
	return false
}

func GetUserRoleArr(user *models.User) ([]models.Role, error) {
	userRoles := new([]models.Role)
	if err := database.DB.Model(&user).Association("Roles").Find(userRoles); err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return *userRoles, nil
}

func GetUserPermissionArr(user *models.User) ([]models.Permission, error) {
	userRoles, err := GetUserRoleArr(user)
	if err != nil {
		return nil, err
	}

	userPerms := new([]models.Permission)
	for _, role := range userRoles {
		tempPermissions := new([]models.Permission)
		if err := database.DB.Model(&role).Association("Permissions").Find(tempPermissions); err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		*userPerms = append(*userPerms, *tempPermissions...)
	}
	return *userPerms, nil
}

func UserPermissionCheck(user *models.User, requirePerms []models.Permission) (hasPerm bool, err error) {
	userPerms, err := GetUserPermissionArr(user)
	if err != nil {
		return false, err
	}
	return PermissionCheck(userPerms, requirePerms), nil
}
