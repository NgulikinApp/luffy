package bcrypt_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	hasher "github.com/NgulikinApp/luffy/hash/bcrypt"
	"github.com/stretchr/testify/assert"
)

const encryptCost = 10

func TestGenerateHash(t *testing.T) {
	password := "ngulikin"

	h := hasher.NewBcrypt()
	hashPassword := h.Generate(password)

	cost, err := bcrypt.Cost([]byte(hashPassword))

	assert.NoError(t, err)
	assert.Equal(t, encryptCost, cost)
	assert.NotEqual(t, password, hashPassword)
	assert.NotEmpty(t, hashPassword)
}

func TestVerifyHash(t *testing.T) {
	password := "ngulikin"

	h := hasher.NewBcrypt()
	hashPassword := h.Generate(password)

	isTrue := h.Verify(password, hashPassword)
	isFalse := h.Verify("hajar", hashPassword)

	assert.True(t, isTrue)
	assert.False(t, isFalse)
}
