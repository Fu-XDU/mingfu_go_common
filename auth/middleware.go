package auth

import (
	"github.com/Fu-XDU/mingfu_go_common/base_response"
	"github.com/Fu-XDU/mingfu_go_common/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

const KeyUuid = "UUID"

func Auth(c *gin.Context) {
	authorization := c.GetHeader(constants.Authorization)
	uuid, ok := VerifyJwt(authorization)
	if !ok {
		c.JSON(http.StatusUnauthorized, base_response.NewErrorResponse(nil, base_response.Unauthorized))
		c.Abort()
		return
	}

	c.Set(KeyUuid, uuid)
	c.Next()
}
