package models

type Author struct {
	BaseModelWithUUID
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
	Posts []Post `json:"posts" gorm:"many2many:author_posts;"`
}
