package clbd

import (
	"github.com/gin-gonic/gin"
)

const (
	ANY  = "any"
	POST = "post"
	GET  = "get"
)

type APIEndPoint struct {
	server *gin.Engine
}

func NewAPIEndPoint() *APIEndPoint {
	api := new(APIEndPoint)
	api.server = gin.Default()

	return api
}

func (api *APIEndPoint) AddHandler(uri string, method string, handler Handler) {
	switch method {
	case GET:
		api.server.GET(uri, handler.Handle)
	case POST:
		api.server.POST(uri, handler.Handle)
	case ANY:
		api.server.Any(uri, handler.Handle)
	}
}

func (api *APIEndPoint) Start() {
	err := api.server.Run(":9090")
	if err != nil {
		panic(err)
	}
}
