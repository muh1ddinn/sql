package storage

import (
	"cars_with_sql/model"
	"cars_with_sql/pkg"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type CustomerRepo struct {
	db *sql.DB
}

func Newcustomer(db *sql.DB) CustomerRepo {
	return CustomerRepo{
		db: db,
	}
}

func (c *CustomerRepo) Create(customer model.Customer) (string, error) {

	id := uuid.New()

	query := `INSERT INTO customerss(
        id,
        first_name,
        last_name,
        gmail,
		phone,
        is_blocked,
        )
    VALUES($1,$2,$3,$4,$5,$6) `

	res, err := c.db.Exec(query,
		id.String(),
		customer.First_name, customer.Last_name, customer.Gmail, customer.Phone, customer.Is_blocked)

	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)

	}
	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

func (c carRepo) GettAll(search string) (model.GetallcustomersResponse, error) {
	var (
		resp   = model.GetallcustomersResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.DATA.Query(`select 
				count(id) OVER(),
				id,
				first_name,
				last_name,
				gmail,
				phone,
				is_blocked,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			customer = model.Customer{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&customer.First_name,
			&customer.Last_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_blocked,
			&customer.Created_at,
			&updateAt); err != nil {
			return resp, err
		}

		customer.Updated_at = pkg.NullStringToString(updateAt)
		resp.Coustomer = append(resp.Coustomer, customer)
	}
	return resp, nil
}
func (c *CustomerRepo) GetByyid(id string) ([]model.Customer, error) {
	cus := []model.Customer{}
	rows, err := c.db.Query(`SELECT 
	id,
				first_name,
				last_name,
				gmail,
				phone,
				is_blocked, FROM customerss
	where id=$1`, id)
	if err != nil {
		fmt.Println("error while getting id country err: ", err)
		return cus, err
	}

	for rows.Next() {
		customer := model.Customer{}
		if err = rows.Scan(&customer.First_name,
			&customer.Last_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_blocked); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		cus = append(cus, customer)
	}

	return cus, nil

}

func (c *CustomerRepo) Delete(id string) error {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}

/*
func (c *customerRepo) Update(customer customerRepo) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at=0
	`

	_, err := c.db.Exec(query,
		customer.Last_name,
		customer.Gmail,
		customer.Phone,
		customer.Is_blocked, customer)

	if err != nil {
		return "", err
	}

	return customer.Id, nil
}
*/
