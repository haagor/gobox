package entrypoint

import (
	friendManager "identity/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Date string `json:"date"`
}

func GetFriendsBornAt(c *gin.Context) {
	d := Body{}

	if err := c.BindJSON(&d); err != nil {
		return
	}

	f := friendManager.GetFriendsBornAt(d.Date)
	c.IndentedJSON(http.StatusOK, f)
}
