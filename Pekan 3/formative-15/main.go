package main

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controllers "formative-15/controllers"
	database "formative-15/database"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	// ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))
}
