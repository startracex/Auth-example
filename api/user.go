package api

import (
	"main/db"
	"main/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func UserApi(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.Status(400)
		c.Abort()
		return
	} else {
		hex, _ := db.TokenMap.Get(token)
		if hex == "" {
			c.Status(400)
			c.Abort()
			return
		}
		objid, _ := primitive.ObjectIDFromHex(hex)
		user := db.LoginDatabase.Find(bson.D{{Key: "_id", Value: objid}})
		utils.SafeInfo(&user)
		c.JSON(200, user)
	}
}
