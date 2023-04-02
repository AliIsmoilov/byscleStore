package postgresql

import (
	"app/api/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type reportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *reportRepo {
	return &reportRepo{
		db: db,
	}
}


func (r *reportRepo) FromStock_ToStock(ctx context.Context, req *models.FromToStock) (int64, error) {

	var (
		giverQuantity		int
		receiverQuantity	int
	)

	// Minus From Giver Store

	err := r.db.QueryRow(ctx, 
		"SELECT quantity FROM stocks WHERE store_id = $1 AND product_id = $2",
		req.FromStore,
		req.ProductId,
	).Scan(&giverQuantity)

	if err != nil{
		return 0, err
	}

	if giverQuantity < req.Quantity{
		return 0, errors.New("not enough products in giver store")
	}


	query1 := `
		UPDATE
			stocks
		SET
			quantity = $1
		WHERE store_id = $2 AND product_id = $3
	`


	res1, err := r.db.Exec(ctx, query1, 
		giverQuantity - req.Quantity,
		req.FromStore,
		req.ProductId,
	)

	if err != nil{
		return 0, err
	}

	// Add to Reciever Store

	err = r.db.QueryRow(ctx, 
		"SELECT quantity FROM stocks WHERE store_id = $1 AND product_id = $2",
		req.ToStore,
		req.ProductId,
	).Scan(&receiverQuantity)
	
	if err != nil{
		return 0, err
	}

	res2, err := r.db.Exec(ctx, query1, 
		receiverQuantity + req.Quantity,
		req.ToStore,
		req.ProductId,
	)

	if err != nil{
		return 0, err
	}

	return res1.RowsAffected()+res2.RowsAffected(), nil
}


func (r *reportRepo) GetListStaffSell(ctx context.Context, req *models.GetListStaffSellRequest) (resp *models.GetListStaffSellResponse, err error) {

	resp = &models.GetListStaffSellResponse{}

	var (
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT 
		COUNT(*) OVER(),
		s.first_name || ' ' || s.last_name AS staff_name,
		c.category_name,
		p.product_name,
		ori.quantity,
		ori.list_price,
		CAST(o.order_date::timestamp AS VARCHAR)
	FROM
		orders AS o
	JOIN order_items AS ori ON o.order_id = ori.order_id
	JOIN products AS p ON p.product_id = ori.product_id
	JOIN categories AS c ON c.category_id = p.category_id
	JOIN staffs AS s ON s.staff_id = o.staff_id
	WHERE o.order_status = 4
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}

	for rows.Next() {
		
		var staff_repo models.StaffSell

		err = rows.Scan(
			&resp.Count,
			&staff_repo.StaffName,
			&staff_repo.Category,
			&staff_repo.Product,
			&staff_repo.Quantity,
			&staff_repo.TotalPrice,
			&staff_repo.Date,
		)

		if err != nil{
			return nil, err
		}

		resp.StaffsReport = append(resp.StaffsReport, &staff_repo)
	}

	return resp, nil
}



func (r *reportRepo) Total_sum(ctx context.Context, req *models.TotalSum) (*models.TotalSumRes, error) {

	var (
		query			string
		exist			int
		promo_code		models.PromoCode
		original_price	float64
		sellPrice		float64
		discount		float64

	)


	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM orders WHERE order_id = $1", req.Order_id).Scan(&exist)
	if err != nil{
		return nil, err
	}

	if exist < 1{
		return nil, errors.New("no order whith such order_id")
	}

	query = `
	SELECT
    	o.order_id,
    	oi.quantity,
    	oi.list_price
	FROM
	    orders AS o
	JOIN order_items AS oi ON o.order_id = oi.order_id
	WHERE o.order_status = 4 AND o.order_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.Order_id)
	if err != nil{
		return nil, err
	}

	

	for rows.Next() {

		var queryResp models.Totalumquery

		err = rows.Scan(
			&queryResp.Order_id,
			&queryResp.Quantity,
			&queryResp.ListPrice,
		)

		if err != nil{
			return 	nil, err
		}

		original_price += queryResp.ListPrice * float64(queryResp.Quantity)

	}

	
	if len(req.Promo_code) > 0{

		
		err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM promo_code WHERE name = $1", req.Promo_code).Scan(&exist)
		if err != nil{
			return nil, err
		}
		fmt.Println(exist)

		if exist < 1{
			return nil, errors.New("invalid Promocode")
		}


		query1 := `
		SELECT 
			promocode_id,
			name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code
		WHERE name = $1`

		err = r.db.QueryRow(context.Background(), query1, req.Promo_code).Scan(
			&promo_code.PromoCodeId,
			&promo_code.Name,
			&promo_code.Discount,
			&promo_code.Discount_Type,
			&promo_code.Ordred_limit_price,
		)

		if err != nil{
			return nil, err
		}

		if original_price > promo_code.Ordred_limit_price{

			if promo_code.Discount_Type == "fixed"{
				sellPrice = (original_price - promo_code.Discount)
			} else if promo_code.Discount_Type == "percent" {
				sellPrice = original_price - (original_price * promo_code.Discount)
			}
	
			discount = promo_code.Discount
		}
	}

	// sell_price := req.ListPrice - (req.ListPrice * req.Discount)


	return &models.TotalSumRes{
		OriginalPrice: original_price,
		SellPrice: sellPrice,
		Discount: discount,
	}, err

}

