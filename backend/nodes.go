package backend

import (
	"fmt"
	"github.com/isbm/mgr-clbd/dbx"
	"log"
)

type Nodes struct {
	db *dbx.Dbx
}

func NewNodesBackend() *Nodes {
	return new(Nodes)
}

// SetDbx sets the dbx reference
func (n *Nodes) SetDbx(d *dbx.Dbx) {
	n.db = d
}

// Initialise cluster schema
func (n *Nodes) StartUp() {
	// Stupid, simple
	if !n.db.DB().HasTable(&ClusterNode{}) {
		n.db.DB().CreateTable(&ClusterNode{})
		log.Println("Created ClusterNode table")
	}

	// Automigrate
	n.db.DB().AutoMigrate(&ClusterNode{})
}

func (n *Nodes) AddNode(node *ClusterNode) {
	log.Println("Add node", node.Fqdn)
	n.db.DB().Create(node)
}

func (n *Nodes) ListNodes() {
	fmt.Println("List nodes")
	clusterNodes := []ClusterNode{}
	n.db.DB().Find(&clusterNodes)

	fmt.Println("Found ", len(clusterNodes), "nodes in the cluster")
	for _, node := range clusterNodes {
		fmt.Print("Node:", node.Fqdn)
	}
}
