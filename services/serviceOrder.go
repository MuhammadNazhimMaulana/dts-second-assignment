package services

import (
	"Assignment_2/database"
	"Assignment_2/models"
	"Assignment_2/requests"
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
	var order []models.Order

	// Get All order
	db.Preload("DetailItem").Find(&order)

	// Response
	return requests.ResponseGet{
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		Orders:       order,
		ResponseCode: "00",
		Status:       "Success",
		Total:        int64(len(order)),
	}, nil
}

func Order(id int) (requests.ResponseGetOne, error) {

	// Dapatkan Database
	db := database.GetDb()

	// Initiating Order
	var order models.Order

	// Get One order
	db.Preload("DetailItem").Find(&order, id)

	// Response
	return requests.ResponseGetOne{
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		Orders:       order,
		ResponseCode: "00",
		Status:       "Success",
		Total:        1,
	}, nil
}

func UpdateOrder(req requests.Request, id int) (requests.Response, error) {
	var res requests.Response

	// Dapatkan Database
	db := database.GetDb()

	// Initiating Order
	var oldOrder models.Order

	// Get One order
	db.Preload("DetailItem").Find(&oldOrder, id)

	// Delete Items
	db.Where("order_item=?", id).Delete(&models.Item{})

	// Detail Item
	oldOrder.DetailItem = []models.Item{}

	// Declaring total
	var total int64
	for _, v_item := range req.Items {
		// Defining Item
		var item models.Item

		// Fill Item
		item.OrderItem = id
		item.Price = v_item.Price
		item.Quantity = int(v_item.Quantity)
		item.ItemCode = v_item.ItemCode
		item.Description = v_item.Description

		// Save New Item
		db.Save(&item)
		total += (v_item.Quantity * v_item.Price)
	}

	// Calling Order
	oldOrder.CustomerName = map[bool]string{true: oldOrder.CustomerName, false: req.CustomerName}[req.CustomerName == ""]
	oldOrder.OrderAt = map[bool]int64{true: 0, false: oldOrder.OrderAt}[oldOrder.OrderAt == 0]

	// Save Update Order
	err := db.Save(&oldOrder).Error

	// Jika Error
	if err != nil {
		return res, err
	}

	return requests.Response{
		Data:         req,
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		OrderID:      oldOrder.OrderID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        total,
	}, nil
}

func Delete(id int) (requests.ResponseDelete, error) {

	// Dapatkan Database
	db := database.GetDb()

	// Delete Item
	db.Where("order_item=?", id).Delete(&models.Item{})

	// Delete Order
	db.Where("order_id=?", id).Delete(&models.Order{})

	// Response
	return requests.ResponseDelete{
		Message:      "Delete Data Berhasil",
		ResponseCode: "00",
		Status:       "Success",
	}, nil
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
