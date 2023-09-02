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

	
	//api/v1/...
	//add new student
	router.POST("/insert_student_doc", v1.InsertDoc)

	// upload file with bulk insert
	router.POST("/upload_files", v1.UploadFile)

	// get document with id
	router.GET("/get_document_by_id/:id", v1.GetDocumentById)
	//get download with id
	router.GET("/get_file/:id", v1.GetFileWithID)

	// update document by id and rev 
	router.PUT("/update_document/:id", v1.UpdateDocumentByIdAndRev)

	// run server on port from .env
	r.Run("localhost:"+os.Getenv("PORT"))
}