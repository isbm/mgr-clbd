package backend

import (
	"fmt"
	"github.com/isbm/mgr-clbd/dbx"
)

type Nodes struct {
	db *dbx.Dbx
}

func NewNodes(db *dbx.Dbx) *Nodes {
	n := new(Nodes)
	n.db = db
	return n
}

func (n *Nodes) ListNodes() {
	fmt.Println("List nodes")
}
