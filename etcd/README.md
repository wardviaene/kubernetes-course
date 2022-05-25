# etcd

## HA Cluster
```
kops create cluster --name=kubernetes.newtech.academy --state=s3://kops-state-b429b --zones eu-west-1a,eu-west-1b,eu-west-1c --master-zones eu-west-1a,eu-west-1b,eu-west-1c --node-count=3 --node-size=t2.micro --master-size=t2.micro --dns-zone=kubernetes.newtech.academy
```

## Test backup
Create an object (wait 15 min after creating the object to make sure it the backup ran)
```
kubectl create configmap readme --from-file=README.md
```

## List backups
bash is not available anymore in this container image, so make sure you use "sh" instead of "bash"
```
kubectl exec -it etcd-main -n kube-system -- sh
./etcd-manager-ctl -backup-store=s3://kops-state-b429b/kubernetes.newtech.academy/backups/etcd/main/ list-backups
```
