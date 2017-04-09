package main

import (
	"fmt"
)

func getPods(){
  /*
  pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
  if err != nil {
    fmt.Printf("Couldn't get all the pods %v\n", err)
    os.Exit(1)
  }

  renderPods(pods)
  */
  fmt.Println("now is a type of: ", "test")
}
/*
func renderPods(pods *api.PodList) {
	for count, pod := range pods.Items {
		fmt.Printf("POD %d\n", count+1)
		pp, _ := json.MarshalIndent(&pod, "", "   ")
		fmt.Println(string(pp))
	}
}
*/
