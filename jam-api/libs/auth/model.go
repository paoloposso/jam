package auth

type AuthenticatedUser struct {
	Email  string
	Token  string
	UserID string
}
