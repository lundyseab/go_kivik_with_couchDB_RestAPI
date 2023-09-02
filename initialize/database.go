package initialize

import (
	"context"
	"log"
	"os"

	"github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
	"github.com/joho/godotenv"
)


var DB *kivik.DB

func init () {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// connect to database
	DB = ConnectToDB()
}

func ConnectCouchDB() *kivik.Client{
	client, err := kivik.New("couch", os.Getenv("DB_HOST"))

	if err != nil {
		panic(err)
	} else {
		log.Println("Connected!")
	}
	err =  client.Authenticate(context.TODO(), couchdb.BasicAuth(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Println("Authentication failed")
	}
	return client
}

func ConnectToDB() *kivik.DB{
	// connect to couchDB
	client := ConnectCouchDB()

	// create new database
	client.CreateDB(context.TODO(), os.Getenv("DATABASE"))
	// connect to database
	db := client.DB(context.TODO(), os.Getenv("DATABASE"))
	return db
}