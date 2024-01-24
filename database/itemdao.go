package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Trickster22/L0/models"
)

func InsertItems(db *sql.DB, items []models.Item) {
	for _, item := range items {
		InsertItem(db, item)
	}
}

func GetItems(db *sql.DB, trackNumber string) []models.Item {
	log.Println("Get items from DB by tracknumber =", trackNumber)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM items WHERE track_number ='%s'", trackNumber))
	items := make([]models.Item, 0)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var item models.Item
		err = rows.Scan(
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
		)
		CheckError(err)
		items = append(items, item)
	}
	CheckError(err)
	return items
}

func getItem(db *sql.DB, chrtId int) models.Item {
	log.Println("Get item from DB by chrt_id =", chrtId)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM items WHERE chrt_id = %d", chrtId))
	CheckError(err)
	defer rows.Close()
	var item models.Item
	for rows.Next() {
		err = rows.Scan(
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
		)
		CheckError(err)
	}
	return item
}

func IsItemExist(db *sql.DB, chrtId int) bool {
	var emptyItem models.Item
	return emptyItem != getItem(db, chrtId)
}

func UpdateItem(db *sql.DB, item models.Item) {
	log.Println("Update item in DB")
	updateStr := `UPDATE items SET track_number = $1, price = $2, rid = $3, name = $4, sale = $5, size = $6, total_price = $7, nm_id = $8, brand = $9, status = $10 WHERE chrt_id = $11`
	_, err := db.Exec(
		updateStr,
		item.Track_number,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.Total_price,
		item.Nm_id,
		item.Brand,
		item.Status,
		item.Chrt_id,
	)
	CheckError(err)
	log.Println("Item updated succesfully")
}

func InsertItem(db *sql.DB, item models.Item) {
	log.Println("Insert item in DB")
	insertStr := `INSERT INTO "items"("chrt_id", "track_number", "price", "rid", "name", "sale", "size", "total_price", "nm_id", "brand", "status") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(
		insertStr,
		item.Chrt_id,
		item.Track_number,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.Total_price,
		item.Nm_id,
		item.Brand,
		item.Status,
	)
	CheckError(err)
	log.Println("Item inserted succesfully")
}

func InsertOrUpdateItem(db *sql.DB, item models.Item) {
	if IsItemExist(db, item.Chrt_id) {
		UpdateItem(db, item)
	} else {
		InsertItem(db, item)
	}
}
