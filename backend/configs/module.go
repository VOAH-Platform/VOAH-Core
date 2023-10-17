package configs

import "github.com/google/uuid"

const (
	ModuleID          = 0
	ModuleName        = "VOAH-Template-Project"
	ModuleVersion     = "0.0.1"
	ModuleDescription = "VOAH Template Project"
)

var (
	ModuleDeps = []string{}
)

type ObjectType string

const (
	SystemObject  ObjectType = "system"
	RootObject    ObjectType = "root"
	ProjectObject ObjectType = "project"
	TeamObject    ObjectType = "team"
	CompanyObject ObjectType = "company"
)

func (p ObjectType) IsValid() bool {
	switch p {
	case SystemObject,
		RootObject,
		ProjectObject,
		TeamObject,
		CompanyObject:
		return true
	}
	return false
}

type PermissionScope string

const (
	AdminPermissionScope  PermissionScope = "admin"
	EditPermissionScope   PermissionScope = "edit"
	InvitePermissionScope PermissionScope = "invite"
	ReadPermissionScope   PermissionScope = "read"
)

func (p PermissionScope) IsValid() bool {
	switch p {
	case AdminPermissionScope,
		EditPermissionScope,
		InvitePermissionScope,
		ReadPermissionScope:
		return true
	}
	return false
}

var CompanyID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")

var (
	AllObjectsTypes = []interface{}{}
)

var APIKey string
