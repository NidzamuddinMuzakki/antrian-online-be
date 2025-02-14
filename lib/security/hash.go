package security

import (
	"antrian-golang/config"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt(length int) (string, error) {
	rand.NewSource(time.Now().UnixNano())
	salt := make([]byte, length)
	for i := range salt {
		salt[i] = byte(rand.Intn(128))
	}
	return hex.EncodeToString(salt), nil
}

func HashPasswordV2(password string, salt string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password + salt))
	if err != nil {
		return "", err
	}
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return salt + ":" + hashedPassword, nil
}

func HashPassword(password string, salt string) (string, error) {
	h := sha256.New()
	h.Write([]byte(password + salt + config.Cold.GeneratePassword))

	hashed := h.Sum(nil)
	hash, err := bcrypt.GenerateFromPassword(hashed, 10)
	if err != nil {
		return "", err
	}

	return salt + ":" + string(hash), nil
}

func ComparePassword(hash string, password string, salt string) error {
	h := sha256.New()
	h.Write([]byte(password + salt + config.Cold.GeneratePassword))

	hashed := h.Sum(nil)

	err := bcrypt.CompareHashAndPassword([]byte(hash), hashed)
	if err != nil {
		return err
	}

	return nil
}
