package models

import "time"

type Post struct {
	BaseModel
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	Body        string    `json:"body"`
	Authors     []Author  `json:"authors" gorm:"many2many:author_posts"`
}
