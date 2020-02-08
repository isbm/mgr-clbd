package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type NodeHandler struct {
	db       *dbx.Dbx
	bknNodes *backend.Nodes
}

func NewNodeHandler() *NodeHandler {
	nh := new(NodeHandler)
	nh.bknNodes = backend.NewNodes(nh.db)
	return nh
}

// SetDbx sets the Dbx instance pointer
func (ph *NodeHandler) SetDbx(db *dbx.Dbx) {
	ph.db = db
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
