package api

import (
	"log"
	msg "github.com/soxueren/greenplum-node/pkg/ws"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("init messagesservice")
}

func GetMessage(c *gin.Context) {
	msg.Wshandler(c.Writer, c.Request)
}
