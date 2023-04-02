package models

type StaffSell struct{
	StaffName	string	`json:"staff_name"`
	Category	string	`json:"category"`
	Product 	string	`json:"product"`
	Quantity	int		`json:"quantity"`
	TotalPrice	float64	`json:"total_price"`
	Date		string	`json:"date"`
}

type GetListStaffSellResponse struct {
	Count			int				`json:"count"`
	StaffsReport	[]*StaffSell	`json:"staffs_report"`
}

type GetListStaffSellRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type TotalSum struct {
	Order_id	int	`json:"order_id"`
	Promo_code	string	`json:"promo_code"`
}

type Totalumquery struct {
	Order_id	int		`json:"order_id"`
	Quantity	int		`json:"quantity"`
	ListPrice	float64	`json:"list_price"`
}

type TotalSumRes struct {
	OriginalPrice	float64	`json:"original_price"`
	Discount		float64	`json:"discount"`
	SellPrice		float64	`json:"sell_price"`
}