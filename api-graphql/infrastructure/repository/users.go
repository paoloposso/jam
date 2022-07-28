package repository

import "api-graphql/users"

type User struct {
}

func (this User) Insert(user users.User) (id string, err error) {
	println("user inserted")
	return "aaa", nil
}

func (this User) Update(user users.User) error {
	println("user updated")
	return nil
}
