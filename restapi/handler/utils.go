package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

func getSignedInUserId(c *gin.Context) model.UserID {
	return c.MustGet(UserIdKey).(model.UserID)
}

func getUserIdFromPath(c *gin.Context) model.UserID {
	return model.UserID(c.Param("user_id"))
}
