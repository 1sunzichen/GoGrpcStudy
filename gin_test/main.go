package main

import "github.com/gin-gonic/gin"



func pong(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"pong",
	})
}
func main(){
	r:=gin.Default();
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	v1:=r.Group("/v1")
	{
		v1.POST("/login",pong)
	}
	r.Run()//listen and serve  on 0.0.0.0:8080
}
