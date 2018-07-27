package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PodEnv sets up the pod environment
type PodEnv struct {
	PodIP        string `json:"pod_ip,omitempty"`
	PodName      string `json:"pod_name,omitempty"`
	NodeName     string `json:"node_name,omitempty"`
	PodNamespace string `json:"pod_namespace,omitempty"`
}

func reflect(rw http.ResponseWriter, req *http.Request) {
	pEnv := PodEnv{
		PodIP:        os.Getenv("POD_IP"),
		PodName:      os.Getenv("POD_NAME"),
		NodeName:     os.Getenv("NODE_NAME"),
		PodNamespace: os.Getenv("POD_NAMESPACE"),
	}
	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	_ = encoder.Encode(pEnv)
}

func main() {
	router := mux.NewRouter()
	http.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", reflect)
	http.ListenAndServe(":3000", router)
}
