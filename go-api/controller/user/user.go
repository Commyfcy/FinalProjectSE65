package user

import (
	"net/http"
	"se/jwt-api/orm"

	"github.com/gin-gonic/gin"
)

var hmacSampleSecret []byte

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Sucessful",
		"users": users})
}
