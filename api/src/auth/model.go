package auth

type Login struct {
	Email    string
	Password string
}

type AuthResponse struct {
	Token  string
	UserID string
}
