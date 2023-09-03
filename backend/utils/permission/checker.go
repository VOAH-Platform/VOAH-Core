package permission

import (
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
)

func PermissionCheck(userPerms []models.Permission, requirePerms []models.Permission) bool {
	for _, requrequirePerm := range requirePerms {
		for _, userPerm := range userPerms {
			if userPerm.Type == models.RootObject || userPerm.Type == models.SystemObject {
				return true
			} else if requrequirePerm.Type == userPerm.Type && requrequirePerm.Target == userPerm.Target {
				if models.AdminPermissionScope == userPerm.Scope {
					break
				} else if requrequirePerm.Scope == userPerm.Scope {
					return true
				}
			}
		}
	}
	return false
}

func UserPermissionCheck(user models.User, requirePerms []models.Permission) (hasPerm bool, err error) {
	userRoles := new([]models.Role)
	if err := database.DB.Model(&user).Association("Roles").Find(userRoles); err != nil {
		return false, err
	}
	userPerms := new([]models.Permission)
	for _, role := range *userRoles {
		tempPermissions := new([]models.Permission)
		if err := database.DB.Model(&role).Association("Permissions").Find(tempPermissions); err != nil {
			return false, err
		}
		*userPerms = append(*userPerms, *tempPermissions...)
	}
	return PermissionCheck(*userPerms, requirePerms), nil
}
