package users

type Repository interface {
	Insert(user User) (id string, err error)
	Update(user User) error
	GetByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}
