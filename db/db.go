package db

import (
	//"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	//"github.com/jmoiron/sqlx"
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

const (
	MAX_LIMIT = 10
)

type Repo interface {
	Get()
}

//var instance *sqlx.DB
var instance *sqlx.DB

func getDbSourceName() (sourceName string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")

	sourceName = fmt.Sprintf("%s:%s@/%s?charset=utf8", user, pass, name)

	//fmt.Println(sourceName)
	return
}

// InitDB This method inits DB
func initDB() {
	var err error

	instance, err = sqlx.Connect("mysql", getDbSourceName())
	//instance, err = sql.Open("mysql", getDbSourceName())

	if err != nil {
		log.Panic(err)
	}

	if err = instance.Ping(); err != nil {
		log.Panic(err)
	}
}

func DB() *sqlx.DB {
	if instance == nil {
		initDB()
	}

	return instance
}

func GetStructDBFields(somestruct interface{}) (selectstring string) {
	var selectQuery []string

	t := reflect.TypeOf(Guest{})
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		selectQuery = append(selectQuery, t.Field(i).Tag.Get("db"))
	}

	selectstring = strings.Join(selectQuery, ",")

	if selectstring == "" {
		log.Panic("GetStructDBFields returned empty for Struct")
	}

	return
}
