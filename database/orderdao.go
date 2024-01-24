package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Trickster22/L0/models"
)

func InsertOrder(db *sql.DB, order models.Order) {
	log.Println("Insert order in DB")
	insertStr := `INSERT INTO orders("orderuid", "tracknumber","entry", "locale", "internalsignature", "customerid", "delivery_service", "shardkey", "smid", "datecreated", "oofshard") 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(
		insertStr,
		order.Order_uid,
		order.Track_number,
		order.Entry,
		order.Locale,
		order.Internal_signature,
		order.Customer_id,
		order.Delivery_service,
		order.Shardkey,
		order.Sm_id,
		order.Date_created,
		order.Oof_shard,
	)
	CheckError(err)

	InsertPayment(db, order.Payment)
	InsertItems(db, order.Items)
	InsertDelivery(db, order.Delivery, order.Order_uid)
	log.Println("Order inserted succesfully")
}

func GetOrder(db *sql.DB, orderUid string) models.Order {
	log.Println("Get order from DB by orderUid =", orderUid)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM orders WHERE orderuid = '%s'", orderUid))
	CheckError(err)
	defer rows.Close()
	var order models.Order
	for rows.Next() {
		err = rows.Scan(
			&order.Order_uid,
			&order.Track_number,
			&order.Entry,
			&order.Locale,
			&order.Internal_signature,
			&order.Customer_id,
			&order.Delivery_service,
			&order.Shardkey,
			&order.Sm_id,
			&order.Date_created,
			&order.Oof_shard,
		)
		CheckError(err)
	}
	order.Delivery = GetDelivery(db, orderUid)
	order.Payment = GetPayment(db, orderUid)
	order.Items = GetItems(db, order.Track_number)
	return order
}

func GetOrders(db *sql.DB) []models.Order {
	log.Println("Get all orders from DB")
	rows, err := db.Query("SELECT * FROM orders")
	CheckError(err)
	defer rows.Close()
	orders := make([]models.Order, 0)
	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Order_uid,
			&order.Track_number,
			&order.Entry,
			&order.Locale,
			&order.Internal_signature,
			&order.Customer_id,
			&order.Delivery_service,
			&order.Shardkey,
			&order.Sm_id,
			&order.Date_created,
			&order.Oof_shard,
		)
		CheckError(err)
		orderUid := order.Order_uid
		order.Delivery = GetDelivery(db, orderUid)
		order.Payment = GetPayment(db, orderUid)
		order.Items = GetItems(db, order.Track_number)
		orders = append(orders, order)
	}
	return orders
}

func UpdateOrder(db *sql.DB, order models.Order) {
	log.Println("Update order in DB")
	updateStr := `UPDATE orders SET tracknumber = $1, entry = $2, locale = $3, internalsignature = $4, 
	customerid = $5, delivery_service = $6, shardkey = $7, smid = $8, datecreated = $9, oofshard = $10
	WHERE orderuid = $11`
	_, err := db.Exec(
		updateStr,
		order.Track_number,
		order.Entry,
		order.Locale,
		order.Internal_signature,
		order.Customer_id,
		order.Delivery_service,
		order.Shardkey,
		order.Sm_id,
		order.Date_created,
		order.Oof_shard,
		order.Order_uid,
	)
	CheckError(err)
	InsertOrUpdateDelivery(db, order)
	InsertOrUpdatePayment(db, order)
	for _, item := range order.Items {
		InsertOrUpdateItem(db, item)
	}
	log.Println("Order updated succesfully")
}
