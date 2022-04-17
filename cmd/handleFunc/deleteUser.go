package handleFunc

import (
	"authDB/pkg/application/services"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := services.DeleteUser(id)
	response.Send(w, "DeleteUser")

}
