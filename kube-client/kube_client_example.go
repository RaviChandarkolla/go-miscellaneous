package main

import (
	"context"
	"fmt"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	_ "os"
	"path/filepath"
	_ "strings"
)

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//feature := "system-feature"

	//labelJob := fmt.Sprintf("app.kubernetes.io/name=%s", feature)
	//podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{LabelSelector: labelJob})
	podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		println(err.Error())
	}
	println(podList)
	println(podList.Items[0].Name)
	//podName := podList.Items[0].Name

	//podName := "mbiq-fluentd-0"
	namespace := "default"

	for _, podItem := range podList.Items {
		podName := podItem.ObjectMeta.Name
		println(podName)
		for _, container := range podItem.Spec.Containers {
			containerName := container.Name
			println(containerName)
			logs, err := getPodLogs(clientset, podName, namespace, containerName)
			if err != nil {
				fmt.Printf("Error getting logs: %v\n", err)
				return
			}

			fmt.Printf("Logs for pod %s in namespace %s:\n%s\n", podName, namespace, logs)
			fmt.Println("======================================")
		}

	}

}

func getPodLogs(clientset *kubernetes.Clientset, podName, namespace, containerName string) (string, error) {
	// Create a PodLogsOptions object

	noLines := int64(5)
	opts := &corev1.PodLogOptions{
		Container: containerName, // Specify container name if not the first container
		//Container: "fluentd",
		TailLines: &noLines, // Specify number of lines to retrieve, optional
	}

	// Get the logs for the specified pod
	podLogReq := clientset.CoreV1().Pods(namespace).GetLogs(podName, opts)
	podLogs, err := podLogReq.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	// Read the log lines
	logs, err := ioutil.ReadAll(podLogs)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}
