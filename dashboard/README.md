# Setting up the dashboard

## Start dashboard

Create dashboard:
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
```

## Create user

Create sample user (if using RBAC - on by default on new installs with kops / kubeadm):
```
kubectl create -f sample-user.yaml

```

## Get login token:
```
kubectl -n kube-system create token admin-user
kubectl -n kube-system get secret | grep admin-user
kubectl -n kube-system describe secret admin-user-token-<id displayed by previous command>
```

## Login to dashboard
Go to http://api.yourdomain.com:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login

Login: admin
Password: the password that is listed in ~/.kube/config (open file in editor and look for "password: ..."

Choose for login token and enter the login token from the previous step
