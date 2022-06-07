package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	friend "identity/entity"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "goadapter"
	password = "goadapter"
	dbname   = "identity"
)

func getFriendsByBirthday(birthDate string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"SELECT email, first_name, last_name, birth_date FROM friends WHERE birth_date = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	l := "2006-01-02"
	s := birthDate
	d, err := time.Parse(l, s)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(d)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var friend friend.Friend
	for rows.Next() {
		rows.Scan(&friend.Email, &friend.First_name, &friend.Last_name, &friend.Birth)
		fmt.Println(friend.Email)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func getAllFriends() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"SELECT * FROM friends")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var friend friend.Friend
	for rows.Next() {
		rows.Scan(&friend.Email, &friend.First_name, &friend.Last_name, &friend.Birth)
		fmt.Println(friend.Email)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	getAllFriends()
	getFriendsByBirthday("1993-10-24")
}
