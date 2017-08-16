package main

import (
//	"k8sGoRestful"
	"k8sGoRestful/handlers"
//	"k8sGoRestful/models"

//	"fmt"


	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/grouper"
	"os"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/sigmon"
)

func main(){
	handler := handlers.New()
//	k8sServer := http.Server{Addr: "172.27.35.4:8080", Handler: handler}
//	log.Fatal(k8sServer.ListenAndServe())
	server := http_server.New("172.27.35.4:8080", handler)
	members := grouper.Members{
		{"server", server},
	}
	group := grouper.NewParallel(os.Interrupt, members)
	monitor := ifrit.Invoke(sigmon.New(group))



	err := <-monitor.Wait()
	if err != nil {

		os.Exit(1)
	}
/*
	client := k8sGoRestful.NewClient("http://172.27.35.4:8080")
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
/*

	server := &http.Server{Addr: "127.0.0.1:8080", Handler: handler}
	log.Fatal(server.ListenAndServe())
*/
	//go-restful
	/*
	wsContainer := restful.NewContainer()

	k8sGoRestful.CreateRoutes(wsContainer)
	server := &http.Server{Addr: "172.18.191.5:8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
	*/
}



