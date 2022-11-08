package core

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const XRequestIDKey = "X-Request-ID"

func generateMsgIDFromContext(c *gin.Context) string {
	rid := c.GetHeader(XRequestIDKey)
	if rid == "" {
		rid = uuid.Must(uuid.NewV4()).String()
		c.Request.Header.Set(XRequestIDKey, rid)
		c.Header(XRequestIDKey, rid)
	}
	return rid
}
