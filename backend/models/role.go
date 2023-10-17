package models

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey" json:"role-id"`
	Type        string       `gorm:"not null;size:30" json:"type"`
	Name        string       `gorm:"not null;size:30;unique" json:"name"`
	Description string       `gorm:"not null;size:200" json:"description"`
	Users       []User       `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"users"`
	Permissions []Permission `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
