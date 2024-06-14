package user

type User struct {
	Id   int
	Name string
}

type UserService interface {
	GetUser(id int) (User, error)
}
