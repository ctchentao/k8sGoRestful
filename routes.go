package k8sGoRestful

import(
//	"github.com/emicklei/go-restful"
//	"k8sGoRestful/handlers"
	"github.com/tedsuo/rata"
)

const (
	// pods
	CreatePods = "CreatePods"
	GetPods    = "GetPods"
	DeletePods = "DeletesPods"
)

var Routes = rata.Routes{
	//Pods
	{Path: "/api/v1/namespaces/:namespace/pods/:name", Method: "GET", Name: GetPods},
	{Path: "/api/v1/namespaces/:namespace/pods", Method: "POST", Name: CreatePods},
	{Path: "/api/v1/namespaces/:namespace/pods/:name", Method: "DELETE", Name: DeletePods},
}

/*
func CreateRoutes(container *restful.Container){
	ws := new(restful.WebService)
	ws.Path("/api/v1/namespaces")
	ws.Route(ws.GET("/{namespace}/pods/{name}").To(handlers.GetPods)).
		Param(ws.PathParameter("namespace","identifier of the user").DataType("string")).
		Param(ws.PathParameter("name","identifier of the user").DataType("string"))
	ws.Route(ws.POST("/{namespace}/pods").To(handlers.CreatePods)).
		Param(ws.PathParameter("namespace","identifier of the user").DataType("string"))
	ws.Route(ws.DELETE("/{namespace}/pods/{name}").To(handlers.DeletePods)).
		Param(ws.PathParameter("namespace","identifier of the user").DataType("string")).
		Param(ws.PathParameter("name","identifier of the user").DataType("string"))
	container.Add(ws)

}
*/