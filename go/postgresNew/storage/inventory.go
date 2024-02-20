package storage

import (
	"backend_course/lessons/postgresNew/country"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Inventory struct {
	db *sql.DB
}

func NewInventory(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}

func (i *Inventory) Create(c country.Country) error {
	id := uuid.NewString()
	_, err := i.db.Exec(
		`INSERT INTO countryy (id,name,code,created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP)
	`, id, c.Name, c.Code)
	if err != nil {
		fmt.Println("error while creating country err: ", err)
		return err
	}

	return nil
}

func (i *Inventory) Update(c country.Country) error {
	_, err := i.db.Exec(`UPDATE countryy SET
name=$1,
code=$2,
updated_at=CURRENT_TIMESTAMP WHERE id=$3`, c.Name, c.Code, c.Id)
	if err != nil {
		fmt.Println("err while updating countryy err: ", err)

		return err
	}
	return nil
}

func (i *Inventory) GetAll() ([]country.Country, error) {
	countryy := []country.Country{}
	rows, err := i.db.Query(`select 
	id,
	name,
	code,
	created_at from countryy WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all countryy err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := country.Country{}
		if err = rows.Scan(&c.Id, &c.Name, &c.Code, &c.CreatedAt); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		countryy = append(countryy, c)
	}

	return countryy, nil
}

func (i *Inventory) DELETE(n string) error {
	_, err := i.db.Exec(
		"DELETE FROM countryy WHERE id = $1",
		n,
	)
	if err != nil {
		fmt.Println("error while deleting country err: ", err)
		return err
	}

	return nil
}
