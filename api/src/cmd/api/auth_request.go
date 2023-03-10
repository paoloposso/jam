package main

type auth_request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
