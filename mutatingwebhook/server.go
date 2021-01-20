package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

//Server contains the functions handling server requests
type Server struct {
	ServerTLSConf *tls.Config
	ClientTLSConf *tls.Config
	CaPEM         []byte
}

func (s Server) getCA(w http.ResponseWriter, req *http.Request) {
	if len(s.CaPEM) == 0 {
		fmt.Fprintf(w, "No certificate found\n")
		return
	}

	// if base64 parameter is set, return in base64 format
	req.ParseForm()
	if _, hasParam := req.Form["base64"]; hasParam {
		fmt.Fprintf(w, string(base64.StdEncoding.EncodeToString(s.CaPEM)))
		return
	}

	fmt.Fprintf(w, string(s.CaPEM))
}

func (s Server) postWebhook(w http.ResponseWriter, r *http.Request) {
	var request AdmissionReviewRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON body in invalid format: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	if request.APIVersion != "admission.k8s.io/v1" || request.Kind != "AdmissionReview" {
		http.Error(w, fmt.Sprintf("wrong APIVersion or kind: %s - %s", request.APIVersion, request.Kind), http.StatusBadRequest)
		return

	}
	fmt.Printf("debug: %+v\n", request.Request)
	response := AdmissionReviewResponse{
		APIVersion: "admission.k8s.io/v1",
		Kind:       "AdmissionReview",
		Response: Response{
			UID:     request.Request.UID,
			Allowed: true,
		},
	}

	// add label if we're creating a pod
	if request.Request.Kind.Group == "" && request.Request.Kind.Version == "v1" && request.Request.Kind.Kind == "Pod" && request.Request.Operation == "CREATE" {
		patch := `[{"op": "add", "path": "/metadata/labels/myExtraLabel", "value": "webhook-was-here"}]`
		patchEnc := base64.StdEncoding.EncodeToString([]byte(patch))
		response.Response.PatchType = "JSONPatch"
		response.Response.Patch = patchEnc
	}

	out, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON output marshal error: %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	fmt.Printf("Got request, response: %s\n", string(out))
	fmt.Fprintln(w, string(out))
}
