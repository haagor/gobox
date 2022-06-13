package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"

	postgresDB "github.com/haagor/gobox/katas/birthdayGreetings/identity/adapter"
	friend "github.com/haagor/gobox/katas/birthdayGreetings/identity/entrypoint"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresDB.Host, postgresDB.Port, postgresDB.User, postgresDB.Password, postgresDB.Dbname)

	var err error
	postgresDB.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer postgresDB.Db.Close()

	router := gin.Default()
	router.GET("/friends", friend.GetFriendsBornAt)

	router.Run("localhost:8080")
}
