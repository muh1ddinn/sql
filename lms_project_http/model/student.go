package model

type Student struct {
	Id         string
	Full_name  string
	Email      string
	Age        int
	Paid_sum   float64
	Status     string
	Login      string
	Password   string
	Group_id   string
	Created_at string
	Updated_at string
}

type Getstudent struct {
	Student []Student
	Count   int64
}
