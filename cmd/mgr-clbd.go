package main

import (
	"fmt"
	"github.com/isbm/go-nanoconf"
	"github.com/isbm/mgr-clbd"
	"github.com/isbm/mgr-clbd/dbx"
	"github.com/isbm/mgr-clbd/handlers"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func run(ctx *cli.Context) error {
	cfg := nanoconf.NewConfig(ctx.String("config"))

	cmdPort := fmt.Sprintf("%d", ctx.Int("kvs-port"))
	if cmdPort == "0" || cmdPort == "4000" {
		cmdPort = ""
	}

	db := dbx.NewDbxConnection().
		SetUser(cfg.Find("kvs").String("user", "")).
		SetPassword(cfg.Find("kvs").String("password", "")).
		SetDBName(cfg.Find("kvs").String("name", "")).
		SetDBHost(cfg.Find("kvs").String("fqdn", "")).
		SetDBPort(cfg.Find("kvs").DefaultInt("port", cmdPort, 4000))

	cmdPort = fmt.Sprintf("%d", ctx.Int("api-port"))
	if cmdPort == "0" || cmdPort == "8080" {
		cmdPort = ""
	}

	nodehandler := hdl.NewNodeHandler("node")
	nodehandler.SetConfig(cfg)

	ep := clbd.NewAPIEndPoint("/api/v1", db).
		SetPort(cfg.Find("api").DefaultInt("port", cmdPort, 8080)).
		AddHandler(hdl.NewPingHandler("cluster")).
		AddHandler(nodehandler).
		AddHandler(hdl.NewSystemsHandler("systems")).
		AddHandler(hdl.NewZoneHandler("zones")).
		SetStaticDirectoryRoot(cfg.Find("general").String("static-root", ""))

	fmt.Println("For Swagger: http://localhost:9080/swagger/index.html")

	ep.Start()

	return nil
}

func main() {
	app := &cli.App{
		Version: "0.1 Alpha",
		Name:    "mgr-clbd",
		Usage:   "Uyuni Cluster Director Daemon",
		Action:  run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Path to configuration file",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "kvs-port",
				Value: 4000,
				Usage: "Specify KV store port (override configuration)",
			},
			&cli.IntFlag{
				Name:  "api-port",
				Value: 8080,
				Usage: "Specify API server port (override configuration)",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
