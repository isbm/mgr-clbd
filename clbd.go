package clbd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/dbx"
	"github.com/isbm/mgr-clbd/handlers"
	"log"
	"path"
	"strings"
)

type APIEndPoint struct {
	server *gin.Engine
	root   string
	port   int
	db     *dbx.Dbx
	mw     *Middleware
}

func NewAPIEndPoint(root string, db *dbx.Dbx) *APIEndPoint {
	api := new(APIEndPoint)
	api.root = "/" + strings.Trim(root, "/")
	api.server = gin.Default()
	api.port = 8080
	api.db = db

	// Setup middleware
	api.mw = NewMiddleware(api.root)
	for _, method := range api.mw.Methods {
		api.server.Use(method)
	}

	return api
}

func (api *APIEndPoint) SetPort(port int) *APIEndPoint {
	api.port = port
	return api
}

// getFullURN joins root and the given URN into a full path
func (api *APIEndPoint) getFullURN(urn string) string {
	return path.Join(api.root, strings.Trim(urn, "/"))
}

// Add handler to the server with all declared API endpoints
func (api *APIEndPoint) AddHandler(handler hdl.Handler) *APIEndPoint {
	handler.SetDbx(api.db)
	for _, hmeta := range handler.Handlers() {
		urn := api.getFullURN(hmeta.Route)
		for _, method := range hmeta.Methods {
			switch method {
			case hdl.GET:
				api.server.GET(urn, hmeta.Handle)
			case hdl.POST:
				api.server.POST(urn, hmeta.Handle)
			default:
				api.server.Any(urn, hmeta.Handle)
			}
		}
		log.Println("Added handler at", urn)
	}
	return api
}

func (api *APIEndPoint) Start() {
	err := api.server.Run(fmt.Sprintf(":%d", api.port))
	if err != nil {
		panic(err)
	}
}
