package router


import "github.com/gin-gonic/gin"


func Initialize(){
		router := gin.Default()
	router.GET("/v1/api", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run() // listen and serve on 0.0.0.0:8080
}