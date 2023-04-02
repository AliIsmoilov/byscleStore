package models

type Order struct {
	OrderId      int          `json:"order_id"`
	CustomerId   int          `json:"customer_id"`
	CustomerData *Customer    `json:"customer_data"`
	OrderStatus  int16        `json:"order_status"`
	OrderDate    string       `json:"order_date"`
	RequiredDate string       `json:"required_date"`
	ShippedDate  string       `json:"shipped_date"`
	StoreId      int          `json:"store_id"`
	StoreData    *Store       `json:"store_data"`
	StaffId      int          `json:"staff_id"`
	StaffData    *Staff       `json:"staff_data"`
	OrderItems   []*OrderItem `json:"order_items"`
}

type OrderPrimaryKey struct {
	OrderId int `json:"order_id"`
}

type CreateOrder struct {
	CustomerId   int    `json:"customer_id"`
	OrderStatus  int16  `json:"order_status"`
	OrderDate    string `json:"order_date"`
	RequiredDate string `json:"required_date"`
	ShippedDate  string `json:"shipped_date"`
	StoreId      int    `json:"store_id"`
	StaffId      int    `json:"staff_id"`
}

type UpdateOrder struct {
	OrderId      int    `json:"order_id"`
	CustomerId   int    `json:"customer_id"`
	OrderStatus  int16  `json:"order_status"`
	OrderDate    string `json:"order_date"`
	RequiredDate string `json:"required_date"`
	ShippedDate  string `json:"shipped_date"`
	StoreId      int    `json:"store_id"`
	StaffId      int    `json:"staff_id"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}

// -----------------------ITEM------------------
type OrderItem struct {
	OrderId     int      `json:"order_id"`
	ItemId      int      `json:"item_id"`
	ProductId   int      `json:"product_id"`
	ProductData *Product `json:"product_data"`
	Quantity    int      `json:"quantity"`
	ListPrice   float64  `json:"list_price"`
	Discount    float64  `json:"discount"`
	SellPrice	float64	 `json:"sell_price"`	
}

type OrderItemPrimaryKey struct {
	OrderId int `json:"order_id"`
	ItemId  int `json:"item_id"`
}

type CreateOrderItem struct {
	OrderId int `json:"order_id"`
	// ItemId      int     `json:"item_id"`
	ProductId   int      `json:"product_id"`
	// ProductData *Product `json:"product_data"`
	Quantity    int      `json:"quantity"`
	ListPrice   float64  `json:"list_price"`
	Discount    float64  `json:"discount"`
}
