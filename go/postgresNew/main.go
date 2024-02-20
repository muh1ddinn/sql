package main

import (
	"backend_course/lessons/postgresNew/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()
	defer db.Close()
	/*
	   	inv := storage.NewInventory(db)
	   	country := country.Country{
	   		Name: "Uzbekkistan",
	   		Code: 99888,
	   	}
	   	err := inv.Create(country)
	   	if err != nil {
	   		return
	   	}
	   	fmt.Println("Country created successfully")

	   	countries, err := inv.GetAll()
	   	if err != nil {
	   		return
	   	}
	   	fmt.Println("Countries: ", countries)

	   }
	*/
	n := "7be8319b-03d0-46a7-a00f-c10d14301793"
	inv := storage.NewInventory(db)

	err := inv.DELETE(n)
	if err != nil {
		return
	}
	fmt.Println("Country deleted successfully")

}
func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=muhiddin password=1 database=backend_c sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
