package users

type Repository interface {
	Insert(User) error
}
