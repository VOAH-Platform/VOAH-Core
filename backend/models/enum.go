package models

import "github.com/google/uuid"

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
)

func (p PermissionScope) IsValid() bool {
	switch p {
	case AdminPermissionScope,
		EditPermissionScope,
		InvitePermissionScope:
		return true
	}
	return false
}

var CompanyID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")
