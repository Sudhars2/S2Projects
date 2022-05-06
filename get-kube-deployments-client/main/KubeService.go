package main

import (
	"context"
	//"flag"
	"fmt"
	//"path/filepath"


	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/rest"
	"reflect"
	"strconv"

)

type KubeServices struct {
	Name string
    ApplicationGroup string
    RunningPodsCount string

}


func GetKubeServices(ServVal string,newconfig *rest.Config) []KubeServices{
	
	config := newconfig 
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//To find the methods supported by clientset
	// fooType := reflect.TypeOf(clientset.AppsV1().Deployments("default"))
	// for i := 0; i < fooType.NumMethod(); i++ {
    // method := fooType.Method(i)
    // fmt.Println(method.Name)
	var Servicelist []KubeServices
	for {
		
		 deps, erro := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
		 if erro != nil {
			panic(erro.Error())
		}
		 fooType := reflect.TypeOf(deps)
		 fmt.Println(fooType)
	if ServVal == "services" {	
	for _, i := range deps.Items {
		Servicelist = append(Servicelist,KubeServices{i.ObjectMeta.Name,i.ObjectMeta.Labels["applicationGroup"],strconv.Itoa(int(i.Status.Replicas))})
		// fmt.Println(i.ObjectMeta.Name)
		// fmt.Println(i.ObjectMeta.Labels)
		// fmt.Println(int(i.Status.Replicas))
	}
} else { 
	for _, j := range deps.Items {
		if j.ObjectMeta.Labels["applicationGroup"] == ServVal {
		Servicelist = append(Servicelist,KubeServices{j.ObjectMeta.Name,j.ObjectMeta.Labels["applicationGroup"],strconv.Itoa(int(j.Status.Replicas))})
		
			}
		}
	}

	return Servicelist
		
	}
}