package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/configg"
)

type Store struct {
	DB       *sql.DB
	Students studentsLms
}

func New(cfg configg.Config) (Store, error) {

	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,

		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	fmt.Println("err opening :", err)
	if err != nil {
		return Store{}, err
	}

	newstudents := Newstudent(db)

	return Store{
		DB:       db,
		Students: newstudents,
	}, nil
}
