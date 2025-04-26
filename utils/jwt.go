package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func CreateTokens(userID uuid.UUID) (string, string, error) {
    // Создание токена доступа
    accessToken := jwt.New(jwt.SigningMethodHS256)
    accessToken.Claims = jwt.MapClaims{
        "user_id": userID.String(), // Преобразуем UUID в строку
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Срок действия 24 часа
    }
    accessTokenString, err := accessToken.SignedString([]byte("secret_key"))
    if err != nil {
        return "", "", err
    }

    // Создание токена обновления
    refreshToken := jwt.New(jwt.SigningMethodHS256)
    refreshToken.Claims = jwt.MapClaims{
        "user_id": userID.String(), // Преобразуем UUID в строку
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Срок действия 72 часа
    }
    refreshTokenString, err := refreshToken.SignedString([]byte("secret_key"))
    if err != nil {
        return "", "", err
    }

    return accessTokenString, refreshTokenString, nil
}
