package backend

import (
	"fmt"
)

type Zones struct {
	BaseBackend
}

func NewZonesBackend() *Zones {
	return new(Zones)
}

func (z *Zones) StartUp() {
	z.VerifyTables(&ClusterZone{})
}

// CreateZone creates a new zone
func (z *Zones) CreateZone(name string, descr string) error {
	var err error
	zone := &ClusterZone{
		Name: name, Description: descr,
	}
	if z.db.DB().NewRecord(zone) {
		if err = z.db.DB().Create(&zone).Error; err != nil {
			logger.Errorln(err.Error())
		} else {
			logger.Infof("Added a new Zone %s", name)
		}
	} else {
		err = fmt.Errorf("Unable to add a new Zone '%s'", name)
		logger.Errorln(err.Error())
	}

	return err
}

// RemoveZone removes an empty zone. If zone still contains nodes,
// it won't be removed, but error will be issued.
func (z *Zones) RemoveZone(name string) error {
	var err error

	return err
}

// NodesInZone returns an amount of attached cluster nodes to the zone
func (z *Zones) NodesInZone(name string) int {

	return 0
}
