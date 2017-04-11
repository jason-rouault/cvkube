package cvkube

import (
	"fmt"
	"encoding/json"
	"os"

	client "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/api"
)

// Get all pods for a given namespace
func GetPods(myClient *client.Clientset, namespace string){
	pods, err := myClient.Pods(namespace).List(metav1.ListOptions{})
  if err != nil {
    fmt.Printf("Couldn't get all the pods %v\n", err)
    os.Exit(1)
  }
	renderPods(pods)
}

func renderPods(pods *api.PodList) {
	for count, pod := range pods.Items {
		fmt.Printf("POD %d\n", count+1)
		pp, _ := json.MarshalIndent(&pod, "", "   ")
		fmt.Println(string(pp))
	}
}
