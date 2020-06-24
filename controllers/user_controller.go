package controllers

import (
	"errors"
	"github.com/OoThan/go-restful-api-with-mongodb/connector"
	"github.com/OoThan/go-restful-api-with-mongodb/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const BookCollection = "books"

var (
	errNotExit         = errors.New("Books are not exist ")
	errInvalidID       = errors.New("Invalid ID ")
	errInvalidBody     = errors.New("Invalid request body ")
	errInsertionFailed = errors.New("Error in the book insertion ")
	errUpdateFailed    = errors.New("Error in the book update ")
	errDeletionFailed  = errors.New("Error in the deletion ")
)

func GetAllBooks(c *gin.Context) {
	db := connector.GetMongoDB()
	books := models.Books{}
	err := db.C(BookCollection).Find(bson.M{}).All(&books)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExit.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "books": &books})
}

func CreateBook(c *gin.Context) {
	db := connector.GetMongoDB()
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	book.ID = bson.NewObjectId()
	book.UpdatedAt = time.Now()
	book.CreatedAt = time.Now()
	err = db.C(BookCollection).Insert(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "book": &book})
}

func GetBook(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	book, err := models.BookInfo(id, BookCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "book": &book})
}

func UpdateBook(c *gin.Context) {
	db := connector.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	existingBook, err := models.BookInfo(id, BookCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	err = c.Bind(&existingBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingBook.ID = id
	existingBook.UpdatedAt = time.Now()
	err = db.C(BookCollection).Update(bson.M{"_id": &id}, existingBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdateFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "book": &existingBook})
}

func DeleteBook(c *gin.Context) {
	db := connector.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := db.C(BookCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": "success", "message": "Book deleted successfully"})
}
