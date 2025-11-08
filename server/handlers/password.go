package handlers

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"

func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	return string(password)
}

func GetPassword(c *gin.Context) {
	password := generatePassword(16)
	c.JSON(http.StatusOK, gin.H{
		"password": password,
	})
}

func PostPassword(c *gin.Context) {
	var body struct {
		Size int `json:"size"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if body.Size <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size must be greater than 0"})
		return
	}

	password := generatePassword(body.Size)
	c.JSON(http.StatusOK, gin.H{
		"password": password,
	})
}
