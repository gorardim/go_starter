package component

import (
	"fmt"
	"testing"
	"time"

	"app/api/admin"
	"app/pkg/jwtutil"
	"github.com/stretchr/testify/assert"
)

func TestUserAuthComponent_GenerateToken(t *testing.T) {
	sign, err := jwtutil.Sign(secret, time.Hour*24*365, &admin.AuthUser{
		UserId: "10001",
	})
	assert.NoError(t, err)
	fmt.Println(sign)
}
