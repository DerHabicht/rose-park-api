package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}

type BaseModelWithSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `json:"-"`
}

type BaseModelWithUUID struct {
	BaseModelWithSoftDelete
	PublicID uuid.UUID `json:"id"`
}
