package database

import "api-graphql/users"

type UserRepository struct {
}

func (this UserRepository) Insert(user users.User) (id string, err error) {
	println("user inserted")
	return "aaa", nil
}

func (this UserRepository) Update(user users.User) error {
	println("user updated")
	return nil
}
