package entrypoint

import (
	"net/http"

	usecase "github.com/haagor/gobox/katas/birthdayGreetings/identity/usecase"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Date string `json:"date"`
}

func GetFriendsBornAt(db usecase.DBAdapter) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		d := Body{}

		if err := c.BindJSON(&d); err != nil {
			return
		}

		f := usecase.GetFriendsBornAt(db, d.Date)
		c.IndentedJSON(http.StatusOK, f)
	}

	return gin.HandlerFunc(fn)
}
