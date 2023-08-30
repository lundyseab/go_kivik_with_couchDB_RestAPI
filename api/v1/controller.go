package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/lundyseab/go_kivik_with_couchDB_RestAPI/initialize"
	"github.com/lundyseab/go_kivik_with_couchDB_RestAPI/models"
)


func InsertDoc(ctx *gin.Context) {

	// ---------------------------
	var student models.Student

	if err := ctx.BindJSON(&student); err != nil {
		return
	}

	docID, rev, err := database.DB.CreateDoc(context.TODO(), &student)
	if err != nil {
		panic(err)
	}
	returnDoc := map[string]interface{}{
		"docId": docID,
		"rev":   rev,
		"name": student.Name,
		"age": student.Age,
		"id": student.ID,
	}
	ctx.IndentedJSON(http.StatusCreated, returnDoc)
}
