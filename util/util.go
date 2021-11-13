package util

import (
	"encoding/json"
	"fmt"

	"github.com/Fadhelbulloh/local-elastic/model"
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
	"github.com/olivere/elastic/v7"
)

func ErrHandler(code int, c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(code, gin.H{"status": false, "message": err.Error(), "data": nil})
		return true
	}
	return false
}

// response error hendler
func ErrorHandleResponse(c *gin.Context, response model.Response) bool {
	if !response.Status {
		c.JSON(response.StatusCode, response)
		return true
	}
	return false
}

// QueryLog Logging query
func QueryLog(boolQuery *elastic.BoolQuery) bool {
	// Query Log
	src, err := boolQuery.Source()
	if err != nil {
		golog.Error(err)
		return true
	}

	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		golog.Error(err)
		return true
	}
	golog.Print(fmt.Sprintf("\n{\"query\":\n%v}", string(data)), "\n")

	return false
}

func Index(tipe string) string {
	var index string
	switch tipe {
	case "location":
		index = "catalog-politica-location"
	case "event":
		index = "catalog-politica-event"
	case "person":
		index = "catalog-person-politica"
	}
	return index
}
