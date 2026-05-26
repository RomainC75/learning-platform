package user_mngmt_infra

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {
}

func NewBcrypt() *BCrypt {
	return &BCrypt{}
}

func (b *BCrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (b *BCrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}
