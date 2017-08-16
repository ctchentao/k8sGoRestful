package handlers

import(
	"k8s.io/client-go/1.5/pkg/api/v1"
//	"k8s.io/client-go/1.5/pkg/api/resource"
	"log"
	"net/http"
	"k8sGoRestful/models"
	batchv1 "k8s.io/client-go/1.5/pkg/apis/batch/v1"
	"io/ioutil"
	"github.com/gogo/protobuf/proto"
//	"k8s.io/client-go/1.5/pkg/api"
//	"k8s.io/client-go/1.5/pkg/apis/batch"
)

type JobsHandler struct {

}

func NewJobsHandler() *JobsHandler {
	return &JobsHandler{}
}

func (h *JobsHandler) k8sCreateJobs(jobs *models.Jobs, namespace string) error{
	labelsMap := convertMap(jobs.GetLabel())
	job := &batchv1.Job{
		ObjectMeta: v1.ObjectMeta{Name: jobs.GetName(), Namespace: namespace, Labels: labelsMap},
		Spec: batchv1.JobSpec{
			Completions: jobs.GetCompletions(),
			Parallelism: jobs.GetParallelism(),

			Template: v1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name: jobs.GetName(),
					Namespace: namespace,
					Labels: labelsMap,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Image: jobs.GetImage(),
							Resources: v1.ResourceRequirements{
								Requests: v1.ResourceList{
								//	{"cpu": ""},
								//	{"memory": ""},
								},
							},
							VolumeMounts: []v1.VolumeMount{
								{
									MountPath: jobs.GetVolumes().GetMountPath(),
									Name: jobs.GetVolumes().GetName(),
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyOnFailure,
					Volumes: []v1.Volume{
						{
							Name: jobs.GetVolumes().GetName(),
							VolumeSource: v1.VolumeSource{
								EmptyDir: &v1.EmptyDirVolumeSource{},
							},
						},

					},
				},
			},
		},
	}
	clientset := newClientset()
	jobInstance, err := clientset.Batch().Jobs(namespace).Create(job)
	if err == nil {
		log.Printf("pod %s have create\n", jobInstance.ObjectMeta.Name)
	}
	return err
}

func (h *JobsHandler) CreateJobs(w http.ResponseWriter, req *http.Request){
	data, err := ioutil.ReadAll(req.Body)
	response := &models.PodsResponse{}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jobs := &models.Pods{}
	if err := proto.Unmarshal(data,jobs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeResponse(w,response)
}