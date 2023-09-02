package main

import (
	v1 "github.com/lundyseab/go_kivik_with_couchDB_RestAPI/api/v1"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {
	
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	// CORS 
	r.Use(cors.New(cors.Config{
	  AllowOrigins: []string{"*"},
	}))
	r.MaxMultipartMemory = 1024 * 1024
	// create group
	router := r.Group("/api/v1")

	//add new student
	//api/v1/...
	router.POST("/insert_student_doc", v1.InsertDoc)
	router.POST("/upload_files", v1.UploadFile)
	router.GET("/get_file/:id", v1.GetFileWithID)

	r.Run("localhost:"+os.Getenv("PORT"))
}