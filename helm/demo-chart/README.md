# Node demo app Chart

## Download dependencies
```
helm dependency update
```

## Install Chart
```
helm install .
```

## Upgrade Chart
```
helm upgrade --set image.tag=v0.0.2,mariadb.db.password=$DB_APP_PASS RELEASE .
```
Initialize helm

kubectl create helm-rbac.yml

helm init --service-account tiller

