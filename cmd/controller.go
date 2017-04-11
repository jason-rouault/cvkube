package main

import (
	"flag"
	"os"

	"github.com/jason-rouault/cvkube"
	"github.com/docopt/docopt-go"
)

var (
	kubeconfig = flag.String("kubeconfig", "/Users/jasonrouault/.kube/config", "absolute path to the kubeconfig file")
)

func main() {
	usage := `

		Usage:
	  controller.go [--config <config>][--namespace <namespace>] get-all (policies|pods|namespaces)
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

	myClient := cvkube.GetClient(kubeconfig)

	// Display all policies. Similar to kubectl describe policies in json
	if arguments["get-all"].(bool) && arguments["policies"].(bool) {
		cvkube.GetNetworkPolicies(myClient, namespace)
		os.Exit(0)
	}


// Display all pods. Similar to kubectl describe pods in json
	if arguments["get-all"].(bool) && arguments["pods"].(bool){
    cvkube.GetPods(myClient, namespace)
		os.Exit(0)
	}

// Display all namespaces. Similar to kubectl describe namespaces in json
	if arguments["get-all"].(bool) && arguments["namespaces"].(bool){
    cvkube.GetNamespaces(myClient)
		os.Exit(0)
	}

	// Get all the policies that get applied to a Pod.
	if arguments["get-policies"].(bool) {
		cvkube.GetPodNetworkPolicies(myClient, namespace, arguments["<pod>"].(string))
		os.Exit(0)
	}

	// Get all the IngressRules that get applied to a Pod.
	if arguments["get-rules"].(bool) {
		cvkube.GetPodIngressRules(myClient, namespace, arguments["<pod>"].(string), arguments["human"].(bool))
		os.Exit(0)
	}
}
