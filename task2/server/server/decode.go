package server

import (
	"encoding/base64"
	"encoding/json"
	"mime"
	"net/http"
)

func (server *Server) DecodeHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		type DecodeRequest struct {
			InputString string `json:"inputString"`
		}
		type DecodeResponse struct {
			OutputString string `json:"outputString"`
		}

		contentType := request.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
		if mediaType != "application/json" {
			http.Error(response, "expect application/json Content-Type",
				http.StatusUnsupportedMediaType)
			return
		}

		decoder := json.NewDecoder(request.Body)
		decoder.DisallowUnknownFields()
		var requestString DecodeRequest
		if err := decoder.Decode(&requestString); err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		responseString, err := base64.StdEncoding.DecodeString(requestString.InputString)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}
		server.Answer(response, &DecodeResponse{string(responseString)})

	} else {
		http.Error(response, "Only POST Method of decode is allowed",
			http.StatusMethodNotAllowed)
	}
}
