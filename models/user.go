package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Структура пользователя
type User struct {
     ID       uuid.UUID `gorm:"primaryKey" json:"id"`
    Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
    Phone    string `json:"phone" validate:"required"`
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Password string `json:"password" validate:"required,min=8"`
    City     string `json:"city" validate:"required"`
    Address  string `json:"address" validate:"required"`
    UserName string `json:"userName" gorm:"unique;not null" validate:"required,min=3,max=100"`
}

// Обработчик валидации
var validate = validator.New()

// Метод для проверки модели пользователя на валидацию
func (user *User) Validate() error {
    return validate.Struct(user)
}


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
