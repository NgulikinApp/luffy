package bcrypt

import (
	"github.com/NgulikinApp/luffy/hash"
	"golang.org/x/crypto/bcrypt"
)

const EncryptCost = 10

type hasher struct{}

func (h *hasher) Generate(password string) string {

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)

	return string(hashedPass)
}

func (h *hasher) Verify(guess string, password string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(guess)); err != nil {
		return false
	}

	return true
}

func NewBcrypt() hash.Hash {
	return &hasher{}
}
