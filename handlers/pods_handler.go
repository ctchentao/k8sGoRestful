package handlers

import(
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	"log"
	"net/http"
	"k8sGoRestful/models"
	"io/ioutil"
	"github.com/tedsuo/rata"
	"github.com/gogo/protobuf/proto"
)

type PodsHandler struct {

}
func NewPodsHandler() *PodsHandler {
	return &PodsHandler{}
}

func (h *PodsHandler)k8sCreatePods(pods *models.Pods, namespace string) error{
	clientset := newClientset()
	pod := new(v1.Pod)
	pod.TypeMeta = unversioned.TypeMeta{Kind: "Pod", APIVersion: "v1"}
	labelsMap := convertMap(pods.GetLabel())

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
	podInstance, err := clientset.Core().Pods(namespace).Create(pod)
	if err == nil {
		log.Printf("pod %s have create\n", podInstance.ObjectMeta.Name)
	}
	return err
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
	pods := &models.Pods{}
	if err := proto.Unmarshal(data,pods); err !=nil {
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
	//todo
}

func (h *PodsHandler)DeletePods(w http.ResponseWriter, req *http.Request){
	//todo
}
