package adapter

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	friend "github.com/haagor/gobox/katas/birthdayGreetings/identity/entity"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "goadapter"
	Password = "goadapter"
	Dbname   = "identity"
)

var Db *sql.DB

func GetFriendsByBirthDate(birthDate time.Time) []friend.Friend {
	stmt, err := Db.Prepare(
		"SELECT email, first_name, last_name, birth_date FROM friends WHERE birth_date = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(birthDate)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res []friend.Friend
	var f friend.Friend
	for rows.Next() {
		rows.Scan(&f.Email, &f.FirstName, &f.LastName, &f.Birth)
		res = append(res, f)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func getAllFriends() {
	stmt, err := Db.Prepare(
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
		rows.Scan(&friend.Email, &friend.FirstName, &friend.LastName, &friend.Birth)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func setFriend(f friend.Friend) {
	stmt, err := Db.Prepare("INSERT INTO friends(email, first_name, last_name, birth_date) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(f.Email, f.FirstName, f.LastName, f.Birth)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
