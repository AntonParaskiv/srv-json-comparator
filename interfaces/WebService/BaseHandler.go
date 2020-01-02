package WebService

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/gin-gonic/gin"
)

// TODO: move to rpc2
type Request struct {
	Method       string `json:"method"`
	MethodParams struct {
		First  interface{} `json:"first"`
		Second interface{} `json:"second"`
	} `json:"methodParams"`
}

func (ws *WebService) BaseHandler(c *gin.Context) {
	var request Request
	// TODO: handle error
	c.Bind(&request)

	first := JsonEntity.NewFromInterface(request.MethodParams.First)
	second := JsonEntity.NewFromInterface(request.MethodParams.Second)

	if request.Method == "isEqual" {
		isEqual, _ := ws.jsonCompareInteractor.IsEqual(first, second)

		c.JSON(200, gin.H{
			"isEqual": isEqual,
		})
	}

}
