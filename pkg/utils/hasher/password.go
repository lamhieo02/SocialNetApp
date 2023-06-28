package utilshasher

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const (
	pepper = "109339Lam@" // secret pepper
	saltSize = 16 // Size of the salt in bytes
)

func HashPassword(password string) (string, string, error) {
	salt := generateSalt(16)
	// Add salt and pepper to the password
	saltedPassword := salt + password + pepper
	// Generate hash
	hashedPassword, err := generateHash(saltedPassword)
	if err != nil {
		return "", "",err
	}
	return hashedPassword, salt, nil
}

func ComparePassword(hashedPassword, password, salt string) error {
	// Generate the salted and hashed password for the entered password
	enteredSaltedPassword := salt + password + pepper

	// Compare the entered password with the stored hashed password
	err := comparePasswords(hashedPassword, enteredSaltedPassword)
	return err
}


func generateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateSalt(size int) string {
	salt := make([]byte, size)
	_, _ = rand.Read(salt)
	saltString := base64.URLEncoding.EncodeToString(salt)

	return saltString
}
