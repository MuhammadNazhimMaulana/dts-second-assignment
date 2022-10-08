package requests

type Request struct {
	CustomerName string `json:"customerName"`
	Items        []Item `json:"items"`
}

type Order struct {
	CustomerName string `json:"customerName"`
	OrderID      string `json:"order_id"`
	DetailItem   []Item `gorm:"-"`
}

type Item struct {
	Description string `json:"description"`
	ItemCode    string `json:"itemCode"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}

type Response struct {
	Item
	Data         Request `json:"data"`
	DateTrans    string  `json:"dateTrans"`
	OrderID      string  `json:"orderID"`
	ResponseCode string  `json:"responseCode"`
	Status       string  `json:"status"`
	Total        int64   `json:"total"`
}

type ResponseGet struct {
	DateTrans    string  `json:"dateTrans"`
	Orders       []Order `json:"order"`
	ResponseCode string  `json:"responseCode"`
	Status       string  `json:"status"`
	Total        int64   `json:"total"`
}
