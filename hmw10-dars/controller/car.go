package controller

import (
	"cars_with_sql/model"
	"fmt"
)

func (c *Controller) Createcaaa() {
	car := getCarInfo()

	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car created successfully with ID: %v\n", id)

}

func getCarInfo() model.Car {

	car := model.Car{}
	fmt.Println(`enter the car datas
	name,
	brand,
	model,
	hourse_power,
	colour,
	engine_cap`)

	// var (
	// 	Name, Brand, Model, Colour string
	// 	Year, HoursePower          int
	// 	EngineCap                  float32
	// )

	fmt.Scan(&car.Name, &car.Model,
		&car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap)

	return car

}

func (c *Controller) Getalll() {

	gett, err := c.Store.Car.GetAll()
	if err != nil {
		return
	}
	fmt.Println("Countries: ", gett)

}

func getcarinfo() string {

	car := string
	fmt.Println(`enter the car get i coloum
	id`)
	fmt.Scan(&car.id)

	return car
}
