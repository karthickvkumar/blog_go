package main

import (
	ctrl "./controllers"
	con "blogger/config"
	"github.com/gorilla/mux"
	"net/http"
)

var db = con.DBConn()

func main() {
	defer db.Close()
	router := mux.NewRouter()

	blog := router.Path("/").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.MainHandler)

	blog = router.Path("/about/{id}").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.AboutHandler)

	blog = router.Path("/about/delete/{id}").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.DeleteHandler)

	blog = router.Path("/create").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.CreateHandler)
	//blog.Methods("POST").HandlerFunc(ctrl.SaveHandler)

	blog = router.Path("/save").Subrouter()
	blog.Methods("POST").HandlerFunc(ctrl.SaveHandler)

	blog = router.Path("/edit/{id}").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.UpdateHandler)

	blog = router.Path("/update").Subrouter()
	blog.Methods("POST").HandlerFunc(ctrl.EditHandler)

	blog = router.Path("/index").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.IndexPageHandler)
	blog = router.Path("/internal").Subrouter()
	blog.Methods("GET").HandlerFunc(ctrl.InternalPageHandler)

	blog = router.Path("/login").Subrouter()
	blog.Methods("POST").HandlerFunc(ctrl.LoginHandler)

	blog = router.Path("/logout").Subrouter()
	blog.Methods("POST").HandlerFunc(ctrl.LogoutHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
