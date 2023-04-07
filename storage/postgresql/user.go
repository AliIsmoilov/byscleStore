package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	id := uuid.New().String()

	query := `
		INSERT INTO users(
			id, 
			first_name,
			last_name,
			login,
			password,
			phone_number,
			updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, now()) RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Login,
		req.Password,
		req.Phone_number,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query    string
		user models.User
	)

	query = `
		SELECT
			id, 
			first_name,
			last_name,
			login,
			password,
			phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.UserId).Scan(
		&user.UserId,
		&user.FirstName,
		&user.LastName,
		&user.Login,
		&user.Password,
		&user.Phone_number,
		&user.Created_at,
		&user.Updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			first_name,
			last_name,
			login,
			password,
			phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24-MI-SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24-MI-SS')
		FROM users
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
		var user models.User
		err = rows.Scan(
			&resp.Count,
			&user.UserId,
			&user.FirstName,
			&user.LastName,
			&user.Login,
			&user.Password,
			&user.Phone_number,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (r *userRepo) UpdatePut(ctx context.Context, req *models.UpdateUser) (int64, error) {

	query := `
		UPDATE
		users
		SET
			id = $1, 
			first_name = $2,
			last_name = $3,
			login = $4,
			password = $5,
			phone_number = $6,
			updated_at = now()
		WHERE id = $8
	`

	result, err := r.db.Exec(ctx, query, 
		req.UserId,
		req.FirstName,
		req.LastName,
		req.Login,
		req.Password,
		req.Phone_number,
		req.UserId,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// func (r *customerRepo) UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error) {
// 	var (
// 		query string
// 		set   string
// 	)

// 	if len(req.Fields) <= 0 {
// 		return 0, errors.New("no fields")
// 	}

// 	i := 0
// 	for key := range req.Fields {
// 		if i == len(req.Fields)-1 {
// 			set += fmt.Sprintf(" %s = :%s ", key, key)
// 		} else {
// 			set += fmt.Sprintf(" %s = :%s, ", key, key)
// 		}
// 		i++
// 	}

// 	query = `
// 		UPDATE
// 		customers
// 		SET
// 		` + set + `
// 		WHERE customer_id = :customer_id
// 	`

// 	req.Fields["customer_id"] = req.ID

// 	query, args := helper.ReplaceQueryParams(query, req.Fields)

// 	fmt.Println(query)

// 	result, err := r.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.RowsAffected(), nil
// }

// func (r *customerRepo) Delete(ctx context.Context, req *models.CustomerPrimaryKey) (int64, error) {
// 	query := `
// 		DELETE 
// 		FROM customers
// 		WHERE customer_id = $1
// 	`

// 	result, err := r.db.Exec(ctx, query, req.CustomerId)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.RowsAffected(), nil
// }
