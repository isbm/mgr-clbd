package clbd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strings"
)

const (
	ANY  = "any"
	POST = "post"
	GET  = "get"
)

type APIEndPoint struct {
	server *gin.Engine
	root   string
	port   int
	dbx    *Dbx
	mw     *Middleware
}

func NewAPIEndPoint(root string, dbx *Dbx) *APIEndPoint {
	api := new(APIEndPoint)
	api.root = "/" + strings.Trim(root, "/")
	api.server = gin.Default()
	api.port = 8080
	api.dbx = dbx

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
func (api *APIEndPoint) AddHandler(handler Handler) *APIEndPoint {
	handler.SetDbx(api.dbx)
	for _, hmeta := range handler.Handlers() {
		urn := api.getFullURN(hmeta.Route)
		for _, method := range hmeta.Methods {
			switch method {
			case GET:
				api.server.GET(urn, hmeta.Handle)
			case POST:
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
