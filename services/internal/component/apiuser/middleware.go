package apiuser

import (
	"net/http"

	"app/api/api"
	"app/services/internal/repo"

	"github.com/gin-gonic/gin"
)

var BindUser = func(userRepo *repo.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authUser, ok := api.UserFormContext(c.Request.Context())
		if ok {
			user, err := userRepo.FindById(c.Request.Context(), authUser.UserId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    api.ErrInvalidParam,
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.Request = c.Request.WithContext(NewContext(c.Request.Context(), user))
		}
		c.Next()
	}
}
