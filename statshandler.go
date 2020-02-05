package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatsHandler struct {
}

func NewStatsHandler() *StatsHandler {
	sh := new(StatsHandler)
	return sh
}

func (sh *StatsHandler) Handle(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
