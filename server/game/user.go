package game

import (
	"github.com/google/uuid"
)

type User struct {
	id              string
	username        string
	password        string
	currentMatchID  string
	currentPlayerID string
}

func NewUser() *User {
	return &User{
		id: uuid.New().String(),
	}
}

func (user *User) GetID() string {
	return user.id
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) SetUsername(username string) {
	user.username = username
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

func (user *User) GetCurrentMatchID() string {
	return user.currentMatchID
}

func (user *User) SetCurrentMatchID(matchID string) {
	user.currentMatchID = matchID
}

func (user *User) GetCurrentPlayerID() string {
	return user.currentPlayerID
}

func (user *User) SetCurrentPlayerID(playerID string) {
	user.currentPlayerID = playerID
}
