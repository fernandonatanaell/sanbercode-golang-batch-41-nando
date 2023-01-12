package controllers

import (
	"database/sql"
	"fmt"
)

// Struct
type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

// Function
func CreateEmployee(db *sql.DB, err error) {
	var employee = Employee{}

	sqlStatement := `INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@mail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployee(db *sql.DB, err error) {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Printf("Employee Datas : %+v\n", results)
}

func UpdateEmployee(db *sql.DB, err error) {
	sqlStatement := `UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airellhidayat@gmail.com", "CurDevs", 24)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee(db *sql.DB, err error) {
	sqlStatement := `DELETE FROM employees WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount:", count)
}
