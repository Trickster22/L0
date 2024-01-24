package main

import (
	"github.com/Trickster22/L0/cache"
	"github.com/Trickster22/L0/database"
	"github.com/Trickster22/L0/nats"
	"github.com/Trickster22/L0/web"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func main() {
	db := database.Connect()
	defer db.Close()

	cache := cache.New()
	ordersFromDb := database.GetOrders(db)
	for _, order := range ordersFromDb {
		cache.Set(order.Order_uid, order)
	}
	sc, err := stan.Connect("test-cluster", "test", stan.NatsURL("nats://localhost:4222"))
	database.CheckError(err)
	defer sc.Close()
	sc.Subscribe("testChanel", nats.Sub(cache, db))

	router := gin.Default()
	router.GET("/order/:uid", web.GetOrderByOrderUid(cache))
	router.Run()
}
