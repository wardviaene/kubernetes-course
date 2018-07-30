# Helm

## Install helm
```
wget https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-linux-amd64.tar.gz
tar -xzvf helm-v2.9.1-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm
```

## Initialize helm

```
kubect create -f helm-rbac.yaml
helm init --service-account tiller
```
