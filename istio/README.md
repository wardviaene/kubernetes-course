# istio

## Kops configuration
```
kops edit cluster kubernetes.newtech.academy
```
Add:
```
  kubeAPIServer:
    admissionControl:
    - NamespaceLifecycle
    - LimitRanger
    - ServiceAccount
    - PersistentVolumeLabel
    - DefaultStorageClass
    - DefaultTolerationSeconds
    - MutatingAdmissionWebhook
    - ValidatingAdmissionWebhook
    - ResourceQuota
    - NodeRestriction
    - Priority
```

## download (1.0.1):
```
cd ~
wget https://github.com/istio/istio/releases/download/1.0.1/istio-1.0.1-linux.tar.gz
tar -xzvf istio-1.0.1-linux.tar.gz
cd istio-1.0.1
echo 'export PATH="$PATH:/home/ubuntu/istio-1.0.1/bin"' >> ~/.profile
```

## Download (latest):
```
cd ~
curl -L https://git.io/getLatestIstio | sh -
echo 'export PATH="$PATH:/home/ubuntu/istio-0.7.1/bin"' >> ~/.profile # change 0.7.1 in your version
cd istio-1.0.1 # change 1.0.1 in your version
```

## Istio install

Apply CRDs:

```
kubectl apply -f ~/istio-1.0.1/install/kubernetes/helm/istio/templates/crds.yaml
```

Wait a few seconds.


Option 1: with no mutual TLS authentication
```
kubectl apply -f ~/istio-1.0.1/install/kubernetes/istio-demo.yaml
```

Option 2: or with mutual TLS authentication
```
kubectl apply -f ~/istio-1.0.1/install/kubernetes/istio-demo-auth.yaml
```

## Example app

### Example app (from istio)
```
export PATH="$PATH:/home/ubuntu/istio-1.0.1/bin"
kubectl apply -f <(istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml)
```

### Hello world app 
```
export PATH="$PATH:/home/ubuntu/istio-1.0.1/bin"
kubectl apply -f <(istioctl kube-inject -f helloworld.yaml)
kubectl apply -f helloworld-gw.yaml
```
