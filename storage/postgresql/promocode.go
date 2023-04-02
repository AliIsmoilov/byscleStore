package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type promoCodeRepo struct {
	db *pgxpool.Pool
}

func NewPromoCodeRepo(db *pgxpool.Pool) *promoCodeRepo {
	return &promoCodeRepo{
		db: db,
	}
}

func (p *promoCodeRepo) CreatePromoCode(ctx context.Context, req *models.CreatePromoCode) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO promo_code(
			promocode_id, 
			name,
			discount,
			discount_type,
			order_limit_price 
		)
		VALUES (
			(
				SELECT COALESCE(MAX(promocode_id), 0) + 1 FROM promo_code
			),
			$1, $2, $3, $4) RETURNING promocode_id
	`
	err := p.db.QueryRow(ctx, query,
		req.Name,
		req.Discount,
		req.Discount_Type,
		req.Ordred_limit_price,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *promoCodeRepo) GetByID(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error) {

	var (
		query    	string
		promocode 	models.PromoCode
	)

	query = `
		SELECT
			promocode_id, 
			name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code
		WHERE promocode_id = $1
	`

	err := p.db.QueryRow(ctx, query, req.PromoCodeId).Scan(
		&promocode.PromoCodeId,
		&promocode.Name,
		&promocode.Discount,
		&promocode.Discount_Type,
		&promocode.Ordred_limit_price,
	)
	if err != nil {
		return nil, err
	}

	return &promocode, nil
}

func (r *promoCodeRepo) GetListPromoCode(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error) {

	resp = &models.GetListPromoCodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			promocode_id, 
			name,
			discount,
			discount_type,
			order_limit_price
		FROM promo_code
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var promocode models.PromoCode
		err = rows.Scan(
			&resp.Count,
			&promocode.PromoCodeId,
			&promocode.Name,
			&promocode.Discount,
			&promocode.Discount_Type,
			&promocode.Ordred_limit_price,
		)
		if err != nil {
			return nil, err
		}

		resp.PromoCodes = append(resp.PromoCodes, &promocode)
	}

	return resp, nil
}

func (r *promoCodeRepo) UpdatePromoCode(ctx context.Context, req *models.UpdatePromoCode) (int64, error) {

	query := `
		UPDATE
			promo_code
		SET
			name = $1,
			discount = $2,
			discount_type = $3,
			order_limit_price = $4
		WHERE promocode_id = $5
	`

	result, err := r.db.Exec(ctx, query, 
		req.Name,
		req.Discount,
		req.Discount_Type,
		req.Ordred_limit_price,
		req.PromoCodeId,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (p *promoCodeRepo) Delete(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM promo_code
		WHERE promocode_id = $1
	`

	result, err := p.db.Exec(ctx, query, req.PromoCodeId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}