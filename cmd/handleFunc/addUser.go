package handleFunc

import (
	"authDB/pkg/application/services"
	"authDB/pkg/domain"
	"encoding/json"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request domain.User

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Cant decode into json")); err != nil {
			log.Println("handleFunc/AddUser/Decode/Write")
			return
		}
		log.Println("handleFunc/AddUser/Decode")
		return
	}

	response := services.AddUser(request)
	response.Send(w, "AddUser")
}
