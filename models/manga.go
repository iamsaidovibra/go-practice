package models

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Author User   `json:"author" gorm:"foreignKey:AuthorID"`
	//File string `json:"file" gorm:"not null"`//TODO: add manga itself (file)
	AuthorID uint `json:"author_id" gorm:"not null"`
	//Comments []Comment `json:comments gorm:"foreignKey:ArticleID"`
	//Like     []Like    `json:comments gorm:"foreignKey:ArticleID"` //TODO: Read about foreing keys/tables relations and eventually add them here
}
