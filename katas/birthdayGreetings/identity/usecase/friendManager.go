package usecase

import (
	"log"
	"time"

	friend "github.com/haagor/gobox/katas/birthdayGreetings/identity/entity"
)

type DBAdapter interface {
	GetFriendsByBirthDate(birthDate time.Time) []friend.Friend
	GetAllFriends()
	SetFriend(f friend.Friend)
}

func GetFriendsBornAt(adapter DBAdapter, birth string) [][3]string {
	var res [][3]string

	l := "2006-01-02"
	b, err := time.Parse(l, birth)
	if err != nil {
		log.Fatal(err)
	}

	f := adapter.GetFriendsByBirthDate(b)
	for _, v := range f {
		res = append(res, [3]string{v.Email, v.FirstName, v.LastName})
	}

	return res
}
