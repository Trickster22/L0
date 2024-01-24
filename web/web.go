package web

import (
	"net/http"

	"github.com/Trickster22/L0/cache"
	"github.com/gin-gonic/gin"
)

func GetOrderByOrderUid(cache *cache.Cache) func(c *gin.Context) {
	f := func(c *gin.Context) {
		uid := c.Param("uid")
		cacheOrder, status := cache.Get(uid)
		if status {
			c.JSON(http.StatusOK, cacheOrder)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Нет записей",
			})
		}
	}
	return gin.HandlerFunc(f)
}
