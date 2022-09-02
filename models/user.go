package models

//* validator gorm:"not null;default:null" prevent empty string pass in
type User struct {
	Uid       int   `json:"id" gorm:"primary_key"`
	Name 	 string `json:"name" binding:"required" gorm:"not null;type:varchar(100);default:null"`
	Username string `json:"username" binding:"required" gorm:"unique;not null;type:varchar(100);default:null"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null;type:varchar(40);default:null"`
	Password string `json:"password" binding:"required" gorm:"not null;type:varchar(100);default:null"`
}