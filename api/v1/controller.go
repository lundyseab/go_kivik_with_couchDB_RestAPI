package v1

import (
	"context"
	"net/http"
	database "bc_hw3/initialize"
	"github.com/gin-gonic/gin"
)


type newDoc struct {
	Name string `json:"name"`
	Age   int    `json:"age"`
}

func InsertDoc(ctx *gin.Context) {

	// ---------------------------
	var doc newDoc

	if err := ctx.BindJSON(&doc); err != nil {
		return
	}

	docID, rev, err := database.DB.CreateDoc(context.TODO(), &doc)
	if err != nil {
		panic(err)
	}
	returnDoc := map[string]interface{}{
		"docId": docID,
		"rev":   rev,
		"name": doc.Name,
		"age": doc.Age,
	}
	ctx.IndentedJSON(http.StatusCreated, returnDoc)
}
