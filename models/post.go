package models

import (
	iso639 "github.com/emvi/iso-639-1"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type Post struct {
	BaseModelWithSoftDelete
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	Language    string    `json:"language"`
	Body        string    `json:"body"`
	Authors     []Author  `json:"authors" gorm:"many2many:author_posts"`
}

func verifyISO639(code interface{}) error {
	c, ok := code.(string)
	if !ok {
		return errors.Errorf("%v is not a valid ISO-639-1 code", code)
	}

	if !iso639.ValidCode(c) {
		return errors.Errorf("%s is not a valid ISO-639-1 code", code)
	}

	return nil
}

func (p *Post) BeforeSave(tx *gorm.DB) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Language, validation.Required, validation.By(verifyISO639)),
	)
}
