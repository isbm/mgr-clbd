package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PingHandler struct {
	urn     string
	methods []string
}

func NewPingHandler() *PingHandler {
	ph := new(PingHandler)
	ph.urn = "ping"
	ph.methods = []string{POST}

	return ph
}

// URN returns uniform resource name of the handler to be installed at.
func (ph *PingHandler) URN() string {
	return ph.urn
}

// Methods returns available methods that this handler supposed to handle
func (ph *PingHandler) Methods() []string {
	return ph.methods
}

// SetDbx implements interface method to set a Dbx instance. Unused in this case.
func (ph *PingHandler) SetDbx(dbx *Dbx) {

}

// Handle implements the entry point of the handler
func (ph *PingHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"time": time.Now().Format(time.RFC1123Z),
	})
}
