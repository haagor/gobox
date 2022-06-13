package adapter

import (
	"time"

	friend "github.com/haagor/gobox/katas/birthdayGreetings/identity/entity"
)

type DBAdapter interface {
	GetFriendsByBirthDate(birthDate time.Time) []friend.Friend
	getAllFriends()
	setFriend(f friend.Friend)
}
