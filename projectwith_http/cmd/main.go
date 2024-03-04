package main

import (
	"cars_with_sql/config"
	"cars_with_sql/controller"
	"cars_with_sql/storage"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	c := controller.NewController(store)

	http.HandleFunc("/car", c.Car)
	fmt.Println("programm is running on localhost:2100...")
	http.ListenAndServe(":2100", nil)

	/*c.Createcaaa()
	//c.Getalll()
	//c.Getbyidd("0658679e-0d24-4db8-8394-af38f95a027e")
	//c.Deletee("020df0f2-3d4c-48be-b182-61a7ff801195")
	c.Updatee()

	*/

}
