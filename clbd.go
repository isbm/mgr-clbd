package clbd

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
}

func NewAPIEndPoint(root string) *APIEndPoint {
	api := new(APIEndPoint)
	api.root = "/" + strings.Trim(root, "/")
	api.server = gin.Default()
	api.port = 8080

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

// Add handlers to the server
func (api *APIEndPoint) AddHandler(handler Handler) *APIEndPoint {
	urn := api.getFullURN(handler.URN())
	for _, method := range handler.Methods() {
		switch method {
		case GET:
			api.server.GET(urn, handler.Handle)
		case POST:
			api.server.POST(urn, handler.Handle)
		case ANY:
			api.server.Any(urn, handler.Handle)
		}
	}
	return api
}

func (api *APIEndPoint) Start() {
	err := api.server.Run(fmt.Sprintf(":%d", api.port))
	if err != nil {
		panic(err)
	}
}
