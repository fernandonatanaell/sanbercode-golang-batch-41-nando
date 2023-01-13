package repository

import (
	"database/sql"
	"quiz-3/structs"
	"time"
)

func GetAllBooks(db *sql.DB) (err error, results []structs.Book) {
	sql := "SELECT * FROM book"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		var cAt, uAt string

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&cAt,
			&uAt,
			&book.CategoryID)
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

		book.Created_at = createdAt
		book.Updated_at = updatedAt

		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	errs := db.QueryRow(sql,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.Created_at.Format("2006-01-02 15:04:05"),
		book.Updated_at.Format("2006-01-02 15:04:05"),
		book.CategoryID,
	)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, updated_at = $8, category_id = $9 WHERE id = $10"

	errs := db.QueryRow(sql,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		time.Now().Format("2006-01-02 15:04:05"),
		book.CategoryID,
		book.ID,
	)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id = $1"

	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}

func GetBook(db *sql.DB, ID int) (book structs.Book, err error) {
	sql := "SELECT * FROM book WHERE id = $1 LIMIT 1"

	err = db.QueryRow(sql, ID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageUrl,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.Created_at,
		&book.Updated_at,
		&book.CategoryID,
	)

	return
}

func GetAllBooksFromCategory(db *sql.DB, category_id int) (books []structs.Book, err error) {
	sql := "SELECT * FROM book WHERE category_id  = $1"

	rows, err := db.Query(sql, category_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		var cAt, uAt string

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&cAt,
			&uAt,
			&book.CategoryID)
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

		book.Created_at = createdAt
		book.Updated_at = updatedAt

		books = append(books, book)
	}

	return
}
