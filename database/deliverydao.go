package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Trickster22/L0/models"
)

func InsertDelivery(db *sql.DB, delivery models.Delivery, orderUid string) {
	log.Println("Insert delivery in DB")
	insertStr := `INSERT INTO deliveries ("name", "phone", "zip", "city", "address", "region", "email", "orderuid") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(
		insertStr,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
		orderUid,
	)
	CheckError(err)
	log.Println("delivery inserted succesfully")
}

func GetDelivery(db *sql.DB, orderUid string) models.Delivery {
	log.Println("Get delivery from DB by orderUid =", orderUid)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM deliveries WHERE orderuid ='%s'", orderUid))
	CheckError(err)
	defer rows.Close()
	var delivery models.Delivery
	for rows.Next() {
		var trash string
		err = rows.Scan(
			&delivery.Name,
			&delivery.Phone,
			&delivery.Zip,
			&delivery.City,
			&delivery.Address,
			&delivery.Region,
			&delivery.Email,
			&trash,
		)
		CheckError(err)
	}
	return delivery
}

func isDeliveryExist(db *sql.DB, orderUid string) bool {
	var emptyDelivery models.Delivery
	return emptyDelivery != GetDelivery(db, orderUid)
}

func UpdateDelivery(db *sql.DB, delivery models.Delivery, orderUid string) {
	log.Println("Update delivery in DB by orderUid =", orderUid)
	updateStr := `UPDATE deliveries SET name = $1, phone = $2, zip = $3, city = $4, address = $5, region = $6, email = $7 WHERE orderuid = $8`
	_, err := db.Exec(
		updateStr,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
		orderUid,
	)
	CheckError(err)
	log.Println("delivery updated succesfully")
}

func InsertOrUpdateDelivery(db *sql.DB, order models.Order) {
	if isDeliveryExist(db, order.Order_uid) {
		UpdateDelivery(db, order.Delivery, order.Order_uid)
	} else {
		InsertDelivery(db, order.Delivery, order.Order_uid)
	}
}
