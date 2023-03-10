package auth

type Repository interface {
	GetUserByEmail(email string) (userId, password string, err error)
}
