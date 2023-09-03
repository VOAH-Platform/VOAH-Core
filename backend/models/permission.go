package models

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID     uuid.UUID       `gorm:"type:uuid;primaryKey" json:"permission-id"`
	Type   ObjectType      `gorm:"not null;size:30" json:"type"`
	Target uuid.UUID       `gorm:"not null" json:"target"`
	Scope  PermissionScope `gorm:"not null;size:50" json:"scope"`
	RoleID uuid.UUID       `gorm:"type:uuid;not null" json:"role-id"`
}
