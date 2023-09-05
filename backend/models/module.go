package models

import "time"

type Module struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	Enabled          bool      `gorm:"not null" json:"enabled"`
	Version          string    `gorm:"size:20;not null" json:"version"`
	Name             string    `gorm:"size:40;not null" json:"name"`
	Description      string    `gorm:"size:300;not null" json:"description"`
	HostURL          string    `gorm:"size:300;not null" json:"host-url"`
	HostInternalURL  string    `gorm:"size:300;not null" json:"-"`
	PermissionTypes  string    `gorm:"size:300;not null" json:"permission-types"`
	PermissionScopes string    `gorm:"size:300;not null" json:"permission-scopes"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created-at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated-at"`
}
