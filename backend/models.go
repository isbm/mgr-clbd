// Backend DB models

package backend

import (
	"github.com/jinzhu/gorm"
)

const (
	CL_ST_NEW = iota
	CL_ST_ORPHAN
	CL_ST_ACCEPTED
	CL_ST_DELETED
)

const (
	ND_ST_NEW = iota
	ND_ST_STAGED
	ND_ST_ACCEPTED
	ND_ST_OFFLINE // Scheduled offline
)

type ClusterZone struct {
	gorm.Model
	ID           int64
	Name         string
	Description  string
	ClusterNodes []ClusterNode `gorm:"foreignkey:ZoneID"`
}

// Cluster node object
type ClusterNode struct {
	ID            int64
	Fqdn          string
	MachineId     string // systemd
	ClientSystems int64
	LoadAverage   float64 // 5 min
	Status        uint
	ZoneID        uint
}

// Client system object
type ClientSystem struct {
	ID        int64
	PubkeyFp  string
	Fqdn      string
	MachineId string
	Status    uint
}
