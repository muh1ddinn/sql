package storage

import (
	"database/sql"
	"databasewithpsql/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Country struct {
	db *sql.DB
}

func NewCountry(db *sql.DB) Country {

	return Country{
		db: db,
	}
}

func (ct *Country) Create(countryy models.Country) error {

	id := uuid.NewString()
	_, err := ct.db.Exec(`
	
	INSERT INTO countryy(
    
		id,
		name,
		code,
		created_at,
		updated_at,
		)VALUES(
			$1,$2,$3,CURRENT_TIMESTAMP,NOW())
      `, id, countryy.Name, countryy.Code)
	if err != nil {
		fmt.Println("error while inseting to db ", err)
		return nil
	}

	return nil
}

func (ct *Country) Update(id string, countryy models.Country) error {
	_, err := ct.db.Exec(`
		   UPDATE states SET 
		   name = $1,
		   code = $2,
		   updated_at = NOW()
		   where = $3
   `, countryy.Name, countryy.Code, countryy.Id)

	if err != nil {
		log.Println("error: ", err)
		return nil
	}
	return nil

}

func (ct *Country) GetByID(id string) (models.Country, error) {

	var country models.Country

	err := ct.db.QueryRow(`
	
	SELECT 
	id,
	name,
	code,
	created_at,
	updated_at,
	deleted_at
	FROM countryy
    where id =$1`, id).Scan(&country.Id,
		&country.Name,
		&country.Code,
		&country.Created_at,
		&country.Updated_at,
		&country.Deleted_at)
	if err != nil {

		fmt.Println("err while scanning from db", err)

		return country, err

	}
	return country, nil
}

func (ct *Country) GetAll() ([]models.Country, error) {

	var counttry []models.Country

	ctry, err := ct.db.Query(`
SELECT 
id,
name,
code,
created_at,
updated_at
FROM countryy `)
	if err != nil {
		fmt.Println("error while getting from db", err)
		return counttry, err
	}

	for ctry.Next() {

		ccountry := models.Country{}

		err := ctry.Scan(&ccountry.Id, &ccountry.Name, &ccountry.Code,
			&ccountry.Created_at, &ccountry.Updated_at)

		if err != nil {

			fmt.Println("error while scaning getall time from database ", err)

			return nil, err
		}

		counttry = append(counttry, ccountry)

	}

	return counttry, nil

}

func (ct *Country) Delete(id string) error {

	err := ct.db.QueryRow(`DELETE FROM countryy Where id =$1 `, id)
	if err != nil {

		fmt.Println("error while deleting database ", err)

		return nil

	}
	return nil

}
