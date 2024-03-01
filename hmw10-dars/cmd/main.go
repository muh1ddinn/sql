package main

import (
	"cars_with_sql/config"
	"cars_with_sql/controller"
	"cars_with_sql/storage"
	"fmt"
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
	//c.Createcaaa()
	c.Getalll()
}
