package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NodeHandler struct {
	urn     string
	methods []string
	dbx     *Dbx
}

func NewNodeHandler() *NodeHandler {
	ph := new(NodeHandler)
	ph.urn = "nodes"
	ph.methods = []string{POST, GET}

	return ph
}

// URN returns uniform resource name of the handler to be installed at.
func (ph *NodeHandler) URN() string {
	return ph.urn
}

// Methods returns available methods that this handler supposed to handle
func (ph *NodeHandler) Methods() []string {
	return ph.methods
}

func (ph *NodeHandler) SetDbx(dbx *Dbx) {
	ph.dbx = dbx
}

// Handle implements the entry point of the handler
func (ph *NodeHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
