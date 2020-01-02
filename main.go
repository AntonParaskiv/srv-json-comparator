package main

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractor"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractor"
	"github.com/gin-gonic/gin"
)

func main() {
	jsonArrayToObjectConvertInteractor := JsonArrayToObjectConvertInteractor.New()
	jsonCompareInteractor := JsonCompareInteractor.New().SetArrayToObjectConverter(jsonArrayToObjectConvertInteractor)
	webService := WebService.New().SetJsonCompareInteractor(jsonCompareInteractor)

	r := gin.Default()
	r.POST("/", webService.BaseHandler)
	// TODO: handle error
	r.Run() // listen and serve on 0.0.0.0:8080
}
