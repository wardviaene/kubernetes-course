# Files
There is no quickstart-for-gke.sh anymore. This has been replaced by quickstart.sh.

# setup storage
```
kubectl create -f storage.yml
```

# setup Operator
```
./quickstart.sh
```

# Create cluster

```
kubectl apply -f postgres-example.yaml
```

# Show cluster pods

```
kubectl get pods -n postgres-operator
```

# show secrets
```
kubectl get secrets -n postgres-operator hippo-pguser-hippo -o yaml |grep user |cut -d ':' -f2 |cut -d ' ' -f2 |base64 --decode
kubectl get secrets -n postgres-operator hippo-pguser-hippo -o yaml |grep password |cut -d ':' -f2 |cut -d ' ' -f2 |base64 --decode
kubectl get secrets -n postgres-operator hippo-pguser-hippo -o yaml |grep host |cut -d ':' -f2 |cut -d ' ' -f2 |base64 --decode
```

# connect to psql

Use user, password, and host from previous step.

```
kubectl run -n postgres-operator -it --rm --image=postgres:10.4 psql-client -- psql -h hippo-primary.postgres-operator.svc -U hippo -W postgres
```

Note: When you see 'If you don't see a command prompt, try pressing enter.', you can enter the password


# Create read replic
Once you add replicas: 2 to the yaml definition, and you apply it, you'll see the new replica being spun up
```
kubectl apply -f postgres-example-scale.yaml
kubectl get pods -n postgres-operator
```

# Shutdown cluster
```
kubectl patch postgrescluster/hippo -n postgres-operator --type merge --patch '{"spec":{"shutdown": true}}'
```
