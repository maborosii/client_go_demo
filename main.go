package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.APIPath = "/api"
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	pod := v1.Pod{}
	restClient.Get().Namespace("test").Resource("pods").Name("test-nginx-deployment-585449566-4542q").Do(context.TODO()).Into(&pod)
	println(pod.Name)

}
