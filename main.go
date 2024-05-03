package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Senders struct {
	Id          int    `JSON:"sender_id"`
	Name        string `JSON:"name"`
	PhoneNumber string `JSON:"phone_number"`
	Address     string `JSON:"address"`
	AddressNote string `JSON:"address_note"`
}

type Receivers struct {
	Id          int    `JSON:"receiver_id"`
	Name        string `JSON:"name"`
	PhoneNumber string `JSON:"phone_number"`
	Address     string `JSON:"address"`
	AddressNote string `JSON:"address_note"`
}

type Statuses struct {
	Id   int    `JSON:"status_id"`
	Name string `JSON:"name"`
}

type PaymentMethod struct {
	Id   int    `JSON:"payment_id"`
	Name string `JSON:"name"`
}

type Transaction struct {
	Id             int    `JSON:"transaction_id"`
	OrderNumber    string `JSON:"order_number"`
	SenderId       int    `JSON:"sender_id"`
	ReceiverId     int    `JSON:"receiver_id"`
	ItemTypeId     int    `JSON:"item_type_id"`
	PayementId     int    `JSON:"payment_id"`
	StatusId       int    `JSON:"status_id"`
	DeliveryTypeId int    `JSON:"delivery_type_id"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "cesena10"
	dbname   = "delivery"
)

func main() {
	connectDb := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connectDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/senders", viewSenders)
	router.GET("/statuses", viewStatuses)
	router.GET("/payments", viewPaymentMethod)
	router.GET("/receivers", viewReceivers)
	router.GET("/transactions", viewTransaction)
	router.PUT("/statuses/:order_number", updateStatus)

	router.Run(":8080")

}

func updateStatus(c *gin.Context) {
	orderNumber := c.Param("order_number")

	var status Transaction
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update status transaksi berdasarkan order number
	result, err := db.Exec("UPDATE transactions SET status_id=$1 WHERE order_number=$2", status.StatusId, orderNumber)
	if err != nil {
		log.Printf("Failed to update status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	if rowsAffected == 0 {
		log.Printf("No rows updated for order number: %s", orderNumber)
		c.JSON(http.StatusNotFound, gin.H{"error": "Order number not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

func viewTransaction(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.Id, &transaction.OrderNumber, &transaction.SenderId, &transaction.ReceiverId, &transaction.ItemTypeId, &transaction.PayementId, &transaction.StatusId, &transaction.DeliveryTypeId); err != nil {
			log.Println(err)
			continue
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)

}

func viewReceivers(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM receivers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var receivers []Receivers
	for rows.Next() {
		var receiver Receivers
		if err := rows.Scan(&receiver.Id, &receiver.Name, &receiver.PhoneNumber, &receiver.Address, &receiver.AddressNote); err != nil {
			log.Println(err)

			continue
		}
		receivers = append(receivers, receiver)
	}

	c.JSON(http.StatusOK, receivers)

}

func viewSenders(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM senders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var senders []Senders
	for rows.Next() {
		var sender Senders
		if err := rows.Scan(&sender.Id, &sender.Name, &sender.PhoneNumber, &sender.Address, &sender.AddressNote); err != nil {
			log.Println(err)
			continue
		}
		senders = append(senders, sender)
	}

	c.JSON(http.StatusOK, senders)

}

func viewPaymentMethod(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM payment_method")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var payments []PaymentMethod
	for rows.Next() {
		var payment PaymentMethod
		if err := rows.Scan(&payment.Id, &payment.Name); err != nil {
			log.Println(err)
			continue
		}
		payments = append(payments, payment)
	}

	c.JSON(http.StatusOK, payments)

}

func viewStatuses(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM statuses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var statuses []Statuses
	for rows.Next() {
		var status Statuses
		if err := rows.Scan(&status.Id, &status.Name); err != nil {
			log.Println(err)
			continue
		}
		statuses = append(statuses, status)
	}

	c.JSON(http.StatusOK, statuses)

}
