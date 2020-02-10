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
	ID          int64
	Name        string `gorm:"type:varchar(100);index;unique"`
	Description string
}

// Cluster node object
type ClusterNode struct {
	gorm.Model
	Fqdn          string
	MachineId     string // systemd
	ClientSystems int64
	LoadAverage   float64 // 5 min
	Status        uint
	Zone          ClusterZone `gorm:"foreignkey:ZoneId"`
	ZoneId        uint
}

// Client system object
type ClientSystem struct {
	ID        int64
	PubkeyFp  string
	Fqdn      string
	MachineId string
	Status    uint
}
