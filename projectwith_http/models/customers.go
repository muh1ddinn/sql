package model

type Customers struct {
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

type Getcus struct {
	Coustomer []Customers
	Count     int64
}

"id" uuid PRIMARY KEY,
  "full_name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "age" int NOT NULL,
  "status" varchar(60) NOT NULL CHECK("status" IN ('active', 'inactive')) DEFAULT 'active',
  "login" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP