package main

import "net/http"

import "fmt"

import "os"

import "log"

var podIP string
var nodeName string

const podOPEnvVar = "POD_IP"
const nodeNameEnvVar = "NODE_NAME"

func init() {
	podIP = os.Getenv(podOPEnvVar)
	if podIP == "" {
		log.Fatalf("Missing %s env variable", podOPEnvVar)
	}

	nodeName = os.Getenv(nodeNameEnvVar)
	if nodeName == "" {
		log.Fatalf("Missing %s env variable", nodeNameEnvVar)
	}

	log.Printf("Pod IP %s on Node %s\n", podIP, nodeName)

}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello from Pod IP %s on Node %s", podIP, nodeName)
	})

	log.Println("started app...")
	http.ListenAndServe(":8080", nil)
}
