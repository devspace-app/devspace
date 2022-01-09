package main

import (
	"context"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func GetClientSet() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()

	if err == rest.ErrNotInCluster {
		log.Println("Not running in cluster, attempting kubeconfig authentication")
	} else if err != nil {
		return nil, err
	}

	// create the kubeconfig instead
	if err == rest.ErrNotInCluster {
		home := homedir.HomeDir()
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))

		if err != nil {
			return nil, err
		}
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func main() {
	clientset, err := GetClientSet()

	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to Kubernetes cluster")

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		log.Println(pod.Labels)
	}
}
