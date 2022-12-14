package requests

import "Assignment_2/models"

type Request struct {
	CustomerName string `json:"customerName"`
	Items        []Item `json:"items"`
}

type Item struct {
	Description string `json:"description"`
	ItemCode    string `json:"itemCode"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

type Response struct {
	Data         Request `json:"data"`
	DateTrans    string  `json:"dateTrans"`
	OrderID      uint    `json:"orderID"`
	ResponseCode string  `json:"responseCode"`
	Status       string  `json:"status"`
	Total        int64   `json:"total"`
}

type ResponseGet struct {
	DateTrans    string         `json:"dateTrans"`
	Orders       []models.Order `json:"order"`
	ResponseCode string         `json:"responseCode"`
	Status       string         `json:"status"`
	Total        int64          `json:"total"`
}

type ResponseGetOne struct {
	DateTrans    string       `json:"dateTrans"`
	Orders       models.Order `json:"order"`
	ResponseCode string       `json:"responseCode"`
	Status       string       `json:"status"`
	Total        int64        `json:"total"`
}

type ResponseDelete struct {
	Message      string `json:"message"`
	ResponseCode string `json:"responseCode"`
	Status       string `json:"status"`
}
