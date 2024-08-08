package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "github.com/joho/godotenv"
    "os"
    "time"
    "log"
)

var jwtSecret []byte

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⛔ ERROR", err)
    }
    jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
    UserID uint
    Role   string
    jwt.StandardClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}
