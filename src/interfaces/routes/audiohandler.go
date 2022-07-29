package routes

import (
	"audioTest/src/application"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

type AudioHandler struct {
	audioService application.IAudioService
}

func (handler *AudioHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Post("/upload", handler.Upload)
	r.Get("/info", handler.GetInfo)
	return r
}

func (handler *AudioHandler) Upload(writer http.ResponseWriter, request *http.Request) {
	fileBytes, err := ioutil.ReadAll(request.Body)
	params := request.URL.Query()
	fileName := params["name"][0]
	if fileName == "" {
		writer.WriteHeader(400)
		return
	}

	err = handler.audioService.StoreFile(fileBytes, fileName)
	if err != nil {
		writer.WriteHeader(400)
		return
	}
}

func (handler *AudioHandler) GetInfo(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	name, exists := request.Form["name"]
	if !exists {
		writer.WriteHeader(400)
		return
	}

	response, err := handler.audioService.GetInfo(name[0])
	if err != nil {
		writer.WriteHeader(400)
		return
	}

	if response == nil {
		writer.WriteHeader(404)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		fmt.Println(err)
		return
	}
}

func (handler *AudioHandler) GetList(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(400)
		return
	}
	duration, exists := request.Form["maxduration"]
	if !exists {
		writer.WriteHeader(400)
		return
	}

	response, err := handler.audioService.GetInfo(duration[0])
	if err != nil {
		writer.WriteHeader(400)
		return
	}

	if response == nil {
		writer.WriteHeader(404)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		fmt.Println(err)
		return
	}

}
