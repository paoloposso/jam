package auth

type AuthService struct {
}

func (auth AuthService) Login(email string, password string) (AuthResponse, error) {
	return AuthResponse{}, nil
}
