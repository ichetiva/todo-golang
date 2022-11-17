package security

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func Hash(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func Match(password, hash string) bool {
	return Hash(password) == hash
}

func MakeToken() string {
	randString := getRandomString()
	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)
	return hex.EncodeToString(bs)
}

func getRandomString() string {
	rand.Seed(time.Now().Unix())

	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 50)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
