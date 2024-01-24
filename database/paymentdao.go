package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Trickster22/L0/models"
)

func GetPayment(db *sql.DB, orderUid string) models.Payment {
	log.Println("Get payment from DB by orderUid =", orderUid)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM payments WHERE transaction ='%s'", orderUid))
	var payment models.Payment
	CheckError(err)
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(
			&payment.Transaction,
			&payment.Requset_id,
			&payment.Currency,
			&payment.Provider,
			&payment.Amount,
			&payment.Payment_dt,
			&payment.Bank,
			&payment.Delivery_cost,
			&payment.Goods_total,
			&payment.Custom_fee,
		)
		CheckError(err)
	}
	return payment
}

func InsertPayment(db *sql.DB, payment models.Payment) {
	log.Println("Insert payment in DB")
	insertStr := `INSERT INTO "payments"("transaction", "requestid", "currency", "provider", "amount", "paymentdt", "bank", "delivery_cost", "goods_total", "custom_fee") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(
		insertStr,
		payment.Transaction,
		payment.Requset_id,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.Payment_dt,
		payment.Bank,
		payment.Delivery_cost,
		payment.Goods_total,
		payment.Custom_fee,
	)
	CheckError(err)
	log.Println("Payment inserted succesfully")
}

func isPaymentExist(db *sql.DB, orderUid string) bool {
	var emptyPayment models.Payment
	return emptyPayment != GetPayment(db, orderUid)
}

func UpdatePayment(db *sql.DB, payment models.Payment) {
	log.Println("Update payment in DB")
	updateStr := `UPDATE payments SET requestid = $1, currency = $2, provider = $3, amount = $4, paymentdt = $5,
	bank = $6, delivery_cost = $7, goods_total = $8, custom_fee = $9 WHERE transaction = $10`
	_, err := db.Exec(
		updateStr,
		payment.Requset_id,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.Payment_dt,
		payment.Bank,
		payment.Delivery_cost,
		payment.Goods_total,
		payment.Custom_fee,
		payment.Transaction,
	)
	CheckError(err)
	log.Println("Payment update succesfully")
}

func InsertOrUpdatePayment(db *sql.DB, order models.Order) {
	if isPaymentExist(db, order.Order_uid) {
		UpdatePayment(db, order.Payment)
	} else {
		InsertPayment(db, order.Payment)
	}
}
