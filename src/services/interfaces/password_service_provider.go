package interfaces

type PasswordServiceProvider interface {
	EncodePassword(password string) string
}
