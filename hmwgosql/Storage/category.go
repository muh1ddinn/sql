package storage

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"pssql.working/hmw/models"
)

type category struct {
	db *sql.DB
}

func Newcategory(db *sql.DB) category {
	return category{
		db: db,
	}
}

func (c *category) Create(category models.Category) error {
	id := uuid.NewString()
	_, err := c.db.Exec(`
		INSERT INTO category (
		id, 
		name, 
		updated_at,
		created_at
		)VALUES ($1, $2, NOW(),Current_timestamp)
	`, id, category.Name)
	if err != nil {
		fmt.Println("error while getting from db", err)
		return nil
	}
	return nil
}

/*
func (c *category) GetbyID(id string) (models.Category, error) {
	var category models.Category
	err := c.db.QueryRow(`
	SELECT
	id,
	name,
	created_at,
	update_at FROM category
    WHERE id = $1`,
		id).Scan(&category.Id, &category.Name, &category.Created_at, &category.Upated_at)
	if err != nil {

		fmt.Println("err while saning from db ", err)
		return category, err
	}

	return category, nil
} */

func (c *category) GetByID(id string) (models.Category, error) {
	var category models.Category
	err := c.db.QueryRow(`
        SELECT 
            id,
            name,
            created_at,
            updated_at
        FROM 
            category
        WHERE 
            id = $1`,
		id).Scan(&category.Id, &category.Name, &category.Created_at, &category.Upated_at)
	if err != nil {
		fmt.Println("Error while scanning from db:", err)
		return category, err
	}

	return category, nil
}

func (c *category) Update(category models.Category) error {

	_, err := c.db.Exec(`

UPDATE category
SET
name = $1,
updated_at = NOW()
WHERE id = $2
`, category.Name)

	if err != nil {

		fmt.Println("err while updating from db", err)

		return nil
	}

	return nil
}
