package main

import (
	"github.com/isbm/go-nanoconf"
	"github.com/isbm/mgr-clbd"
)

func main() {
	cfg := nanoconf.NewConfig("./mgr-clbd.conf")
	ep := clbd.NewAPIEndPoint("/api/v1").
		SetPort(cfg.Find("api").Int("port", "")).
		AddHandler(clbd.NewPingHandler())
	ep.Start()
}
