package models

type Blog struct {
	BaseModel
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Posts   []Post   `json:"posts"`
	Authors []Author `json:"authors" gorm:"many2many:blog_authors"`
}
