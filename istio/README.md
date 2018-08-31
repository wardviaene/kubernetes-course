# istio

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

## Istio install

Apply CRDs:

```
kubectl apply -f install/kubernetes/helm/istio/templates/crds.yaml
```

Wait a few seconds.


Option 1: with no mutual TLS authentication
```
kubectl apply -f install/kubernetes/istio.yaml
```

Option 2: or with mutual TLS authentication
```
kubectl apply -f install/kubernetes/istio-auth.yaml
```

## Example app

Example app (from istio)
```
export PATH="$PATH:/home/ubuntu/istio-1.0.1/bin"
kubectl apply -f <(istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml)
```


### Traffic management

Add default route to v1:
```
istioctl create -f samples/bookinfo/routing/route-rule-all-v1.yaml
```

Route traffic to v2 if rule matches
```
istioctl replace -f samples/bookinfo/routing/route-rule-reviews-test-v2.yaml
```

Route 50% of traffic between v1 and v3:
```
istioctl replace -f samples/bookinfo/routing/route-rule-reviews-50-v3.yaml
```

## Distributed tracing

Enable zipkin:
```
kubectl apply -f install/kubernetes/addons/zipkin.yaml
```

Enable Jaeger:
```
kubectl delete -f install/kubernetes/addons/zipkin.yaml # if zipkin was installed, delete it first
kubectl apply -n istio-system -f https://raw.githubusercontent.com/jaegertracing/jaeger-kubernetes/master/all-in-one/jaeger-all-in-one-template.yml

```

