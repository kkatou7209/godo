package password

type PasswordHasher interface {
	// Hash password.
	Hash(password string) (string, error)
	// Verify password.
	Verify(password string, hash string) bool
}