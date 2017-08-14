package main

import(
	//	"fmt"
//	"github.com/emicklei/go-restful"
//	"k8sGoRestful"
	"k8sGoRestful/handlers"

//	"k8sGoRestful/models"
//	"fmt"
//	"net/http/httptest"
	"net/http"
	"log"
)

func main(){
	handler := handlers.New()
/*	server := httptest.NewServer(handler)

	client := k8sGoRestful.NewClient(server.URL)
	fmt.Println(server.URL)
	pods := &models.Pods{
		Name: "testapi",
		Label: []models.Label{
			{Name: "name", Value: "testapi"},
			{Name: "company", Value: "IBM"},
			},
		Image: "nginx:1.7.9",
		ContainPort: 12345,
	}
	b, err := client.CreatePods("default", pods)
	if b {
		fmt.Println("create pods successfully!")
	} else {
		fmt.Println(err)
	}
*/

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



