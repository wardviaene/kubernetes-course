# install jenkins
```
kubectl create -f serviceaccount.yaml
helm install --name jenkins --set rbac.create=true,master.runAsUser=1000,master.fsGroup=1000 stable/jenkins
```
