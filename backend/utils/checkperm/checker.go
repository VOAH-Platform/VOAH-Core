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
					return true
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

func GetPermissionByRoleArr(roles []models.Role) ([]models.Permission, error) {
	var err error
	perms := new([]models.Permission)
	for _, role := range roles {
		tempPermissions := new([]models.Permission)
		if err = database.DB.Model(&role).Association("Permissions").Find(tempPermissions); err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		*perms = append(*perms, *tempPermissions...)
	}
	return *perms, nil
}

func GetUserPermissionArr(user *models.User) ([]models.Permission, error) {
	var err error
	userRoles, err := GetUserRoleArr(user)
	if err != nil {
		return nil, err
	}
	return GetPermissionByRoleArr(userRoles)
}

func UserPermissionCheck(user *models.User, requirePerms []models.Permission) (hasPerm bool, err error) {
	userPerms, err := GetUserPermissionArr(user)
	if err != nil {
		return false, err
	}
	return PermissionCheck(userPerms, requirePerms), nil
}
