package main

import (
    
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "fmt"
    "path/filepath"
    "k8s.io/client-go/util/homedir"
	"k8s.io/client-go/rest"
    "flag"
    "k8s.io/client-go/tools/clientcmd"
    
)

type config struct {

    config *rest.Config
}



func (conf *config) getConfig() *rest.Config {
    
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
        fmt.Println(home)
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
        
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
        fmt.Println(kubeconfig)
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, errone := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if errone != nil {
		panic(errone.Error())
	}
    conf.config = config
    return config
   
	
}


func (conf *config) GetAllService(w http.ResponseWriter, r *http.Request) {
    var newconfig *rest.Config
    if conf.config == nil {
    newconfig = conf.getConfig()
    conf.config = newconfig
    } else {
        newconfig = conf.config
    }
   services:=  GetKubeServices("services",newconfig)
   fmt.Println(services)
   for _,i := range services {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)       
   json.NewEncoder(w).Encode(i)
   }

}


func (conf *config) GetAppGroupSerice(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    name := vars["applicationGroup"]
    var newconfig *rest.Config
    if conf.config == nil {
    newconfig = conf.getConfig()
    conf.config = newconfig
    } else {
        newconfig = conf.config
    }
    appgrp := GetKubeServices(name,newconfig)
    fmt.Println(appgrp)
    for _,j := range appgrp {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)       
       json.NewEncoder(w).Encode(j)
       }
        
    
}
func main() {
    newconfig := config{nil}
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/services", newconfig.GetAllService) 
    r.HandleFunc("/services/{applicationGroup}", newconfig.GetAppGroupSerice) 
   
    //log.Fatal(http.ListenAndServe(":8 080 ", r))
    if err := http.ListenAndServe(":8080", r); err != nil {
        panic(err)
    }

}