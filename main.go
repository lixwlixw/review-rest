package main

import (
	rf "review-rest/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/api/v1/approve", rf.GetDeploymentList)
	r.Run(":8008")

}
