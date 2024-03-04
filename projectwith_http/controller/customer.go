package controller

import (
	"cars_with_sql/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.c(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllCars(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateCars(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteCar(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) createcustomer(w http.ResponseWriter, r *http.Request) {

	customer := model.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		errStr := fmt.Sprintf("error while request body,err:%v\n", err)

		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, err)
		return

	}

	id, err := c.Store.Car.Create(customer)
	if err != nil {
		fmt.Println("error while creating car,err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
	}

	handleResponse(w, http.StatusOK, id)

}

/*
func (c Controller) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

		customer := model.Customer{}

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			errStr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
			fmt.Println(errStr)
			handleResponse(w, http.StatusBadRequest, errStr)
			return
		}


		}
	 	customer.Id = r.URL.Query().Get("id")

		err := uuid.Validate(car.Id)
		if err != nil {
			fmt.Println("error while validating, err: ", err)
			handleResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := c.Store.customer.Update(car)
		if err != nil {
			fmt.Println("error while creating car, err: ", err)
			handleResponse(w, http.StatusInternalServerError, err)
			return
		}

		handleResponse(w, http.StatusOK, id)
	}
*/
func (c Controller) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	cars, err := c.Store.(search)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) DeleteCar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Car.Delete(id)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetbyCar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	carrs, err := c.Store.Car.GetAll(id)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, carrs)
}
