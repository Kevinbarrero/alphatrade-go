package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinbarrero/alphatrade-go/db"
)

type GetCoinRequest struct {
	symbol    string `json:"symbol"`
	Kline     string `json:"kline"`
	StartTime uint64 `json:"start_time"`
	EndTime   uint64 `json:"end_time"`
}

func (server *Server) getCoin(ctx *gin.Context) {
	var req GetCoinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// doesn't have to be here
	db.UpdateCoins()
	coin := server.store.Get

}
