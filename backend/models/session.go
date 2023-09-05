package models

import (
	"time"

	"github.com/google/uuid"
)

type DeviceType int

const (
	Web     DeviceType = 1
	Android DeviceType = 2
	IOS     DeviceType = 3
	Windows DeviceType = 4
	MacOS   DeviceType = 5
	Linux   DeviceType = 6
)

type Session struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"-"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null" json:"-"`
	DeviceID     uuid.UUID  `gorm:"type:uuid;not null;unique" json:"-"`
	DeviceType   DeviceType `gorm:"not null" json:"device-type"`
	DeviceDetail string     `gorm:"not null;size:30" json:"device-detail"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created-at"`
}
