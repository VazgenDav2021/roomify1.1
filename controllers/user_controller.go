package controllers

import (
	"errors"
	"net/http"
	"roomify-backend/config"
	"roomify-backend/models"
	"roomify-backend/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetAccount - Получение информации о пользователе
func GetAccount(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.GetErrorMessage(404))
	}

	// Безопасный ответ без пароля
	response := struct {
		ID       uuid.UUID   `json:"id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		City     string `json:"city"`
	}{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.UserName,
		Phone:    user.Phone,
		Address:  user.Address,
		City:     user.City,
	}

	return c.JSON(http.StatusOK, response)
}

// RegisterUser - Регистрация нового пользователя
func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetErrorMessage(400))
	}

	if err := utils.ValidateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetErrorMessage(400)+": "+err.Error())
	}

	// Проверки на уникальность
	if err := checkUserUniqueness(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}

	// Генерация токенов
	accessToken, refreshToken, err := utils.CreateTokens(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// LoginUser - Вход пользователя
func LoginUser(c echo.Context) error {
	var loginData models.User
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetErrorMessage(400))
	}

	if err := utils.ValidateUser(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetErrorMessage(400)+": "+err.Error())
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.GetErrorMessage(404))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.GetErrorMessage(2313))
	}

	accessToken, refreshToken, err := utils.CreateTokens(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// DeleteUser - Удаление пользователя
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.GetErrorMessage(404))
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Пользователь успешно удален",
	})
}

// UpdateUser - Обновление пользователя
func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.GetErrorMessage(404))
	}

	var updatedData models.User
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetErrorMessage(400))
	}

	// Обновляем поля
	if updatedData.Email != "" {
		user.Email = updatedData.Email
	}
	if updatedData.Phone != "" {
		user.Phone = updatedData.Phone
	}
	if updatedData.Name != "" {
		user.Name = updatedData.Name
	}
	if updatedData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
		}
		user.Password = string(hashedPassword)
	}
	if updatedData.City != "" {
		user.City = updatedData.City
	}
	if updatedData.Address != "" {
		user.Address = updatedData.Address
	}
	if updatedData.UserName != "" {
		user.UserName = updatedData.UserName
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetErrorMessage(500))
	}

	return c.JSON(http.StatusOK, user)
}

// checkUserUniqueness - проверка уникальности email, username и phone
func checkUserUniqueness(user *models.User) error {
	var existingUser models.User

	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New(utils.GetErrorMessage(4313))
	}

	if err := config.DB.Where("user_name = ?", user.UserName).First(&existingUser).Error; err == nil {
		return errors.New(utils.GetErrorMessage(4312))
	}

	if err := config.DB.Where("phone = ?", user.Phone).First(&existingUser).Error; err == nil {
		return errors.New(utils.GetErrorMessage(4311))
	}

	return nil
}

