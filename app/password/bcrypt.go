package password

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {}

func NewBycryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (b *BcryptPasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (b *BcryptPasswordHasher) Verify(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}