package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// responds with the list of all albums as JSON.
func getSample(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sampleData)
}

// set api route(should be seprate in api folder)
func Get_api() {

	router := gin.Default()
	router.GET("/sampleData", getSample)

	router.Run("localhost:8080")
}
