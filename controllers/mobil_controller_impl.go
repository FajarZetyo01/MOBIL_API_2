package controllers

import (
	"MOBIL_API_2/helpers"
	"MOBIL_API_2/model/web"
	"MOBIL_API_2/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MobilControllerImpl struct {
	MobilService services.MobilService
}

func NewMobilController(mobilService services.MobilService) *MobilControllerImpl {
	return &MobilControllerImpl{
		MobilService: mobilService,
	}
}

func (controller *MobilControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mobilCreateRequest := web.MobilCreateRequest{}
	helpers.ReadFromRequestBody(request, &mobilCreateRequest)
	mobilResponse := controller.MobilService.Create(request.Context(), mobilCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mobilResponse,
	}
	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *MobilControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mobilUpdateRequest := web.MobilUpdateRequest{}
	helpers.ReadFromRequestBody(request, &mobilUpdateRequest)
	mobilId := params.ByName("mobilId")
	id, err := strconv.Atoi(mobilId)
	helpers.PanicIfError(err)
	mobilUpdateRequest.Id = id
	mobilResponse := controller.MobilService.Update(request.Context(), mobilUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mobilResponse,
	}
	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *MobilControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mobilId := params.ByName("mobilId")
	id, err := strconv.Atoi(mobilId)
	helpers.PanicIfError(err)
	controller.MobilService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller *MobilControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mobilId := params.ByName("mobilId")
	id, err := strconv.Atoi(mobilId)
	helpers.PanicIfError(err)
	mobilResponse := controller.MobilService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mobilResponse,
	}
	helpers.WriteToResponseBody(writer, webResponse)
}

func (controller MobilControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mobilResponses := controller.MobilService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mobilResponses,
	}
	helpers.WriteToResponseBody(writer, webResponse)
}
