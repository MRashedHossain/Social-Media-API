package accounts

import (
	"fmt"
	"user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/", ServerRunning)
	router.GET("/user", user.GetUserList)
}
func ServerRunning() {
	fmt.Println("Welcome!!!")
	fmt.Println("The Server is running!!!")
}
