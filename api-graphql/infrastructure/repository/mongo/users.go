package mongo

import "api-graphql/users"

type Users struct {
}

func (this Users) Insert(user users.User) (id string, err error) {
	println("user inserted")
	return "aaa", nil
}
