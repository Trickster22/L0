package cache

import (
	"log"
	"sync"

	"github.com/Trickster22/L0/models"
)

type Cache struct {
	sync.RWMutex
	orders map[string]models.Order
}

func New() *Cache {
	log.Println("Initializing in memory cache")
	orders := make(map[string]models.Order)
	cache := Cache{
		orders: orders,
	}
	return &cache
}

func (c *Cache) Set(key string, order models.Order) {
	log.Println("Put order in cache, order =", order)
	c.Lock()
	defer c.Unlock()
	c.orders[key] = order
}

func (c *Cache) Get(key string) (models.Order, bool) {
	log.Println("Get order from cache")
	c.RLock()
	defer c.RUnlock()
	order, found := c.orders[key]
	if found {
		return order, true
	}
	return order, false
}
