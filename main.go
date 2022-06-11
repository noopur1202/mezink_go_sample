package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//postgres credentials
const (
	DB_USER     = "username"
	DB_PASSWORD = "password"
	DB_NAME     = "fullstack_api"
	DB_HOST     = "sample_db_postgres"
)

func main() {

	router := gin.Default()
	router.POST("/records", getRecords)
	router.Run(":8081")
}

// getRecords responds with the list of all records from postgres db as JSON.
func getRecords(c *gin.Context) {
	var records []Record
	var requestBody Request

	//get reuest payload values
	if err := c.BindJSON(&requestBody); err != nil {
		handleErr(err)
	}
	fmt.Println("Request payload data", requestBody)

	//setting up to postgresDB
	db := setupDB()
	fmt.Println("Getting all records.................")

	//sql query to get records from table
	rows, err := db.Query("SELECT id,name,(SELECT SUM(m) FROM UNNEST(marks) m),created_at FROM records")
	handleErr(err)

	// scanning for each row in record and add to record slice
	for rows.Next() {
		var id int64
		var name string
		var createdAt time.Time
		var totalMarks int64

		err = rows.Scan(&id, &name, &totalMarks, &createdAt)
		handleErr(err)

		//adding only values netween minimum and maximum marks to response slice
		if totalMarks > int64(requestBody.MinCount) && totalMarks < int64(requestBody.MaxCount) {
			records = append(records, Record{ID: id, CreatedAt: createdAt, TotalMarks: totalMarks})
		}
	}

	//sending response to handler
	var response = Response{Code: 0, Message: "Success", Records: records}
	c.IndentedJSON(http.StatusOK, response)
}

// postgres DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	handleErr(err)
	return db
}

// Function for handling errors
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
