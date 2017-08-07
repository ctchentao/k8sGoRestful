package handlers

import (
//	"io/ioutil"
	"net/http"
	"github.com/tedsuo/rata"
	k8s "k8sGoRestful"
	"github.com/gogo/protobuf/proto"
	"strconv"
	"flag"
	"k8s.io/client-go/1.5/tools/clientcmd"
	"k8s.io/client-go/1.5/kubernetes"
	"k8sGoRestful/models"
)

var (
	kubeconfig = flag.String("kubeconfig", "/home/chentao/config.yaml", "absolute path to the kubeconfig file")

)

func New() http.Handler {
	podsHandler := NewPodsHandler()
	jobsHandler := NewJobsHandler()

	actions := rata.Handlers{
		//pods
		k8s.GetPods:    route(podsHandler.GetPods),
		k8s.CreatePods: route(podsHandler.CreatePods),
		k8s.DeletePods: route(podsHandler.DeletePods),

		//jobs
		k8s.CreateJobs: route(jobsHandler.CreateJobs),
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

func writeResponse(w http.ResponseWriter, message proto.Message) {
	responseBytes, err := proto.Marshal(message)
	if err != nil {
		panic("Unable to encode Proto: " + err.Error())
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(responseBytes)))
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.WriteHeader(http.StatusOK)

	w.Write(responseBytes)
}

func newClientset() (*kubernetes.Clientset) {
	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func convertMap(labels []models.Label) map[string]string {
	labelsMap := make(map[string]string)
	for i :=0;i<len(labels);i++{
		labelsMap[labels[i].Name] = labels[i].Value
	}
	return labelsMap
}