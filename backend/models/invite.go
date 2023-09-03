package models

import (
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"invite-id"`
	SenderID      uuid.UUID  `gorm:"not null;type:uuid" json:"sender-id"`
	RecieverEmail string     `gorm:"not null;size:320" json:"reciever-email"`
	TargetType    ObjectType `gorm:"not null;size:30" json:"target-type"`
	TargetID      uuid.UUID  `gorm:"not null;type:uuid" json:"target"`
	ExpireAt      time.Time  `gorm:"not null" json:"expire-at"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created-at"`
}
