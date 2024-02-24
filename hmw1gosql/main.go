package main

import (
	"database/sql"
	storage "databasewithpsql/Storage"
	"databasewithpsql/models"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	{
		fmt.Println("You can choose create: database's infos, update infos, you can get all infos or you can getbyID infos, and Delete infos:")
		fmt.Println()
		fmt.Println("1: Create database infos")
		time.Sleep(1 * time.Second)
		fmt.Println("2: Update database infos")
		time.Sleep(1 * time.Second)
		fmt.Println("3: GETALL database infos")
		time.Sleep(1 * time.Second)
		fmt.Println("4: GETBYID database infos")
		time.Sleep(1 * time.Second)
		fmt.Println("5: DELETEBYID database infos")
		time.Sleep(1 * time.Second)
	}

	db := connectDB()
	defer db.Close()

	ct := storage.NewCountry(db)

	var choose int

	fmt.Println("please choose one n:")
	fmt.Scan(&choose)

	switch choose {
	case 1:
		Country := models.Country{

			Name: "UZBEK",
			Code: 99888,
		}
		err := ct.Create(Country)

		if err != nil {

			fmt.Println(err, "err")

			return
		}

		fmt.Println("created information to database  ")

	case 2:
		NewCountry := models.Country{
			Name: "Turkey",
			Code: 35466,
		}

		err := ct.Update("b45430d9-6e05-4566-b742-5e7a674ab458", NewCountry)
		if err != nil {

			fmt.Println(err, "err")

			return
		}

		fmt.Println("update information to database  ")
	case 3:

		c, err := ct.GetAll()

		if err != nil {

			fmt.Println(err, "err")

			return
		}
		fmt.Println("Country of information : ", c)

	case 4:
		cout, err := ct.GetByID("b45430d9-6e05-4566-b742-5e7a674ab458")

		if err != nil {
			fmt.Println("err while geting rows", err)
			return
		}

		fmt.Println("Country of information : ", cout)

	case 5:
		err := ct.Delete("3b675b95-c45f-4a49-aaff-723724923252")
		if err != nil {
			fmt.Println("err while deleting time ", err)
			return
		}

		fmt.Println("delete done succesfully")

	default:
		fmt.Println("please choose correct number ?")

	}

}
func connectDB() *sql.DB {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=muhiddin password=1 database=backend_c sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
