package model

type Car struct {
	Id          string
	Name        string
	Year        int
	Brand       string
	Model       string
	HoursePower int
	Colour      string
	EngineCap   float32
	CreatedAt   string
	UpdatedAt   string
}

type GetAllCarsResponse struct {
	Cars  []Car `json:"cars"`
	Count int64 `json:"count"`
}
