package clbd

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/dbx"
	_ "github.com/isbm/mgr-clbd/docs"
	"github.com/isbm/mgr-clbd/handlers"
	"github.com/isbm/mgr-clbd/utils"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type APIEndPoint struct {
	server   *gin.Engine
	root     string
	port     int
	db       *dbx.Dbx
	mw       *Middleware
	handlers []hdl.Handler
	logger   *logrus.Logger
}

func NewAPIEndPoint(root string, db *dbx.Dbx) *APIEndPoint {
	api := new(APIEndPoint)
	api.logger = utils.GetTextLogger(logrus.DebugLevel, nil)
	api.root = "/" + strings.Trim(root, "/")
	api.port = 8080
	api.db = db
	api.handlers = make([]hdl.Handler, 0)

	// Setup server
	gin.SetMode(gin.ReleaseMode)
	api.server = gin.New()
	api.server.Use(gin.Recovery())
	api.server.Use(api.GinLogger())

	// Setup middleware
	//api.mw = NewMiddleware(api.root)
	//for _, method := range api.mw.Methods {
	//	api.server.Use(method)
	//}

	url := ginSwagger.URL("http://localhost:9080/swagger/doc.json")
	api.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return api
}

// GinLogger makes a custom logger via logrus
func (api *APIEndPoint) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		timestamp := time.Now()
		latency := timestamp.Sub(start)

		errmsg := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if errmsg != "" {
			errmsg = " " + errmsg
		}

		bodySize := c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		// Color methods
		var method string
		if api.logger.Out == os.Stderr {
			switch c.Request.Method {
			case "POST":
				method = aurora.Bold(aurora.BrightGreen("POST")).String()
			case "GET":
				method = aurora.Bold(aurora.BrightYellow("GET")).String()
			case "DELETE":
				method = aurora.Bold(aurora.BrightRed("GET")).String()
			default:
				method = aurora.Bold(aurora.BrightMagenta(c.Request.Method)).String()
			}
		} else {
			method = c.Request.Method
		}

		msg := fmt.Sprintf("%s (%s) %s - %s%s (latency: %s; size: %d)",
			method, c.ClientIP(), c.Request.Host, path, errmsg, latency, bodySize)
		if c.Writer.Status() != http.StatusOK || errmsg != "" {
			if errmsg != "" {
				api.logger.Errorln(msg)
			} else {
				api.logger.Warningln(msg)
			}
		} else {
			api.logger.Infoln(msg)
		}
	}
}

// SetStaticDirectoryRoot sets static directory to be served.
// Usually it is a fileserver for everything: binaries, states etc.
// It is available as "/pub".
func (api *APIEndPoint) SetStaticDirectoryRoot(root string) *APIEndPoint {
	if root != "" {
		api.server.Use(static.Serve("/pub", static.LocalFile(root, true)))
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
	api.handlers = append(api.handlers, handler)
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
		api.logger.Debugln("Initialise API endpoint at", urn)
	}
	return api
}

func (api *APIEndPoint) BootBackends() {
	for _, handler := range api.handlers {
		backend := handler.Backend()
		if backend != nil {
			backend.StartUp()
		}
	}
}

func (api *APIEndPoint) Start() {
	err := api.db.Open()
	if err != nil {
		panic(err)
	}
	defer api.db.Close()
	api.BootBackends()
	err = api.server.Run(fmt.Sprintf(":%d", api.port))
	if err != nil {
		panic(err)
	}
}
