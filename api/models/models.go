package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Director struct {
	ID        int    `orm:"auto"`
	FirstName string `json:"firstname" orm:"size(255)"`
	LastName  string `json:"lastname" orm:"size(255)"`
}

type Movie struct {
	ID       int       `orm:"auto"`
	Isbn     string    `json:"isbn" orm:"size(255);unique"`
	Title    string    `json:"title" orm:"size(255)"`
	Director *Director `json:"director" orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Director), new(Movie))
	orm.RegisterDataBase("default", "mysql", "root:12Abhi@264eden@tcp(localhost:3307)/movie_management?charset=utf8")
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connection is established")
}
