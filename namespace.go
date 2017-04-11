package cvkube

import (
	"fmt"
	"encoding/json"
	"os"

	client "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/api"
)

func GetNamespaces(myClient *client.Clientset){
	namespaces, err := myClient.Namespaces().List(metav1.ListOptions{})
  if err != nil {
    fmt.Printf("Couldn't get all the namespaces %v\n", err)
    os.Exit(1)
  }
	renderNamespaces(namespaces)
}

func renderNamespaces(namespaces *api.NamespaceList) {
	for count, namespace := range namespaces.Items {
		fmt.Printf("NAMESPACE %d\n", count+1)
		pp, _ := json.MarshalIndent(&namespace, "", "   ")
		fmt.Println(string(pp))
	}
}
