package main

import (
	v1 "bc_hw3/api/v1"
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

	// create group 
	router := r.Group("/api/v1")
	router.POST("/insertDoc", v1.InsertDoc)

	r.Run("localhost:"+os.Getenv("PORT"))
}