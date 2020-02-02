package main

import (
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/EquallerReflect"
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/HasherMd5"
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/JsonMarshaller"
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/WebServerStd"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/EntityEqualRepository/EntityEqualRepository"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/EntityEqualRepository/EquallerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HashRepository"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HasherInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/BaseHandler"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebServerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WebService"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebService/WerServiceInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/HashRepositoryInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractor"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/EntityEqualRepositoryInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractor"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
	"log"
)

// TODO: move mock-errors to dedicated module and incapsulate it
// TODO: comment every package and type/func/method
// TODO: Add logger modules
// TODO: Add Error modules

func main() {
	var (
		jsonArrayToObjectConvertInteractor JsonArrayToObjectConvertInteractorInterface.Interactor
		jsonCompareInteractor              JsonCompareInteractorInterface.Interactor
		webServer                          WebServerInterface.WebServer
		webService                         WerServiceInterface.WebService
		hashRepository                     HashRepositoryInterface.Repository
		entityEqualRepository              EntityEqualRepositoryInterface.Repository
		hasher                             HasherInterface.Hasher
		equaller                           EquallerInterface.Equaller
	)

	jsonMarshaller := JsonMarshaller.New()
	hasher = HasherMd5.New()

	hashRepository = HashRepository.New().
		SetMarshaller(jsonMarshaller).
		SetHasher(hasher)

	jsonArrayToObjectConvertInteractor = JsonArrayToObjectConvertInteractor.New().
		SetHashRepository(hashRepository)

	equaller = EquallerReflect.New()
	entityEqualRepository = EntityEqualRepository.New().
		SetEqualler(equaller)

	jsonCompareInteractor = JsonCompareInteractor.New().
		SetArrayToObjectConverter(jsonArrayToObjectConvertInteractor).
		SetEntityEqualRepository(entityEqualRepository)

	webServer = WebServerStd.New()

	webService = WebService.New().
		SetWebServer(webServer).
		SetMarshaller(jsonMarshaller)

	webHandler := BaseHandler.New().
		SetJsonCompareInteractor(jsonCompareInteractor)

	webService.HandleFunc("/", webHandler)

	err := webServer.Start()
	if err != nil {
		log.Fatalf("start web server failed: %s", err.Error())
		return
	}

}
