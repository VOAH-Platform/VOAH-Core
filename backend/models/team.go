package models

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"team-id"`
	Visible     bool      `gorm:"not null" json:"-"`
	Public      bool      `gorm:"not null" json:"public"`
	Displayname string    `gorm:"not null;size:30" json:"displayname"`
	Description string    `gorm:"not null;size:200" json:"description"`
	Users       []User    `gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE" json:"users"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created-at"`
}
