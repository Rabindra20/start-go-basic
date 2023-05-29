package main
import ("fmt"
	// "log"
	"net/http"
  "github.com/gin-gonic/gin"
//   "component"
)
func GetMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "GetMethod Called With Param: " + IdValue
	c.JSON(http.StatusOK, message)
  ReqPayload := make([]byte, 1024)
	ReqPayload, err := c.GetRawData()
	if err != nil {
		  fmt.Println(err)
		  return
	}
	fmt.Println("Request Payload Data: ", string(ReqPayload))
  }


func main() {
	router := gin.Default()
	// http://localhost:4000/api/v1/PersonId/5000
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
		"admin": "adminpass",
	  }))
	subRouterAuthenticated.GET("/:IdValue", GetMethod)
	listenPort := "4000"
	  
	  router.Run(":"+listenPort)
	}