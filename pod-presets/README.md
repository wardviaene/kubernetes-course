# PodPresets

# Alpha status
As long as the PodPresets is in alpha status, the following changes need to be made in kops:

Add:
```
spec:
  kubeAPIServer:
    appendAdmissionPlugins:
    - PodPreset
    runtimeConfig:
      settings.k8s.io/v1alpha1: "true"
```

# running the demo
First apply the PodPresets:
```
kubectl create -f pod-presets.yaml
```

Then run the deployments
```
kubectl create -f deployments.yaml
```
