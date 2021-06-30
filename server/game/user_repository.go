package game

type UserRepository interface {
	Fetch(userID string) *User
}

type UserRepositoryMemImpl struct {
	users map[string]*User
}

func NewUserRepository() *UserRepositoryMemImpl {
	return &UserRepositoryMemImpl{
		users: make(map[string]*User),
	}
}

func (userRepo *UserRepositoryMemImpl) Fetch(userID string) *User {
	return NewUser()
}
