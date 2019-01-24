package main

import (
	rf "review-rest/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/api/v1/approve", rf.GetDeploymentList)
	r.GET("/api/v1/commitid", rf.GetCommitId)
	r.GET("/api/v1/reviewid", rf.GetReviewId)
	r.POST("/api/v1/review/summary", rf.PostSummary)
	r.Run(":8008")

}
