package models

type User struct {
	Base
	Name     string `json:"name" gorm:"size:100"`
	Password string `json:"password" gorm:"size:100"`
	Email    string `json:"email" gorm:"size:100"`
}
type Company struct {
	Base
	Name     string `json:"name" gorm:"size:100"`
	Password string `json:"password" gorm:"size:100"`
	Email    string `json:"email" gorm:"size:100"`
}
