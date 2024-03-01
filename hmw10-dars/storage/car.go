package storage

import (
	"cars_with_sql/model"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type carRepo struct {
	DATA *sql.DB
}

func Newcar(DATA *sql.DB) carRepo {
	return carRepo{
		DATA: DATA,
	}
}

/*
======================================
create (body) id ,err
update(body) id,err
delete(id) err
get(id) body,err
getALL(serach) []body,count,err

======================================
*/
func (c *carRepo) Create(car model.Car) (string, error) {

	id := uuid.New()

	query := `INSERT INTO cars(
        id,
        name,
        brand,
        model,
		year,
        hourse_power,
        colour,
        engine_cap)
    VALUES($1,$2,$3,$4,$5,$6,$7,$8) `

	res, err := c.DATA.Exec(query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.Year,
		car.HoursePower, car.Colour, car.EngineCap)

	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)

	}
	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

func (c *carRepo) GetAll() ([]model.Car, error) {
	carr := []model.Car{}
	rows, err := c.DATA.Query(`SELECT 
	name,
	brand,
	model,
	year,
	hourse_power,
	colour,
	engine_cap FROM cars`)
	if err != nil {
		fmt.Println("error while getting all country err: ", err)
		return nil, err
	}

	for rows.Next() {
		car := model.Car{}
		if err = rows.Scan(&car.Name, &car.Model,
			&car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		carr = append(carr, car)
	}

	return carr, nil
}

func (c *carRepo) Getbyid(id string) ([]model.Car, error) {
	carrr := []model.Car{}
	rows, err := c.DATA.Query(`SELECT 
	name,
	brand,
	model,
	year,
	hourse_power,
	colour,
	engine_cap FROM cars`)
	if err != nil {
		fmt.Println("error while getting all country err: ", err)
		return carrr, err
	}

	for rows.Next() {
		car := model.Car{}
		if err = rows.Scan(&car.Name, &car.Model,
			&car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap); err != nil {
			fmt.Println("error while scanning country err: ", err)
			return nil, err
		}
		carrr = append(carrr, car)
	}

	return carrr, nil
}
