package main

import (
	"github.com/isbm/mgr-clbd"
)

func main() {
	h := clbd.NewPingHandler()

	ep := clbd.NewAPIEndPoint()
	ep.AddHandler("/ping", "post", h)
	ep.Start()
}
