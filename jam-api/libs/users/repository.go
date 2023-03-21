package users

type Repository interface {
	Insert(User) error
	Get(id string) (*User, error)
}
