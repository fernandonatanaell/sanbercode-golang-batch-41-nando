package repository

import (
	"database/sql"
	"quiz-3/structs"
	"time"
)

func GetAllCategories(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		var cAt, uAt string

		err = rows.Scan(&category.ID, &category.Name, &cAt, &uAt)
		if err != nil {
			panic(err)
		}

		createdAt, err := time.Parse(time.RFC3339, cAt)
		if err != nil {
			panic(err.Error())
		}

		updatedAt, err := time.Parse(time.RFC3339, uAt)
		if err != nil {
			panic(err.Error())
		}

		category.Created_at = createdAt
		category.Updated_at = updatedAt

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (name, created_at, updated_at) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql,
		category.Name,
		category.Created_at.Format("2006-01-02 15:04:05"),
		category.Updated_at.Format("2006-01-02 15:04:05"),
	)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"

	errs := db.QueryRow(sql,
		category.Name,
		time.Now().Format("2006-01-02 15:04:05"),
		category.ID,
	)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetCategory(db *sql.DB, ID int) (category structs.Category, err error) {
	sql := "SELECT * FROM category WHERE id = $1 LIMIT 1"

	err = db.QueryRow(sql, ID).Scan(
		&category.ID,
		&category.Name,
		&category.Created_at,
		&category.Updated_at,
	)

	return
}
