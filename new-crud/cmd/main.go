package main

import (
	"crud/environment"
	"crud/pkg/configure"
	routes "crud/pkg/crud_routes"
	"crud/pkg/models"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	en, _ := environment.Getenv()
	configure.Connection(en)
}

func main() {
	r := mux.NewRouter()
	routes.Userroutes(r)
	go ticker()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func ticker() {
	en, _ := environment.Getenv()
	configure.Connection(en)
	repo := models.Newdbcontroller(configure.GetDB())
	cd := time.Tick(time.Second*1)

	for _ = range cd {
		repo.Perm_del()
	}

}
