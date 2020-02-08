package main

import (
	"github.com/isbm/go-nanoconf"
	"github.com/isbm/mgr-clbd"
)

func main() {
	cfg := nanoconf.NewConfig("./mgr-clbd.conf")

	dbx := clbd.NewDbxConnection().
		SetUser(cfg.Find("db").String("user", "")).
		SetPassword(cfg.Find("db").String("password", "")).
		SetDBName(cfg.Find("db").String("name", "")).
		SetDBHost(cfg.Find("db").String("fqdn", "")).
		SetDBPort(cfg.Find("db").Int("port", ""))

	ep := clbd.NewAPIEndPoint("/api/v1", dbx).
		SetPort(cfg.Find("api").Int("port", "")).
		AddHandler(clbd.NewPingHandler()).
		AddHandler(clbd.NewNodeHandler())
	ep.Start()
}
