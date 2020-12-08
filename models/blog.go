package models

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jinzhu/gorm"
)

type Blog struct {
	BaseModel

	// The name of this blog.
	Name    string   `json:"name"`
	Domain  string   `json:"domain"`
	Posts   []Post   `json:"posts,omitempty"`
	Authors []Author `json:"authors,omitempty" gorm:"many2many:blog_authors"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) error {
	return validation.ValidateStruct(b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Domain, validation.Required, is.URL),
	)
}