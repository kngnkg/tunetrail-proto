package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
)

func getSignedInUserId(c *gin.Context) model.UserID {
	return c.MustGet(UserIdKey).(model.UserID)
}

func getUserIdFromPath(c *gin.Context) model.UserID {
	return model.UserID(c.Param("user_id"))
}

func getFolloweeIdFromPath(c *gin.Context) model.UserID {
	return model.UserID(c.Param("followee_user_id"))
}

func getPostIdFromPath(c *gin.Context) string {
	return c.Param("post_id")
}

func getPaginationFromQuery(c *gin.Context) (*model.Pagination, error) {
	nc := c.DefaultQuery("next_cursor", "")
	pc := c.DefaultQuery("previous_cursor", "")
	lstr := c.DefaultQuery("limit", "10")

	l, err := strconv.Atoi(lstr)
	if err != nil {
		return nil, err
	}

	return &model.Pagination{
		NextCursor:     nc,
		PreviousCursor: pc,
		Limit:          l,
	}, nil
}
