package clbd

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	// Handle implements the entry point of the handler
	Handle(ctx *gin.Context)

	// URN returns uniform resource name of the handler to be installed at.
	URN() string

	// Methods returns available methods that this handler supposed to handle
	Methods() []string
}
