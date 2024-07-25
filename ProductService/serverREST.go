package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": []string{"Book", "Pen"}})
	})
	router.Run(":8080")
}
