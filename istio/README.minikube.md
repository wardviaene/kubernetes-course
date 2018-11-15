# istio

## Launch minikube

* Install the latest release of minikube: https://github.com/kubernetes/minikube

```
minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 --vm-driver=`your_vm_driver_choice`
```


## download (1.0.3):
```
cd ~
wget https://github.com/istio/istio/releases/download/1.0.3/istio-1.0.3-linux.tar.gz # linux
wget https://github.com/istio/istio/releases/download/1.0.3/istio-1.0.3-osx.tar.gz # macos
tar -xzvf istio-1.0.3-linux.tar.gz
cd istio-1.0.3
echo 'export PATH="$PATH:~/istio-1.0.3/bin"' >> ~/.profile
```

## Download (latest):
```
cd ~
curl -L https://git.io/getLatestIstio | sh -
echo 'export PATH="$PATH:~/istio-1.0.3/bin"' >> ~/.profile # change 1.0.3 in your version
cd istio-1.0.3 # change 1.0.3 in your version
```

## Istio install

Apply CRDs:

```
kubectl apply -f ~/istio-1.0.3/install/kubernetes/helm/istio/templates/crds.yaml
```

Wait a few seconds.


Option 1: with no mutual TLS authentication (the next demos assume no mutual TLS)
```
kubectl apply -f ~/istio-1.0.3/install/kubernetes/istio-demo.yaml
```

Option 2: or with mutual TLS authentication
```
kubectl apply -f ~/istio-1.0.3/install/kubernetes/istio-demo-auth.yaml
```

## Disable LoadBalancer

Replace LoadBalancer in NodePort with:

```
kubectl edit -n istio-system svc/istio-ingressgateway
```

Get the URL with:
```
minikube service istio-ingressgateway -n istio-system --url
```

## Example app

### Example app (from istio)
```
export PATH="$PATH:~/istio-1.0.3/bin"
kubectl apply -f <(istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml)
```

### Hello world app 
```
export PATH="$PATH:~/istio-1.0.3/bin"
wget https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld.yaml
kubectl apply -f <(istioctl kube-inject -f helloworld.yaml)
kubectl apply -f https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld-gw.yaml
```

### Mutual TLS example
Create pods, services, destinationrules, virtualservices
```
wget https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld-tls.yaml
kubectl create -f <(istioctl kube-inject -f helloworld-tls.yaml)
kubectl create -f https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld-legacy.yaml
```

### End-user authentication
```
wget https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld-jwt.yaml
kubectl create -f <(istioctl kube-inject -f helloworld-jwt.yaml)
kubectl create -f https://raw.githubusercontent.com/wardviaene/kubernetes-course/master/istio/helloworld-jwt-enable.yaml
```
