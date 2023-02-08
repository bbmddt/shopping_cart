package api_helper

import (
	"shopping_cart/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

// 從context獲得Userid
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
