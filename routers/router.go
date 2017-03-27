package routers

import (
	"net/http"
	"userApi/controllers"
	//"github.com/gorilla/mux"
	"github.com/drone/routes"
)

func init() {
	http.HandleFunc("/users", controllers.UserHandle)
	//http.HandleFunc("/users/{uid}/relationships", controllers.RelationGet)

	//r := mux.NewRouter()
	//r.HandleFunc()

	mux := routes.New()
	mux.Get("/users/:user_id/relationships", controllers.RelationGet)
	mux.Put("/users/:user_id/relationships/:other_user_id", controllers.RelationPut)
	http.Handle("/", mux)
}
