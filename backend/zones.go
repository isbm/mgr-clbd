package backend

type Zones struct {
	BaseBackend
}

func NewZonesBackend() *Zones {
	return new(Zones)
}

func (z *Zones) StartUp() {
	z.VerifyTables(&ClusterZone{})
}
