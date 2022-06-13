package usecase

import (
	"log"
	"time"

	DBAdapter "github.com/haagor/gobox/katas/birthdayGreetings/identity/adapter"
)

func GetFriendsBornAt(birth string) [][3]string {
	var res [][3]string

	l := "2006-01-02"
	s := birth
	b, err := time.Parse(l, s)
	if err != nil {
		log.Fatal(err)
	}

	f := DBAdapter.GetFriendsByBirthDate(b)
	for _, v := range f {
		res = append(res, [3]string{v.Email, v.FirstName, v.LastName})
	}

	return res
}
