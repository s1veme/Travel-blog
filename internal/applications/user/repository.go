package user

type UserRepository interface {
	Create(*User) error
	FindByEmail(string) (*User, error)
}
