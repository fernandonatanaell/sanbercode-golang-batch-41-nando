package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var result gin.H

	categories, err := repository.GetAllCategories(database.DBConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category
	var result string

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.Created_at = time.Now()
	category.Updated_at = time.Now()
	err = repository.InsertCategory(database.DBConnection, category)
	if err != nil {
		result = err.Error()
	} else {
		result = "Success Insert Category"
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param(("id")))
	var result string

	if category, errCategory := repository.GetCategory(database.DBConnection, id); errCategory != nil {
		result = "Category not found!"
	} else {
		err := c.ShouldBindJSON(&category)
		if err != nil {
			panic(err)
		}

		err = repository.UpdateCategory(database.DBConnection, category)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success Update Category"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))
	var result string

	if category, errCategory := repository.GetCategory(database.DBConnection, id); errCategory != nil {
		result = "Category not found!"
	} else {
		err = repository.DeleteCategory(database.DBConnection, category)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success Delete Category"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetAllBooksFromCategory(c *gin.Context) {
	var result gin.H
	category_id, err := strconv.Atoi(c.Param(("id")))

	books, err := repository.GetAllBooksFromCategory(database.DBConnection, category_id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}
