package utils

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
    "errors"
    "context"
)

var jwtKey = []byte("your_secret_key")

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GenerateJWT(userID int64, permission string) (string, error) {
    claims := &jwt.StandardClaims{
        Subject:   string(rune(userID)),
        ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (int64, error) {
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }
    // Преобразуем Subject в int64
    // Здесь нужно конвертировать строку в int64
    return strconv.ParseInt(claims.Subject, 10, 64)
}

func GetUserIDFromContext(ctx context.Context) int64 {
    userID, ok := ctx.Value("userID").(int64)
    if !ok {
        return 0
    }
    return userID
}
