package main

import (
	"flag"
	"encoding/json"
	"fmt"
  "reflect"
	"os"

	"github.com/docopt/docopt-go"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	// api "k8s.io/client-go/pkg/api/v1"
  // "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	client "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
)

var (
	kubeconfig = flag.String("kubeconfig", "/Users/jasonrouault/.kube/config", "absolute path to the kubeconfig file")
)

func main() {
	usage := `

		Usage:
	  controller.go [--config <config>][--namespace <namespace>] get-all (policies|pods)
	  controller.go [--config <config>][--namespace <namespace>] get-pods <policy>
	  controller.go [--config <config>][--namespace <namespace>] get-policies <pod>
	  controller.go [--config <config>][--namespace <namespace>] get-rules <pod> [human]

		Options:
		--namespace=NAMESPACE Namespace to run the query in
		--config=FILE path to the KubeConfig file.
		`

	arguments, _ := docopt.Parse(usage, nil, true, "cvkube", false)
	//fmt.Println(arguments)
	// Get location of the Kubeconfig file. By default in your home.
	var kubeconfig string
	if arguments["--config"] == nil {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	} else {
		kubeconfig = arguments["--config"].(string)
	}

	// Get namespace, by default it will be "default"
	var namespace string
	if arguments["--namespace"] == nil {
		namespace = "default"
	} else {
		namespace = arguments["--namespace"].(string)
	}

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Error opening Kubeconfig: %v\n", err)
		os.Exit(1)
	}

  /*
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
  */


	myClient, err := client.NewForConfig(config)
		if err != nil {
			fmt.Printf("Error creating REST Kube Client: %v\n", err)
			os.Exit(1)
	}

	// Display all policies. Similar to kubectl describe policies in json
	if arguments["get-all"].(bool) && arguments["policies"].(bool) {

		policies, err := myClient.Extensions().NetworkPolicies(namespace).List(metav1.ListOptions{})
    //polices, err := clientset.ExtensionsV1beta1().PodSecurityPolicy
		if err != nil {
			fmt.Printf("Couldn't get Network Policy: %v\n", err)
			os.Exit(1)
		}
		//renderPolicies(policies)
    fmt.Println("now is a type of: ", reflect.TypeOf(policies))
    for count, policy := range policies.Items {
  		fmt.Printf("POLICY %d\n", count+1)
  		pp, _ := json.MarshalIndent(&policy, "", "   ")
  		fmt.Println(string(pp))
  	}
		os.Exit(0)
	}


// Display all pods. Similar to kubectl describe pods in json
	if arguments["get-all"].(bool) && arguments["pods"].(bool) {
    /*
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Couldn't get all the pods %v\n", err)
			os.Exit(1)
		}

		renderPods(pods)
    fmt.Println("now is a type of: ", reflect.TypeOf(pods))
    */
    getPods()
		os.Exit(0)
	}
}

/*
func renderPolicies(policies *extensions.NetworkPolicyList) {
	for count, policy := range policies.Items {
		fmt.Printf("POLICY %d\n", count+1)
		pp, _ := json.MarshalIndent(&policy, "", "   ")
		fmt.Println(string(pp))
	}
}
*/
/*
func renderPods(pods *api.PodList) {
	for count, pod := range pods.Items {
		fmt.Printf("POD %d\n", count+1)
		pp, _ := json.MarshalIndent(&pod, "", "   ")
		fmt.Println(string(pp))
	}
}
*/
