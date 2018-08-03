# install jenkins
```
kubectl create -f serviceaccount.yaml
helm install --name jenkins --set rbac.install=true,Master.RunAsUser=1000,Master.FsGroup=1000 stable/jenkins
```
