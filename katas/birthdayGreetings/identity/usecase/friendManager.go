package usecase

import (
	postgresDB "identity/adapter"
	"log"
	"time"
)

func getBirthsForDate(birth string) [][3]string {
	var res [][3]string

	l := "2006-01-02"
	s := birth
	b, err := time.Parse(l, s)
	if err != nil {
		log.Fatal(err)
	}

	f := postgresDB.GetFriendsByBirthday(b)
	for _, v := range f {
		res = append(res, [3]string{v.Email, v.First_name, v.Last_name})
	}

	return res
}
