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

var db *sql.DB

func getFriendsByBirthday(birthDate string) {
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
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func getAllFriends() {
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
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func setFriend(f friend.Friend) {
	stmt, err := db.Prepare("INSERT INTO friends(email, first_name, last_name, birth_date) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(f.Email, f.First_name, f.Last_name, f.Birth)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	getAllFriends()
	getFriendsByBirthday("1993-10-24")

	l := "2006-01-02"
	s := "1992-03-21"
	d, err := time.Parse(l, s)
	if err != nil {
		log.Fatal(err)
	}
	f := friend.Friend{Email: "f.t@wanaddo.fr", First_name: "Fabien", Last_name: "Tores", Birth: d}
	setFriend(f)
}
