package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
	"time"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	ph := new(PingHandler)
	return ph
}

func (ph *PingHandler) Backend() backend.Backend {
	return nil
}

// SetDbx implements interface method to set a Dbx instance. Unused in this case.
func (ph *PingHandler) SetDbx(db *dbx.Dbx) {}

// Handlers returns a map of supported handlers and their configuration
func (ph *PingHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   "ping",
			Handle:  ph.OnPing,
			Methods: []string{ANY},
		},
	}
}

// Handle implements the entry point of the handler
func (ph *PingHandler) OnPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"time": time.Now().Format(time.RFC1123Z),
	})
}
