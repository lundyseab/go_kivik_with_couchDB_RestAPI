package v1

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/lundyseab/go_kivik_with_couchDB_RestAPI/initialize"
	"github.com/lundyseab/go_kivik_with_couchDB_RestAPI/models"
)

// insert student documents
func InsertDoc(ctx *gin.Context) {

	// ---------------------------
	var student models.Student

	if err := ctx.BindJSON(&student); err != nil {
		returnDoc := map[string]interface{}{
			"description": "binding file failed, your data must following this form of json {\"id\" : \"value\", \"name\": \"value\", \"age\" : 21}",
			"status": 200,
		}
		ctx.IndentedJSON(http.StatusOK, returnDoc)
	}

	docID, rev, err := database.DB.CreateDoc(context.TODO(), &student)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"docId": docID,
		"rev":   rev,
		"name": student.Name,
		"age": student.Age,
		"id": student.ID,
		"classroom": student.Classroom,
	}
	returnDoc := map[string]interface{}{
		"body": body,
		"status": 200,
	}
	ctx.IndentedJSON(http.StatusOK, returnDoc)
}

// upload one or multiple file to couchDB
func UploadFile(c *gin.Context){

	form, _ := c.MultipartForm()

	files := form.File["files"]  
  
	docs := make([]interface{}, len(files))
  
	for i, file := range files {
  
	  doc := make(map[string]interface{})
	  
	  doc["name"] = file.Filename
	  
	  data, _ := file.Open()
	  bytes, _ := ioutil.ReadAll(data)
	  doc["data"] = base64.StdEncoding.EncodeToString(bytes)
  
	  docs[i] = doc
	}

	 // Insert docs
	 _, err := database.DB.BulkDocs(context.TODO(), docs)
	 if err != nil {
		returnDoc := map[string]interface{}{
			"description": "insert files failed",
			"status": 200,
		}
		c.IndentedJSON(200, returnDoc)
	 }  
	length := len(docs)
	returnDoc := map[string]interface{}{
		"description": strconv.Itoa(length) +" files uploaded successfully",
		"status": 200,
	}
	c.IndentedJSON(200, returnDoc)
  
}

//get document by id 
func GetDocumentById(c *gin.Context){
	id := c.Param("id")

	var student interface{}

	err := database.DB.Get(context.TODO(), id).ScanDoc(&student)
	if err != nil {
		returnDoc := map[string]interface{}{
			"description": "document with ID: "+id +", is not found!",
			"status": 200,
		}
		c.IndentedJSON(200, returnDoc)
	}
	returnDoc := map[string]interface{}{
		"id" : id,
		"body": student,
		"status": 200,
	}
	c.IndentedJSON(200, returnDoc)
}

// get file with id
func GetFileWithID(c *gin.Context){
	id := c.Param("id")

	 // Get document from CouchDB
	client := database.DB
	var doc models.FileDoc
	// Get document with ID "john"
	err := client.Get(context.TODO(), id).ScanDoc(&doc)
	if err != nil {
		returnDoc := map[string]interface{}{
			"description": "file with ID: "+id +", is not found!",
			"status": 200,
		}
		c.IndentedJSON(200, returnDoc)
	}
	
	// Decode data
	file, err := base64.StdEncoding.DecodeString(doc.Data)
	if err != nil {
		returnDoc := map[string]interface{}{
			"description": "Sorry,Decoding fail",
			"status": 200,
		}
		c.IndentedJSON(200, returnDoc)
	}

	 // Set headers for download back with original filename
	 c.Header("Content-Disposition", "attachment; filename="+doc.Name)

	//  return file for download
	c.Data(200, "application/octet-stream", file)

}

func UpdateDocumentByIdAndRev(c *gin.Context){
		id := c.Param("id")
		// Document to update
		var student models.Student

		if err := c.BindJSON(&student); err != nil {
			returnDoc := map[string]interface{}{
				"description": "failed",
				"status": 200,
			}
			c.IndentedJSON(http.StatusOK, returnDoc)
		}
	  
		// Update document using Update()
		rev, err := database.DB.Put(context.TODO(), id, student)
		if err != nil {
			returnDoc := map[string]interface{}{
				"description": "failed to update student document",
				"status": 200,
			}
			c.IndentedJSON(http.StatusOK, returnDoc)
		}
		student.Rev = rev

		returnDoc := map[string]interface{}{
			"body": student,
			"status": 200,
		}
		c.IndentedJSON(http.StatusOK, returnDoc)

}
// func deleteById(id, rev string) {
// 	newRev, err := DBCon.Delete(context.TODO(), id, rev)
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Printf("The tombstone document has revision %s\n", newRev)
//   }
