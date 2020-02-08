package main

import (
	"github.com/isbm/go-nanoconf"
	"github.com/isbm/mgr-clbd"
	"github.com/isbm/mgr-clbd/dbx"
)

func main() {
	cfg := nanoconf.NewConfig("./mgr-clbd.conf")

	db := dbx.NewDbxConnection().
		SetUser(cfg.Find("db").String("user", "")).
		SetPassword(cfg.Find("db").String("password", "")).
		SetDBName(cfg.Find("db").String("name", "")).
		SetDBHost(cfg.Find("db").String("fqdn", "")).
		SetDBPort(cfg.Find("db").Int("port", ""))

	ep := clbd.NewAPIEndPoint("/api/v1", db).
		SetPort(cfg.Find("api").Int("port", "")).
		AddHandler(clbd.NewPingHandler()).
		AddHandler(clbd.NewNodeHandler()).
		AddHandler(clbd.NewSystemsHandler())
	ep.Start()
}
