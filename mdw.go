// Global middleware implementaion

package clbd

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (m *Middleware) MW_CheckToken(ctx *gin.Context) {
	fmt.Println("CHECK TOKEN>>", m._getMethodName(ctx.FullPath()))
	ctx.Next()
}
