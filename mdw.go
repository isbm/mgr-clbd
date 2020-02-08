// Global middleware implementaion

package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
	Methods []gin.HandlerFunc
	root    string
}

func NewMiddleware(root string) *Middleware {
	mw := new(Middleware)
	mw.root = root
	mw.Methods = []gin.HandlerFunc{
		mw.MW_CheckToken,
	}
	return mw
}

func (m *Middleware) _getMethodName(fullpath string) string {
	return fullpath[len(m.root):]
}

// MW_CheckToken provides an example of token verification
func (m *Middleware) MW_CheckToken(ctx *gin.Context) {
	if m._getMethodName(ctx.FullPath()) == "/ping" {
		ctx.Next()
	} else {
		token := ctx.Request.FormValue("token")
		if token == "0" { // Dummy token "0" means "admin"
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised access"})
		}
	}
}
