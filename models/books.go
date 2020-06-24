package models

import (
	"github.com/OoThan/go-restful-api-with-mongodb/connector"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Book struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Isbn      string        `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title     string        `bson:"title,omitempty" json:"title,omitempty"`
	Author    *Author       `bson:"author" json:"author,omitempty"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

type Author struct {
	FirstName string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty" json:"last_name,omitempty"`
}

type Books []Book

func BookInfo(id bson.ObjectId, bookCollection string) (Book, error) {
	db := connector.GetMongoDB()
	book := Book{}
	err := db.C(bookCollection).Find(bson.M{"_id": id}).One(&book)
	return book, err
}
