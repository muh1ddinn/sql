package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	storage "pssql.working/hmw/Storage"
)

func main() {
	db := connectDB()
	defer db.Close()
	ca := storage.Newcategory(db)

	Category := models.Category{

		Name: "TURKEY",
	}
	err := ca.Create(Category)

	if err != nil {
		fmt.Println(err, "err")

		return
	}

	fmt.Println("category created succesfully")
	
	categories, err := ca.GetByID("7b21d19c-8888-446b-97a7-ed73d3547421")
	if err != nil {
		return
	}
	fmt.Println("Categories: ", categories)

}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=muhiddin password=1 database=backend_c sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
