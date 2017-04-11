package cvkube

import(
  "fmt"
  "os"

  client "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
  "k8s.io/client-go/tools/clientcmd"
)

func GetClient(kubeconfig string)(*client.Clientset) {
  // uses the current context in kubeconfig
  config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
  if err != nil {
    fmt.Printf("Error opening Kubeconfig: %v\n", err)
    os.Exit(1)
  }

  myClient, err := client.NewForConfig(config)
    if err != nil {
      fmt.Printf("Error creating REST Kube Client: %v\n", err)
      os.Exit(1)
  }

  return myClient
}
