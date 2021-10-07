package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(secret string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	return string(bytes)
}

func ValidateHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	//fmt.Println(Hash(secret), "banding", string(hash))
	//return Hash(secret) == hash
	return err == nil

}
