package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"user-id"`
	Visible     bool      `gorm:"not null" json:"-"`
	Email       string    `gorm:"not null;size:320;unique" json:"email"`
	PWHash      string    `gorm:"not null;size:60" json:"-"`
	Username    string    `gorm:"not null;size:30;unique" json:"username"`
	Displayname string    `gorm:"not null;size:30" json:"displayname"`
	Position    string    `gorm:"size:30" json:"position"`
	Description string    `gorm:"size:240" json:"description"`
	Sessions    []Session `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	TeamID      uuid.UUID `gorm:"type:uuid" json:"team-id"`
	Roles       []Role    `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"roles"`
	Projects    []Project `gorm:"many2many:project_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"projects"`
	SentInvite  []Invite  `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created-at"`
}
