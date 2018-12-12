package restfulapi

import (
	myquery "review-rest/database/query"
	sq "review-rest/structerr"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Ns struct {
	Ns string `form:"commit" binding:"required"`
}
func GetDeploymentList(c *gin.Context) {

	var commit Ns
	if bindErr := c.Bind(&commit); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":myquery.GetDeployPoolList(commit.Ns),
	})
}
func GetCommitId(c *gin.Context) {

        var commit Ns
        if bindErr := c.Bind(&commit); bindErr != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
                })
                return
        }
        c.JSON(http.StatusOK, gin.H{
                "id":myquery.GetCommitId(commit.Ns),
        })
}
