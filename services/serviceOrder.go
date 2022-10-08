package services

import (
	"Assignment_2/database"
	"Assignment_2/models"
	"Assignment_2/requests"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"
)

func SaveOrder(req requests.Request) (requests.Response, error) {
	var res requests.Response

	// Dapatkan Database
	db := database.GetDb()

	// Items
	var items []models.Item

	// Declaring total
	var total int64
	for _, v_item := range req.Items {
		// Defining Item
		var item models.Item

		// Fill Item
		item.Price = v_item.Price
		item.Price = v_item.Price
		item.Quantity = int(v_item.Quantity)
		item.ItemCode = v_item.ItemCode
		item.Description = v_item.Description

		// Adding to array
		items = append(items, item)
		total += (v_item.Quantity * v_item.Price)
	}

	// Calling Order
	order := models.Order{
		OrderID:      generateRandomOrderId(),
		CustomerName: req.CustomerName,
		OrderAt:      0,
		DetailItem:   items,
	}

	errdb := db.Create(&order).Error

	// Jika Error
	if errdb != nil {
		return res, errdb
	}

	return requests.Response{
		Data:         req,
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		OrderID:      order.OrderID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        total,
	}, nil
}

func AllOrder() (requests.ResponseGet, error) {

	// Dapatkan Database
	db := database.GetDb()

	// Initiating Order
	var order []requests.Order

	// Get All order
	db.Joins("left JOIN items ON items.order_item = orders.id").Find(&order)

	fmt.Println("tes: ", order)

	// Response
	return requests.ResponseGet{
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		Orders:       order,
		ResponseCode: "00",
		Status:       "Success",
		Total:        int64(len(order)),
	}, nil
}

// Generate ID
func generateRandomOrderId() string {
	b := make([]byte, 16)

	_, err := rand.Read(b)

	if err != nil {
		return fmt.Sprintf("%v", currentTime())
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Epoch Time
func currentTime() int64 {
	return time.Now().Unix()
}

// Date Time Epoch
func dateTimeEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err == nil {
		fmt.Println(b)
	}
	return
}
