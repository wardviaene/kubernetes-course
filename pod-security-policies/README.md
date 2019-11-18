# Pod Security Setup

## Minikube

```
mkdir -p ~/.minikube/files/etc/kubernetes/addons/
cp initial-psp.yaml ~/.minikube/files/etc/kubernetes/addons/psp.yaml
minikube start --extra-config=apiserver.enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,PodSecurityPolicy
```
