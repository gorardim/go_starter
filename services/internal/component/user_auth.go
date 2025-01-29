package component

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"app/pkg/gormx/g"

	"app/api/admin"
	"app/pkg/jwtutil"
	"app/pkg/randx"
	"app/services/internal/repo"

	"github.com/gin-gonic/gin"
)

type UserAuthComponent struct {
	CenterUserRepo *repo.CenterUserRepo
}

const secret = "PX7Pq&58kUHYQUbxn1F!5*2Sb&7#i3jU"

func (u *UserAuthComponent) HashPassword(password string, userId string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write([]byte(userId))
	// write salt
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(nil))
}

func (u *UserAuthComponent) GenerateToken(userId string) (string, error) {
	sign, err := jwtutil.Sign(secret, time.Hour*24*365, &admin.AuthUser{
		UserId: userId,
	})
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (u *UserAuthComponent) GenerateUserId() string {
	format := time.Now().Format("0601021504")
	return format + randx.Digit(4)
}

var publicAuthPaths = map[string]bool{
	"/admin/auth/login":   true,
	"/admin/auth/captcha": true,
}

func (u *UserAuthComponent) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if publicAuthPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		token := c.GetHeader("Authorization")
		// jwt
		claim, err := jwtutil.Parse[admin.AuthUser](secret, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"code":    admin.ErrUnauthenticated,
			})
			c.Abort()
			return
		}
		ctx := jwtutil.NewContext(c.Request.Context(), claim.Custom)
		// bind 绑定用户信息
		user, err := u.CenterUserRepo.Get(c.Request.Context(), g.Where("user_id = ?", claim.Custom.UserId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"code":    admin.ErrInvalidParam,
			})
			c.Abort()
			return
		}
		ctx = NewUserContext(ctx, &User{
			Id:       user.Id,
			UserId:   user.UserId,
			Username: user.Username,
		})
		c.Request = c.Request.WithContext(ctx)
	}
}
