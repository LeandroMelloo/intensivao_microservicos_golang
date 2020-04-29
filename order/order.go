package main

import (
	"encoding/json"
	"fmt"
	"order/db"
	"order/queue"
	"os"
	"time"

	uuid "github.com/gofrs/uuid"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float32 `json:"price,string"`
}

type Order struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductId string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,string"`
}

var productsUrl string

func init() {
	productsUrl = os.Getenv("PRODUCT_URL")
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("checkout_queue", connection, in)

	for payload := range in {
		fmt.Println(string(payload))
	}
}

func createOrder(payload []byte) Order {
	var order Order
	json.Unmarshal(payload, &order)

	uuid, _ := uuid.NewV4()
	order.Uuid = uuid.String()
	order.Status = "pendente"
	order.CreatedAt = time.Now()
	saveOrder(order)
	return order
}

func saveOrder(order Order) {
	fmt.Println(order.Uuid)
	json, _ := json.Marshal(order)
	fmt.Println(json)
	connection := db.Connect()

	err := connection.Set(order.Uuid, string(json), 0).Err()
	if err != nil {
		panic(err.Error())
	}
}
