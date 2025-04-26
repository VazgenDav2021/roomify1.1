package utils

import (
	"errors"
	"regexp"
	"roomify-backend/models"
)

// ValidateUser - Валидация данных пользователя при регистрации/логине
func ValidateUser(user *models.User) error {
	// Проверка обязательности email
	if user.Email == "" {
		return errors.New("email является обязательным полем")
	}
	if !isValidEmail(user.Email) {
		return errors.New("неверный формат email")
	}

	// Проверка обязательности пароля
	if user.Password == "" {
		return errors.New("пароль является обязательным полем")
	}
	if len(user.Password) < 6 {
		return errors.New("пароль должен быть не менее 6 символов")
	}
	if !isStrongPassword(user.Password) {
		return errors.New("пароль должен содержать хотя бы одну букву и одну цифру")
	}

	// Проверка обязательности имени пользователя
	if user.UserName == "" {
		return errors.New("имя пользователя является обязательным полем")
	}
	if len(user.UserName) < 3 {
		return errors.New("имя пользователя должно быть не менее 3 символов")
	}

	// Проверка обязательности телефона
	if user.Phone == "" {
		return errors.New("номер телефона является обязательным полем")
	}
	if !isValidPhone(user.Phone) {
		return errors.New("неверный формат номера телефона")
	}

	return nil
}

// Функция для проверки правильности формата email
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Функция для проверки правильности формата номера телефона
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(phone)
}

// Функция для проверки силы пароля
func isStrongPassword(password string) bool {
	// Должен содержать хотя бы одну букву и одну цифру
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	return hasLetter && hasNumber
}
