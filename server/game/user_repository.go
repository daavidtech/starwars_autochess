package game

type UserRepository interface {
	Fetch(userID string) *User
	FetchByUsername(username string) *User
	Save(user *User)
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
	return userRepo.users[userID]
}

func (userRepo *UserRepositoryMemImpl) FetchByUsername(username string) *User {
	for _, user := range userRepo.users {
		if user.GetUsername() == username {
			return user
		}
	}

	return nil
}

func (userRepo *UserRepositoryMemImpl) Save(user *User) {
	userRepo.users[user.GetID()] = user
}
