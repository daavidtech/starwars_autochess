package game

import (
	"github.com/daavidtech/starwars_autochess/match"
	"github.com/google/uuid"
)

type User struct {
	id              string
	currentMatch    *match.Match
	currentPlayerID string
}

func NewUser() *User {
	return &User{
		id: uuid.New().String(),
	}
}

func (user *User) GetCurrentMatch() *match.Match {
	return user.currentMatch
}

func (user *User) SetCurrentMatch(match *match.Match) {
	user.currentMatch = match
}

func (user *User) GetCurrentPlayerID() string {
	return user.currentPlayerID
}

func (user *User) SetCurrentPlayerID(playerID string) {
	user.currentPlayerID = playerID
}
