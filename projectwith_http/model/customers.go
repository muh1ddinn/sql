package model

type Customer struct {
	Id         string
	First_name string
	Last_name  string
	Gmail      string
	Phone      string
	Is_blocked bool
	Created_at string
	Updated_at string
	Deleted_at int
}

type GetallcustomersResponse struct {
	Coustomer []Customer `json:"Customer"`
	Count     int64      `json:"count"`
}
