package models

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jinzhu/gorm"
)

type Blog struct {
	BaseModel

	// Name is the title of the blog.
	Name    string   `json:"name"`

	// Domain is the domain name where this blog lives on the internet.
	Domain  string   `json:"domain"`
}

func (b *Blog) BeforeSave(tx *gorm.DB) error {
	return validation.ValidateStruct(b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Domain, validation.Required, is.URL),
	)
}