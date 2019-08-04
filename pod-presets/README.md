# PodPresets

# Alpha status
As long as the PodPresets is in alpha status, the following changes need to be made in kops:

```
spec:
  kubeAPIServer:
    enableAdmissionPlugins:
    - Initializers
    - NamespaceLifecycle
    - LimitRanger
    - ServiceAccount
    - PersistentVolumeLabel
    - DefaultStorageClass
    - DefaultTolerationSeconds
    - MutatingAdmissionWebhook
    - ValidatingAdmissionWebhook
    - NodeRestriction
    - ResourceQuota
    - PodPreset
    runtimeConfig:
      settings.k8s.io/v1alpha1: "true"

```
create the cluster again on aws

kops create cluster --name=dev.psamman.com --state=s3://kops-state-ammar --zones=eu-west-1a --node-count=3 --node-size=t2.micro --master-size=t2.micro --dns-zone=dev.psamman.com

now edit the cluster using this command 
kops edit cluster dev.psamman.com --state=s3://kops-state-ammar
and paste at the end of the file the new spces
kops update cluster dev.psamman.com --state=s3://kops-state-ammar --yes

# running the demo
First apply the PodPresets:
```
kubectl create -f pod-presets.yaml
```

Then run the deployments
```
kubectl create -f deployments.yaml
```
