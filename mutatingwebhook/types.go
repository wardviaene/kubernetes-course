package main

//AdmissionReviewRespons for replies to incoming webhooks
type AdmissionReviewResponse struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Response   Response `json:"response"`
}
type Response struct {
	UID       string `json:"uid"`
	Allowed   bool   `json:"allowed"`
	Patch     string `json:"patch,omitempty"`
	PatchType string `json:"patchType,omitempty"`
}

//AdmissionReviewRequest for incoming webhooks
type AdmissionReviewRequest struct {
	APIVersion string  `json:"apiVersion"`
	Kind       string  `json:"kind"`
	Request    Request `json:"request"`
}
type Kind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}
type Resource struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}
type RequestKind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}
type RequestResource struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}
type Extra struct {
	SomeKey []string `json:"some-key"`
}
type UserInfo struct {
	Username string   `json:"username"`
	UID      string   `json:"uid"`
	Groups   []string `json:"groups"`
	Extra    Extra    `json:"extra"`
}
type Object struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}
type OldObject struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}
type Options struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}
type Request struct {
	UID                string          `json:"uid"`
	Kind               Kind            `json:"kind"`
	Resource           Resource        `json:"resource"`
	SubResource        string          `json:"subResource"`
	RequestKind        RequestKind     `json:"requestKind"`
	RequestResource    RequestResource `json:"requestResource"`
	RequestSubResource string          `json:"requestSubResource"`
	Name               string          `json:"name"`
	Namespace          string          `json:"namespace"`
	Operation          string          `json:"operation"`
	UserInfo           UserInfo        `json:"userInfo"`
	Object             Object          `json:"object"`
	OldObject          OldObject       `json:"oldObject"`
	Options            Options         `json:"options"`
	DryRun             bool            `json:"dryRun"`
}
