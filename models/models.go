package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type Model interface {
	GetID() uint
	GetIDFieldName() string
	GetOwnerIDFieldName() string
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

func (b BaseModel) GetOwnerIDFieldName() string {
	return ""
}

type ValidationErrors []string

func (v ValidationErrors) Error() string {
	verrs, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("marshalling of validation errors failed: %v", err))
	}

	return string(verrs)
}
