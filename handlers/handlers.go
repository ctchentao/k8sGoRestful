package handlers

import (
//	"io/ioutil"
	"net/http"
	"github.com/tedsuo/rata"
	k8s "k8sGoRestful"
)

func New() http.Handler {
	podsHandler := NewPodsHandler()

	actions := rata.Handlers{
		//pods
		k8s.GetPods:    route(podsHandler.GetPods),
		k8s.CreatePods: route(podsHandler.CreatePods),
		k8s.DeletePods: route(podsHandler.DeletePods),
	}

	handler, err := rata.NewRouter(k8s.Routes,actions)
	if err != nil {
		panic("unable to create router: " + err.Error())
	}

	return handler
}

func route(f http.HandlerFunc) http.Handler {
	return f
}