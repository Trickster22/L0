package nats

import (
	"database/sql"
	"encoding/json"
	"reflect"

	"github.com/Trickster22/L0/cache"
	"github.com/Trickster22/L0/database"
	"github.com/Trickster22/L0/models"
	"github.com/nats-io/stan.go"
)

func Sub(cache *cache.Cache, db *sql.DB) func(m *stan.Msg) {
	f := func(m *stan.Msg) {
		jsonStr := string(m.Data)
		var order models.Order
		err := json.Unmarshal([]byte(jsonStr), &order)
		if err == nil {
			if _, state := cache.Get(order.Order_uid); !state {
				cache.Set(order.Order_uid, order)
				database.InsertOrder(db, order)
			} else {
				if !reflect.DeepEqual(order, database.GetOrder(db, order.Order_uid)) {
					database.UpdateOrder(db, order)
					cache.Set(order.Order_uid, order)
				}
			}

		}

	}
	return f
}
