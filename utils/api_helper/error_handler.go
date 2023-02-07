package api_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 錯誤處理
func HandleError(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	g.Abort()
	return
}
