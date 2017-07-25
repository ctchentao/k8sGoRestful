package main

import(
	//	"fmt"
//	"github.com/emicklei/go-restful"
//	"k8sGoRestful"
	"net/http"
	"log"
	"k8sGoRestful/handlers"
)

func main(){
	handler := handlers.New()

	server := &http.Server{Addr: "172.18.191.5:8080", Handler: handler}
	log.Fatal(server.ListenAndServe())

	//go-restful
	/*
	wsContainer := restful.NewContainer()

	k8sGoRestful.CreateRoutes(wsContainer)
	server := &http.Server{Addr: "172.18.191.5:8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
	*/
}



