# Helm

## Install helm (helm <3.0)
```
wget https://storage.googleapis.com/kubernetes-helm/helm-v2.11.0-linux-amd64.tar.gz
tar -xzvf helm-v2.11.0-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm
```

If you want a more recent helm version (>3.0), use:
```
wget https://get.helm.sh/helm-v3.2.3-linux-amd64.tar.gz
```

Make sure to use the correct instructions below for the different helm versions. You can get the latest helm version from https://github.com/helm/helm/releases. If you want to install helm on Windows/MacOS then have a look at https://helm.sh/docs/intro/install/ for instructions.

## Initialize helm (only for helm <3.0)

```
kubectl create -f helm-rbac.yaml
helm init --service-account tiller
```

## Initialize helm (helm >3.0)

Starting from 3.0, tiller has been removed. There is no need to create a service account for tiller anymore. The only command that is mandatory is the helm repo add command below:

```
helm repo add stable https://kubernetes-charts.storage.googleapis.com/
```

## Setup S3 helm repository
Make sure to set the default region in setup-s3-helm-repo.sh
```
./setup-s3-helm-repo.sh
```

## Package and push demo-chart

```
export AWS_REGION=yourregion # if not set in ~/.aws
helm package demo-chart
helm s3 push ./demo-chart-0.0.1.tgz my-charts
```
