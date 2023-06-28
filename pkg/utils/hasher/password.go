package utilshasher

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	pepper = "109339Lam@" // secret pepper
	saltSize = 16 // Size of the salt in bytes
)

func HashPassword(password string) (string, error) {
	salt := generateSalt(saltSize)
	passwordWithSalt := appendSaltAndPassword(salt, password)
	hashedPassword := hashWithPepper(passwordWithSalt)
	return hashedPassword, nil
}

func ComparePassword(hash, password string) error {
	salt, storedHash := extractSaltAndHash(hash)
	passwordWithSalt := appendSaltAndPassword(salt, password)
	computedHash := hashWithPepper(passwordWithSalt)
	if storedHash != computedHash {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func generateSalt(size int) []byte {
	salt := make([]byte, size)
	_, _ = rand.Read(salt)
	return salt
}

func appendSaltAndPassword(salt []byte, password string) string {
	saltedPassword := append(salt, []byte(password)...)
	return hex.EncodeToString(saltedPassword)
}

func hashWithPepper(passwordWithSalt string) string {
	pepperedPassword := []byte(passwordWithSalt + pepper)
	hash := sha256.Sum256(pepperedPassword)
	return hex.EncodeToString(hash[:])
}

func extractSaltAndHash(hashedPassword string) ([]byte, string) {
	saltedPassword, _ := hex.DecodeString(hashedPassword)
	salt := saltedPassword[:saltSize]
	storedHash := saltedPassword[saltSize:]
	return salt, hex.EncodeToString(storedHash)
}
