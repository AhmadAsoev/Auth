package handleFunc

import (
	"authDB/pkg/application/services"
	"authDB/pkg/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var request domain.User
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Cant decode into json")); err != nil {
			log.Println("handleFunc/AddUser/Decode/Write")
			return
		}
		log.Println("handleFunc/AddUser/Decode")
		return
	}

	response := services.UpdateUser(request, id)
	response.Send(w, "UpdateUser")
}
