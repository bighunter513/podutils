package main

import (
	"context"
	"fmt"
	"net"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetPodIp get pod's ip from k8s pod container
func GetPodIp() string {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("err occurred:", err)
		return GetIp()
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("err2 occurred:", err)
		panic(err.Error())
	}

	pod, err := clientset.CoreV1().Pods("default").Get(
		context.TODO(),
		os.Getenv("HOSTNAME"),
		metav1.GetOptions{})
	if err != nil {
		fmt.Println("err3 occurred:", err)
		return GetIp()
	}

	return pod.Status.PodIP
}

func main() {
	ip := GetPodIp()

	fmt.Println("POD HOSTNAME:", os.Getenv("HOSTNAME"))
	fmt.Println("POD ip:", ip)
}
