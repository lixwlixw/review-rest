package main

import (
	"review-rest/getlogs"
	rf "review-rest/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/api/v1/approve", rf.GetNodeInfo)
	r.Run(":8008")

}
