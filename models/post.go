package models

import "time"

type Post struct {
	BaseModelWithSoftDelete
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	Body        string    `json:"body"`
	Authors     []Author  `json:"authors" gorm:"many2many:author_posts"`
}
