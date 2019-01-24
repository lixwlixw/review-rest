package restfulapi

import (
	myquery "review-rest/database/query"
	sq "review-rest/structerr"
	"net/http"
	"github.com/gin-gonic/gin"
)
type Ns struct {
	Ns string `form:"commit" binding:"required"`
        REID string `form:"id" binding:"required"`
}
type Svcs struct {
 Services   []Svc `json:"reviewid"`
}

type Svc struct {
 Name string `json:"id"`
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

func GetReviewId(c *gin.Context) {
 services := Svcs{}
 values := myquery.GetReviewId()
 for i:=0;i<len(values);i++{
svc := Svc{Name:values[i]}
  services.Services = append(services.Services,svc)
 }

 c.JSON(200, services)
 return
}

func PostSummary(c *gin.Context) {
        var Na Ns
        if bindErr := c.Bind(&Na); bindErr != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
                })
                return
        }
        c.JSON(http.StatusOK, gin.H{
                "update ok": myquery.PostSummary(Na.Ns, Na.REID),
        })
}
