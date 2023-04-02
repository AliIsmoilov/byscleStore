package models

type PromoCode struct {
	PromoCodeId			int    	`json:"promocode_id"`
	Name 				string 	`json:"name"`
	Discount			float64	`json:"discount"`
	Discount_Type		string	`json:"discount_type"`
	Ordred_limit_price	float64	`json:"order_limit_price"`
}

type PromoCodePrimaryKey struct {
	PromoCodeId int `json:"promocode_id"`
}

type CreatePromoCode struct {
	Name 				string 	`json:"name"`
	Discount			float64	`json:"discount"`
	Discount_Type		string	`json:"discount_type"`
	Ordred_limit_price	float64	`json:"order_limit_price"`
}

type UpdatePromoCode struct {
	PromoCodeId			int    	`json:"promocode_id"`
	Name 				string 	`json:"name"`
	Discount			float64	`json:"discount"`
	Discount_Type		string	`json:"discount_type"`
	Ordred_limit_price	float64	`json:"order_limit_price"`
}

type GetListPromoCodeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromoCodeResponse struct {
	Count  int      `json:"count"`
	PromoCodes 		[]*PromoCode `json:"promocodes"`
}
