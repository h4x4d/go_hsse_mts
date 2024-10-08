package server

import (
	"math/rand/v2"
	"net/http"
	"time"
)

func (server *Server) HardOpHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		operationTime := rand.IntN(10) + 10
		time.Sleep(time.Duration(operationTime) * time.Second)

		type HardOpResponse struct {
			Status    string `json:"status"`
			SleepTime int    `json:"sleepTime"`
		}

		if rand.IntN(2) == 1 {
			server.Answer(response, &HardOpResponse{"success",
				operationTime})
		} else {
			server.AnswerWithCode(response, &HardOpResponse{"fail",
				operationTime}, 500)
		}

	} else {
		http.Error(response, "Only GET Method of hard-op is allowed",
			http.StatusMethodNotAllowed)
	}
}
