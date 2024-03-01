package controller

import "cars_with_sql/storage"

type Controller struct {
	Store storage.Store
}

func NewController(store storage.Store) Controller {
	return Controller{
		Store: store,
	}
}
