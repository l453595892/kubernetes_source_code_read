// Example showing how to patch Kubernetes resources.
package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	//  Leave blank for the default context in your kube config.
	context = ""
)

func main() {
	//  Get the local kube config.
	fmt.Printf("Connecting to Kubernetes Context %v\n", context)
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: context}).ClientConfig()
	if err != nil {
		panic(err.Error())
	}

	// Creates the dynamic interface.
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//  Create a GVR which represents an Istio Virtual Service.
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1alpha3",
		Resource: "virtualservices",
	}

	//  List all of the Virtual Services.
	virtualServices, err := dynamicClient.Resource(virtualServiceGVR).Namespace("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for index, virtualService := range virtualServices.Items {
		fmt.Printf("VirtualService %d: %s\n", index+1, virtualService.GetName())
	}
}
