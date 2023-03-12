package main

import (
	"BookStore/pkg"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getBooks(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		if strings.Contains(c.Request.URL.String(), "/books?title=") {
			title := c.Request.URL.Query().Get("title")
			var bks []pkg.Book
			db.Find(&bks)
			for _, i := range bks {
				if strings.Contains(i.Title, title) {
					c.JSON(200, i)
				}
			}
		} else {
			var bks []pkg.Book
			db.Find(&bks)
			c.JSON(http.StatusOK, bks)
		}
	}
}

func getBook(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	id := c.Param("id")
	if err != nil {
		panic(err)
	} else {
		var bks []pkg.Book
		db.Find(&bks)
		ok := false
		newId, _ := strconv.Atoi(id)
		for _, i := range bks {
			if i.ID == newId {
				c.JSON(http.StatusOK, i)
				ok = true
			}
		}
		if !ok {
			c.JSON(404, gin.H{"message": "not found!"})
		}
	}
}

func createBook(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		var b pkg.Book
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal([]byte(reqBody), &b)
		db.Create(&b)
		c.JSON(201, gin.H{"message": "created"})
	}
}

func deleteBook(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	id := c.Param("id")
	if err != nil {
		panic(err)
	} else {
		var bks []pkg.Book
		db.Find(&bks)
		ok := false
		newId, _ := strconv.Atoi(id)
		for _, i := range bks {
			if i.ID == newId {
				db.Delete(&i)
				c.JSON(http.StatusOK, gin.H{"message": "deleted"})
				ok = true
			}
		}
		if !ok {
			c.JSON(404, gin.H{"message": "not found!"})
		}
	}

}

func desc(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		var bks []pkg.Book
		db.Order("ID desc").Find(&bks)
		c.JSON(http.StatusOK, bks)
	}
}

func asc(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		var bks []pkg.Book
		db.Order("ID asc").Find(&bks)
		c.JSON(http.StatusOK, bks)
	}
}

func updateBook(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Dimash2003 dbname=go_application1 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		id := c.Param("id")
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		b := make(map[string]string)
		json.Unmarshal([]byte(reqBody), &b)
		IdNew, _ := strconv.Atoi(id)
		db.Model(&pkg.Book{}).Where("iD = ?", IdNew).Update("Title", b["Title"])
		db.Model(&pkg.Book{}).Where("iD = ?", IdNew).Update("Description", b["Description"])
		c.JSON(200, gin.H{"message": "updated"})
	}

}

func main() {
	r := gin.Default()
	r.GET("/books", getBooks)
	r.GET("/books/desc", desc)
	r.GET("/books/asc", asc)
	r.GET("/books/:id", getBook)
	r.POST("/books", createBook)
	r.DELETE("/books/:id", deleteBook)
	r.PATCH("/books/:id", updateBook)
	r.Run()
}

/*
	Get List of books, +
	Upd title or description by ID, +
	Delete book by ID, +
	Search by title, +
	Add book, +
	Sort by desc and asc, +
	Docker
*/
