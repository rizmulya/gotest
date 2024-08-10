package utils

import (
    "crypto/rand"
    "math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

// Generate random string 23 bytes
func RandStr() string {
    var randStr string
    for i := 0; i < 23; i++ {
        index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        randStr += string(charset[index.Int64()])
    }
    return randStr
}
