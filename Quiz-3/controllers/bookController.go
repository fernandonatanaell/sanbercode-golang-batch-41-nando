package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetAllBooks(c *gin.Context) {
	var result gin.H

	categories, err := repository.GetAllBooks(database.DBConnection)

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

func InsertBook(c *gin.Context) {
	var book structs.Book
	var result string

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	pattern := `^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	re := regexp.MustCompile(pattern)
	match := re.MatchString(book.ImageUrl)

	checkReleaseYear := (book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021)

	if !match || !checkReleaseYear {
		if !match {
			result = "Image_url is not a valid URL!"
		}

		if !checkReleaseYear {
			if result != "" {
				result += " & "
			}

			result += "Release year must between 1980 and 2021!"
		}
	} else {
		if _, errCategory := repository.GetCategory(database.DBConnection, book.CategoryID); errCategory != nil {
			result = "Category not found!"
		} else {
			book.Created_at = time.Now()
			book.Updated_at = time.Now()
			book.Thickness = getThickness(book.TotalPage)

			err = repository.InsertBook(database.DBConnection, book)
			if err != nil {
				result = err.Error()
			} else {
				result = "Success Insert Book"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param(("id")))
	var result string

	if book, errBook := repository.GetBook(database.DBConnection, id); errBook != nil {
		result = "Book not found!"
	} else {
		err := c.ShouldBindJSON(&book)
		if err != nil {
			panic(err)
		}

		pattern := `^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
		re := regexp.MustCompile(pattern)
		match := re.MatchString(book.ImageUrl)

		checkReleaseYear := (book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021)

		if !match || !checkReleaseYear {
			if !match {
				result = "Image_url is not a valid URL!"
			}

			if !checkReleaseYear {
				if result != "" {
					result += " & "
				}

				result += "Release year must between 1980 and 2021!"
			}
		} else {
			book.Thickness = getThickness(book.TotalPage)

			err = repository.UpdateBook(database.DBConnection, book)
			if err != nil {
				result = err.Error()
			} else {
				result = "Success Update Book"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))
	var result string

	if book, errBook := repository.GetBook(database.DBConnection, id); errBook != nil {
		result = "Book not found!"
	} else {
		err = repository.DeleteBook(database.DBConnection, book)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success Delete Book"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func getThickness(totalPage int) string {
	if totalPage <= 100 {
		return "tipis"
	} else if totalPage >= 101 && totalPage <= 200 {
		return "sedang"
	} else {
		return "tebal"
	}
}
