package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Model interface {
	GetID() uint
	GetIDFieldName() string
}

type BaseModel struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type BaseModelWithUUID struct {
	BaseModel
	PublicID uuid.UUID `json:"id"`
}

func (b BaseModel) GetID() uint {
	return b.ID
}

func (b BaseModel) GetIDFieldName() string {
	return "ID"
}
