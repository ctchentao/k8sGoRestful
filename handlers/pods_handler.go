package handlers

import(
	"flag"
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/tools/clientcmd"
	"log"
//	"fmt"
//	"github.com/emicklei/go-restful"
//	"github.com/gorilla/schema"
	"net/http"
//	"io"
	"k8sGoRestful/models"
	"io/ioutil"
	"encoding/json"
	"github.com/tedsuo/rata"
)

var (
//	decoder *schema.Decoder
	kubeconfig = flag.String("kubeconfig", "/home/chentao/config.yaml", "absolute path to the kubeconfig file")
)

type PodsHandler struct {

}
func NewPodsHandler() *PodsHandler {
	return &PodsHandler{}
}

func (h *PodsHandler)k8sCreatePods(pods *models.Pods, namespace string) error{
	var e error
//	fmt.Println(namespace)
	flag.Parse()
	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		e = err
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		e = err
		panic(err)
	}

	pod := new(v1.Pod)
	pod.TypeMeta = unversioned.TypeMeta{Kind: "Pod", APIVersion: "v1"}
	labelsMap := make(map[string]string)
	labels := pods.GetLabel()
	for i :=0;i<len(labels);i++{
		labelsMap[*(labels[i].Name)] = *(labels[i].Value)
	}
	pod.ObjectMeta = v1.ObjectMeta{Name: pods.GetName(), Namespace: namespace, Labels: labelsMap}
	pod.Spec = v1.PodSpec{
		RestartPolicy: v1.RestartPolicyAlways,
		Containers: []v1.Container{
			v1.Container{
				Name:  pods.GetName(),
				Image: pods.GetImage(),
				Ports: []v1.ContainerPort{
					v1.ContainerPort{
						ContainerPort: pods.GetContainPort(),
						Protocol:      v1.ProtocolTCP,
					},
				},
			},
		},
	}
	podname, err := clientset.Core().Pods(namespace).Create(pod)
	log.Printf("pod %s have cretae\n", podname.ObjectMeta.Name)
	return e
}
/*
func CreatePods1(req *restful.Request, resp *restful.Response){
	err := req.Request.ParseForm()
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	pods := new(models.Pods)
	data, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := json.Unmarshal(data,&pods); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	namespace := req.PathParameter("namespace")
	err = k8sCreatePods(pods,namespace)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
	}
	io.WriteString(resp.ResponseWriter, "<html><body>create pods successfully!</body></html>")
}*/

func (h *PodsHandler)CreatePods(w http.ResponseWriter, req *http.Request){
	data, err := ioutil.ReadAll(req.Body)
	response := &models.PodsResponse{}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pods := new(models.Pods)
	if err := json.Unmarshal(data,&pods); err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	namespace := rata.Param(req, "namespace")
	err = h.k8sCreatePods(pods,namespace)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeResponse(w, response)
}

func (h *PodsHandler)GetPods(w http.ResponseWriter, req *http.Request){

}

func (h *PodsHandler)DeletePods(w http.ResponseWriter, req *http.Request){

}
