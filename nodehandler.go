package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NodeHandler struct {
	dbx *Dbx
}

func NewNodeHandler() *NodeHandler {
	return new(NodeHandler)
}

// SetDbx sets the Dbx instance pointer
func (ph *NodeHandler) SetDbx(dbx *Dbx) {
	ph.dbx = dbx
}

// Handlers returns a map of supported handlers and their configuration
func (nh *NodeHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   "nodes",
			Handle:  nh.OnNodes,
			Methods: []string{POST, GET},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *NodeHandler) OnNodes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
